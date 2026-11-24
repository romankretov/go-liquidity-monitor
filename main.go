package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	rpcURL := "http://127.0.0.1:8545"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatal("couldn't reach rpc")
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal("read chain ID")
	}

	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal("block number")
	}

	fmt.Printf("Chain id: %v, Block Number: %v \n", chainID, blockNumber)

}
