package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MoonBaZZe/znn-sdk-go/zenon"
	"github.com/MoonBaZZe/znn-sdk-go/zenon/znn"
	"github.com/MoonBaZZe/znn-sdk-go/zenon/znn/account"
	"github.com/MoonBaZZe/znn-sdk-go/zenon/znn/ledger"
)


func connectToZenon() (*zenon.Zenon, error) {
	client, err := zenon.NewZenon("ws:
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Zenon node: %v", err)
	}
	return client, nil
}


func checkBalance(client *zenon.Zenon, address string) error {
	acc, err := client.Ledger.GetAccountBlockByAddress(context.Background(), znn.ParseAddress(address))
	if err != nil {
		return fmt.Errorf("failed to get account balance: %v", err)
	}

	fmt.Printf("Account Balance for %s:\n", address)
	for _, balance := range acc.BalanceInfoMap {
		fmt.Printf("Token: %s, Balance: %d\n", balance.Token.Standard, balance.Balance.Uint64())
	}
	return nil
}


func sendTransaction(client *zenon.Zenon, fromAddress, toAddress string, amount uint64) error {
	account, err := account.NewAccount(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to create account: %v", err)
	}

	
	txHash, err := account.SendFunds(context.Background(), toAddress, znn.Amount(amount))
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent successfully. Tx Hash: %s\n", txHash)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args]")
		fmt.Println("Commands: ")
		fmt.Println("  balance <address>")
		fmt.Println("  send <fromAddress> <toAddress> <amount>")
		os.Exit(1)
	}

	command := os.Args[1]

	
	client, err := connectToZenon()
	if err != nil {
		log.Fatalf("Error connecting to Zenon: %v", err)
	}
	defer client.Close()

	switch command {
	case "balance":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run main.go balance <address>")
			os.Exit(1)
		}
		address := os.Args[2]
		err = checkBalance(client, address)
		if err != nil {
			log.Fatalf("Error checking balance: %v", err)
		}
	case "send":
		if len(os.Args) != 5 {
			fmt.Println("Usage: go run main.go send <fromAddress> <toAddress> <amount>")
			os.Exit(1)
		}
		fromAddress := os.Args[2]
		toAddress := os.Args[3]
		amount, err := strconv.ParseUint(os.Args[4], 10, 64)
		if err != nil {
			log.Fatalf("Invalid amount: %v", err)
		}
		err = sendTransaction(client, fromAddress, toAddress, amount)
		if err != nil {
			log.Fatalf("Error sending transaction: %v", err)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}

	fmt.Println("dApp execution completed successfully.")
}

