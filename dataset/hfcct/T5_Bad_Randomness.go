package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract rappresenta il contratto intelligente per Hyperledger Fabric
type SmartContract struct {
	contractapi.Contract
}

// InsecureRandomNumberExample genera un numero casuale utilizzando una fonte di entropia debole
func (s *SmartContract) InsecureRandomNumberExample(ctx contractapi.TransactionContextInterface) (int64, error) {
	// Generazione di un numero casuale utilizzando una fonte non sicura
	randomNum, err := s.generateInsecureRandomNumber()
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}

	return randomNum, nil
}

// generateInsecureRandomNumber genera un numero casuale utilizzando la fonte di entropia debole
func (s *SmartContract) generateInsecureRandomNumber() (int64, error) {
	// Utilizzo di una fonte di entropia debole (crypto/rand) per generare un numero casuale
	// Questo è un esempio, nella realtà bisognerebbe utilizzare una fonte più sicura
	max := big.NewInt(100)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
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
