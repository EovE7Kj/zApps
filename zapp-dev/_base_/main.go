package main

import (
    "fmt"
    "log"

    "github.com/zenon-network/znn.go/znn"
    "github.com/zenon-network/znn.go/znn/common"
    "github.com/zenon-network/znn.go/znn/crypto"
)

func main() {
    
    client, err := znn.NewClient("http:
    if err != nil {
        log.Fatalf("Failed to initialize Zenon client: %v", err)
    }
    defer client.Close()

    
    keyPair := crypto.NewKeyPair()

    fmt.Printf("Generated Address: %s\n", keyPair.Address)

    
    accountState, err := client.Ledger.GetAccountState(keyPair.Address)
    if err != nil {
        log.Fatalf("Failed to get account state: %v", err)
    }

    fmt.Printf("Account State: %+v\n", accountState)

    
    txHash, err := deployDApp(client, keyPair, "<smart_contract_code_or_data>")
    if err != nil {
        log.Fatalf("Failed to deploy dApp: %v", err)
    }

    fmt.Printf("Transaction Hash for dApp deployment: %s\n", txHash)
}


func deployDApp(client *znn.Client, keyPair *crypto.KeyPair, dAppData string) (string, error) {
    
    block, err := client.Ledger.CreateTransactionBlock(keyPair.Address, common.AddressFromString("z1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqsgg9njn"), 0, []byte(dAppData))
    if err != nil {
        return "", fmt.Errorf("failed to create transaction block: %w", err)
    }

    
    block.Sign(keyPair)

    
    txHash, err := client.Ledger.SendTransaction(block)
    if err != nil {
        return "", fmt.Errorf("failed to send transaction: %w", err)
    }

    return txHash, nil
}
