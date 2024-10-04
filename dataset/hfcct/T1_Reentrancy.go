package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) SimulateConcurrency(ctx contractapi.TransactionContextInterface) (string, error) {
	// Simulazione di utilizzo improprio di goroutine
	go func() {
		fmt.Println("Goroutine in action")
	}()

	// Attesa per un breve periodo per consentire il completamento della goroutine (che non è supportata in Fabric)
	time.Sleep(100 * time.Millisecond)

	return "Concurrency simulation complete", nil
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
