#!/bin/sh

solc - <auction6.sol --optimize --combined-json abi,bin -o ./out/ --overwrite
abigen --combined-json ./out/combined.json --pkg contracts --out core.go 

