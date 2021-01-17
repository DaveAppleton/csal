package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"af.sorry.devcon6/contracts"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/memorykeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	baseClient *backends.SimulatedBackend
	cAddress   common.Address
	cContract  *contracts.Auction6
)

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	gaslimit := viper.GetUint64("gaslimit")
	funds, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	banker, _ := memorykeys.GetAddress("banker")
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		*banker: {Balance: funds},
	}, gaslimit)

	return baseClient, nil
}
func chkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func checkTxErrB(tx *types.Transaction, err error, expected uint64) {
	chkErr(err)
	baseClient.Commit()
	rctp, err := baseClient.TransactionReceipt(context.Background(), tx.Hash())
	chkErr(err)
	if rctp.Status != expected {
		chkErr(fmt.Errorf("transaction returns %d", rctp.Status))
	}
	fmt.Println("Gas Used ", rctp.GasUsed)
}

func checkTxErr(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusSuccessful)
}

func checkTxErrF(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusFailed)
}

func initViper() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()
}

var reverse map[common.Address]string

func nameOf(c common.Address) string {
	str := reverse[c]
	if len(str) > 0 {
		return str
	}
	return c.Hex()
}

func main() {
	var bidderAddresses []common.Address
	var winnerAddresses []common.Address
	var bidders []string
	var tx *types.Transaction
	var sendTxs []*types.Transaction
	var bidTxs []*types.Transaction
	opts := bind.FilterOpts{}

	reverse = make(map[common.Address]string)

	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	log.Println("new run")
	client, err := getClient()
	chkErr(err)
	sixtyEther := new(big.Int).Mul(etherUtils.OneEther(), big.NewInt(60))

	bankerTx, _ := memorykeys.GetTransactor("banker")
	for j := 0; j < 500; j++ {
		bidder := fmt.Sprintf("bidder%d", j)
		bidderAddress, _ := memorykeys.GetAddress(bidder)
		reverse[*bidderAddress] = bidder
		tx, err = sendEthereum("banker", *bidderAddress, sixtyEther, 23000, client)
		chkErr(err)
		sendTxs = append(sendTxs, tx)
		bidders = append(bidders, bidder)
		bidderAddresses = append(bidderAddresses, *bidderAddress)
		if (j+1)%50 == 0 {
			client.Commit()
		}
	}
	fmt.Println(len(bidders), "bidders")
	client.Commit()
	for j := 0; j < len(bidders); j++ {
		rct, err := client.TransactionReceipt(context.Background(), sendTxs[j].Hash())
		chkErr(err)
		if rct.Status == types.ReceiptStatusFailed {
			log.Fatal("Tx", j, "failed")
		}
	}
	fmt.Println("All accounts primed")
	startStamp := big.NewInt(1547590334)
	endStamp := big.NewInt(1548022334)
	cAddress, tx, cContract, err = contracts.DeployAuction6(bankerTx, client, startStamp, endStamp, etherUtils.OneEther(), etherUtils.PointOneEther())
	checkTxErr(tx, err)
	amount := etherUtils.OneEther()
	blockTime := client.Blockchain().CurrentBlock().Time()
	jump := new(big.Int).Sub(startStamp, big.NewInt(int64(blockTime)))
	client.AdjustTime(time.Duration(jump.Int64()) * time.Second)
	client.Commit()
	for j := 0; j < len(bidders); j++ {

		bidderTx, _ := memorykeys.GetTransactor(bidders[j])
		bidderTx.Value = amount
		tx, err = cContract.Bid(bidderTx)
		chkErr(err)
		bidTxs = append(bidTxs, tx)
		if j%2 == 0 {
			amount = new(big.Int).Add(amount, etherUtils.PointOneEther())
		}
		if (j+1)%25 == 0 {
			client.Commit()
		}
	}
	client.Commit()
	for j := 0; j < len(bidders); j++ {
		rct, err := client.TransactionReceipt(context.Background(), bidTxs[j].Hash())
		chkErr(err)
		if rct.Status == types.ReceiptStatusFailed {
			log.Fatal("Tx", j, "failed")
		}
		fmt.Println("bid", bidders[j], etherUtils.EtherToStr(bidTxs[j].Value()), rct.GasUsed)
		log.Println("bid", bidders[j], etherUtils.EtherToStr(bidTxs[j].Value()), rct.GasUsed)
	}
	fmt.Println("upBid")
	bidderTx, _ := memorykeys.GetTransactor(bidders[0])
	bidderTx.Value = etherUtils.OneEther()
	tx, err = cContract.Bid(bidderTx)
	checkTxErr(tx, err)
	fmt.Println("closing bidding")
	blockTime = client.Blockchain().CurrentBlock().Time()
	jump = new(big.Int).Sub(endStamp, big.NewInt(int64(blockTime)))
	client.AdjustTime(time.Duration(jump.Int64()) * time.Second)
	client.Commit()
	maxxx := len(bidders)
	if len(bidders)%2 == 0 {
		maxxx = len(bidders) - 1
		winnerAddresses = append(winnerAddresses, bidderAddresses[len(bidders)-1])
	}
	for j := 0; j < maxxx; j += 2 {
		if maxxx-j-2 < 0 {
			break
		}
		first := maxxx - j - 2
		second := maxxx - j - 1
		fmt.Println(first, second)
		winnerAddresses = append(winnerAddresses, bidderAddresses[first])
		winnerAddresses = append(winnerAddresses, bidderAddresses[second])
	}
	winnerAddresses = append(winnerAddresses, bidderAddresses[0])
	version := 2
	if version == 1 {
		tx, err = cContract.WinningSequence(bankerTx, winnerAddresses, big.NewInt(0))
		checkTxErr(tx, err)
		increment := 500
		for j := 0; j < len(bidders); j += increment {
			tx, err := cContract.VerifySequence(bankerTx, big.NewInt(int64(increment)))
			checkTxErr(tx, err)
			verified, err := cContract.Verified(nil)
			chkErr(err)
			failed, err := cContract.Failed(nil)
			chkErr(err)
			if verified || failed {
				fmt.Println(j, "verified", verified)
				fmt.Println(j, "failed", failed)
				log.Println(j, "verified", verified)
				log.Println(j, "failed", failed)
				break
			}
		}
	} else {
		pos := 0
		ln := 600
		reset := true
		fmt.Println("Winner Sequence length ", len(winnerAddresses))
		ne, err := cContract.NumberOfEntries(nil)
		chkErr(err)
		fmt.Println("Number of entries rcv", ne.String())
		totalGasUsed := uint64(0)
		for {

			if pos >= len(winnerAddresses) {
				fmt.Println("Break")
				break
			}
			if pos+ln < len(winnerAddresses) {
				tx, err = cContract.NewWinningSequence(bankerTx, winnerAddresses[pos:pos+ln], reset)
			} else {
				tx, err = cContract.NewWinningSequence(bankerTx, winnerAddresses[pos:], reset)
			}
			checkTxErr(tx, err)
			rct, _ := client.TransactionReceipt(context.Background(), tx.Hash())
			totalGasUsed += rct.GasUsed
			ps, err := cContract.Pos(nil)
			chkErr(err)
			fmt.Println("pos after", ps)
			reset = false
			verified, err := cContract.Verified(nil)
			chkErr(err)
			failed, err := cContract.Failed(nil)
			chkErr(err)
			if verified || failed {
				fmt.Println(ps, "verified", verified)
				fmt.Println(ps, "failed", failed)
				log.Println(ps, "verified", verified)
				log.Println(ps, "failed", failed)
				break
			}
			pos += ln
		}
		fmt.Println("total Gas Used : ", totalGasUsed)
		fmt.Println("POS = ", pos)
		f2, err := cContract.FilterToBeContinued(&opts)
		chkErr(err)
		for f2.Next() {
			fmt.Println("to be continued ",
				f2.Event.LPos,
				nameOf(f2.Event.LastAddress))
		}
	}
	fi, err := cContract.FilterVerificationFailed(&opts)
	chkErr(err)
	for fi.Next() {
		fmt.Println("failed at", fi.Event.Position.String())
	}
	vi, err := cContract.FilterVerified(&opts)
	chkErr(err)
	for vi.Next() {
		fmt.Println("Verified !!!")
	}
}
