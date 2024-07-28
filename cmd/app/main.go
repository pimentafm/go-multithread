package main

import (
	"fmt"
	"time"

	"github.com/pimentafm/go-multithread/api"
	"github.com/pimentafm/go-multithread/models"
)

func main() {
	cep := "01001000"
	timeout := 1 * time.Second
	address, source, err := getAddress(cep, timeout)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		printAddress(address, source)
	}
}

func getAddress(cep string, timeout time.Duration) (*models.Address, string, error) {
	resultChan := make(chan *api.APIResponse, 2)
	errChan := make(chan error, 2)

	go api.GetAddressFromBrasilAPI(cep, resultChan, errChan)
	go api.GetAddressFromViaCEP(cep, resultChan, errChan)

	select {
	case res := <-resultChan:
		return res.Address, res.Source, nil
	case err := <-errChan:
		return nil, "", err
	case <-time.After(timeout):
		return nil, "", fmt.Errorf("timeout after %v", timeout)
	}
}

func printAddress(address *models.Address, source string) {
	fmt.Printf("-------- %s ----------\n", source)
	fmt.Printf("CEP: %s\n", address.CEP)
	fmt.Printf("Street: %s\n", address.Street)
	fmt.Printf("Neighborhood: %s\n", address.Neighborhood)
	fmt.Printf("City: %s\n", address.City)
	fmt.Printf("State: %s\n", address.State)
}
