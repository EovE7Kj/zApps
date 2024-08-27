package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/MoonBaZZe/znn-sdk-go/zenon"
	"github.com/MoonBaZZe/znn-sdk-go/zenon/znn/account"
	"github.com/MoonBaZZe/znn-sdk-go/zenon/znn/plasma"
)


func connectToZenon() (*zenon.Zenon, error) {
	client, err := zenon.NewZenon("ws:
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Zenon node: %v", err)
	}
	return client, nil
}


func uploadContent(client *zenon.Zenon, filePath string, owner string) (string, error) {
	
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read HTML file: %v", err)
	}

	
	contentHex := hex.EncodeToString(content)

	
	acc, err := account.NewAccount(context.Background(), owner)
	if err != nil {
		return "", fmt.Errorf("failed to create account: %v", err)
	}

	
	txHash, err := acc.PlasmaDeployCode(context.Background(), contentHex)
	if err != nil {
		return "", fmt.Errorf("failed to deploy content to Zenon network: %v", err)
	}

	fmt.Printf("Content uploaded successfully. Tx Hash: %s\n", txHash)
	return txHash, nil
}


func retrieveContent(client *zenon.Zenon, txHash string) (string, error) {
	
	contentHex, err := client.Plasma.GetContractCode(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve content: %v", err)
	}

	
	content, err := hex.DecodeString(contentHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex content: %v", err)
	}

	return string(content), nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <command> <filePath/address>")
		fmt.Println("Commands: ")
		fmt.Println("  upload <filePath> <ownerAddress>")
		fmt.Println("  retrieve <txHash>")
		os.Exit(1)
	}

	command := os.Args[1]

	
	client, err := connectToZenon()
	if err != nil {
		log.Fatalf("Error connecting to Zenon: %v", err)
	}
	defer client.Close()

	switch command {
	case "upload":
		if len(os.Args) != 4 {
			fmt.Println("Usage: go run main.go upload <filePath> <ownerAddress>")
			os.Exit(1)
		}
		filePath := os.Args[2]
		ownerAddress := os.Args[3]
		txHash, err := uploadContent(client, filePath, ownerAddress)
		if err != nil {
			log.Fatalf("Error uploading content: %v", err)
		}
		fmt.Printf("Uploaded content with transaction hash: %s\n", txHash)
	case "retrieve":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run main.go retrieve <txHash>")
			os.Exit(1)
		}
		txHash := os.Args[2]
		content, err := retrieveContent(client, txHash)
		if err != nil {
			log.Fatalf("Error retrieving content: %v", err)
		}
		fmt.Printf("Retrieved content: \n%s\n", content)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}

	fmt.Println("dApp execution completed successfully.")
}

