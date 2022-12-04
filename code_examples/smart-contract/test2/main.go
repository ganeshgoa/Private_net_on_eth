package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/ganeshgoa/blockchain/code_examples/smart-contract/test2/store"
)

func main() {
	ctx := context.Background()

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// тут нужно передать какие-то параметры (ключ-значение) вместо 111
	keyin := strings.NewReader("111")
	auth, err := bind.NewTransactorWithChainID(keyin, "<<strong_password>>", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, storage, err := store.DeployStorage(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	tx, err = storage.Store(&bind.TransactOpts{}, big.NewInt(12345))
	if err != nil {
		log.Fatalf("Failed to call storage.Store: %v", err)
	}
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	val, err := storage.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to call storage.Retrieve: %v", err)
	}
	fmt.Printf("Value from storage.Retrieve: %v\n", val.String())
}
