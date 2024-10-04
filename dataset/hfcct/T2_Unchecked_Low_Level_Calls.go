package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) UncheckedErrorsExample(ctx contractapi.TransactionContextInterface) error {
	// Esempio di utilizzo di `_` senza gestione degli errori
	_, err := s.getSomeData(ctx)
	if err != nil {
		// Ignorando l'errore
		_ = err // Attenzione: utilizzo di `_` senza gestione corretta dell'errore
	}

	// Operazione successiva che potrebbe non avere informazioni sull'errore

	return nil
}

func (s *SmartContract) getSomeData(ctx contractapi.TransactionContextInterface) (string, error) {
	// Simulazione di una funzione che restituisce dati e un possibile errore
	return "some data", fmt.Errorf("error fetching data")
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
