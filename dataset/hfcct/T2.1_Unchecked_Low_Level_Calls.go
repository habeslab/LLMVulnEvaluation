package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract rappresenta il contratto intelligente per Hyperledger Fabric
type SmartContract struct {
	contractapi.Contract
}

// ExternalLibraryCallingExample esegue una chiamata a una libreria esterna non consentita
func (s *SmartContract) ExternalLibraryCallingExample(ctx contractapi.TransactionContextInterface) (string, error) {
	// Esempio di importazione di un pacchetto non consentito 'database/sql'
	_, err := fmt.Println("Hello, world!")
	if err != nil {
		return "", err
	}

	return "External library calling example complete", nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}
