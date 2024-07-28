package api

import (
	"fmt"

	"github.com/pimentafm/go-multithread/models"
	"github.com/pimentafm/go-multithread/utils"
)

func GetAddressFromBrasilAPI(cep string, resultChan chan<- *APIResponse, errChan chan<- error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	var brasilAPIResponse models.BrasilAPICEPInformation

	err := utils.FetchJSON(url, &brasilAPIResponse)
	if err != nil {
		errChan <- err
		return
	}

	address := models.Address{
		CEP:          brasilAPIResponse.CEP,
		Street:       brasilAPIResponse.Street,
		Neighborhood: brasilAPIResponse.Neighborhood,
		City:         brasilAPIResponse.City,
		State:        brasilAPIResponse.State,
	}

	resultChan <- &APIResponse{Address: &address, Source: "BrasilAPI"}
}
