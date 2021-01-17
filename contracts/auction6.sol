// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.0;


contract auction6 {

struct Bid  {
    uint80  amount;
    uint32  entry;
    uint16  position;
}

mapping(address => Bid) entries;
uint256 start;
uint256 end;
uint256 reserve;
uint256 increment;
uint256 entry;
uint256 public numberOfEntries;
uint256 checkPos;
bool    public failed;
bool    public verified;

address lastAddress;
uint16  public pos;


address[] winners;

    event FirstBid(uint256 entry, address bidder, uint256 bid);
    event IncreasedBid(uint256 entry, address bidder, uint256 oldBid, uint256 increment);
    
    event Verified();
    event VerificationFailed(uint position);

    event ToBeContinued(address lastAddress, uint16 lPos);


    constructor(uint256 _start, uint256 _end, uint256 _reserve, uint256 _increment) {
        start = _start;
        end = _end;
        reserve = _reserve;
        increment = _increment;
    }

    function bid() external payable {
        require(block.timestamp >= start && block.timestamp <= end,"Auction not currently running");
        uint80 oldBid = entries[msg.sender].amount;
        if (oldBid > 0) {
            require(msg.value >= increment,"increment too small");
            emit IncreasedBid(entry,msg.sender,oldBid,msg.value);
        } else {
            require(msg.value >= reserve,"bid does not meet reserve price");
            numberOfEntries++;
            emit FirstBid(entry,msg.sender, msg.value);
        }
        Bid memory _newBid = Bid(oldBid+uint80(msg.value),uint32(entry++),0);
        entries[msg.sender] = _newBid;
    }

    function winningSequence(address[] memory sequence, uint256 count) external {
        require(!verified,"sequence already verified");
        require(sequence.length == numberOfEntries,"sequence length incorrect");
        winners = sequence;
        checkPos = 0;
        failed = false;
        verifySequence(count);
    }

    function verifySequence(uint256 count) public {
        uint startPos;
        require(!verified,"sequence already verified");
        require(!failed, "sequence already failed. Try again");
        require(checkPos < numberOfEntries,"this should never happen");
        if (checkPos == 0) {
            startPos++;
        }
        Bid memory prev = entries[winners[checkPos+startPos-1]];
        if (checkPos == 0) {
            checkPos++;
        }
        for (uint j = startPos; j < count; j++) {
            if (checkPos == numberOfEntries) {
                verified = true;
                emit Verified();
                return;
            }
            Bid memory current = entries[winners[checkPos++]];
            if (
                (prev.amount < current.amount) ||
                ((prev.amount == current.amount) && (prev.entry > current.entry))
            ) {
                failed = true;
                emit VerificationFailed(checkPos-1);
                return;
            }
            prev = current;
        }
        if (checkPos == numberOfEntries) {
            verified = true;
            emit Verified();
        }
    }


    function newWinningSequence(address[] memory sequence, bool reset) public {
        uint startPos = 0;
        if (reset) {
            lastAddress = sequence[0];
            startPos = 1;
            pos = 0;
        }
        uint16 lPos = pos;

        Bid memory prev = entries[lastAddress];
        for (uint j = startPos; j < sequence.length; j++) {
            if (lPos == numberOfEntries-1) {
                verified = true;
                emit Verified();
                pos = lPos;
                return;
            }
            Bid memory current = entries[sequence[j]];
            if (
                (prev.amount < current.amount) ||
                ((prev.amount == current.amount) && (prev.entry > current.entry)) ||
                (current.amount == 0) 
            ) {
                failed = true;
                emit VerificationFailed(lPos);
                pos = lPos;
                return;
            }
            current.position = ++lPos;
            entries[sequence[j]] = current;
            prev = current;
        }
        if (lPos == numberOfEntries-1) {
            verified = true;
            pos = lPos;
            emit Verified();
            return;
        } 
        lastAddress = sequence[sequence.length-1];
        pos = lPos;
        emit ToBeContinued(lastAddress,lPos);
    }


}

/* 
Gas Used  934758
Gas Used  926607
Gas Used  926607
Gas Used  926643
Gas Used  925854
*/