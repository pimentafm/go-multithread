package main

import (
	"fmt"
	"time"

	"github.com/pimentafm/go-multithread/api"
	"github.com/pimentafm/go-multithread/models"
)

func main() {
	cep := "35780000"
	timeout := 1 * time.Second
	address, source, err := getAddress(cep, timeout)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Address from %s:\n%+v\n", source, address)
	}
}

func getAddress(cep string, timeout time.Duration) (*models.Address, string, error) {
	resultChan := make(chan *api.APIResponse, 2)
	errChan := make(chan error, 2)

	go api.GetAddressFromViaCEP(cep, resultChan, errChan)

	return nil, "", fmt.Errorf("timeout after %v", timeout)
}
