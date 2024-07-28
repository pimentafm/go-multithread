package api

import (
	"fmt"

	"github.com/pimentafm/go-multithread/models"
	"github.com/pimentafm/go-multithread/utils"
)

func GetAddressFromViaCEP(cep string, resultChan chan<- *APIResponse, errChan chan<- error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	var address models.Address

	err := utils.FetchJSON(url, &address)
	if err != nil {
		errChan <- err
		return
	}

	resultChan <- &APIResponse{Address: &address, Source: "ViaCEP"}
}
