package api

import (
	"fmt"

	"github.com/pimentafm/go-multithread/models"
	"github.com/pimentafm/go-multithread/utils"
)

func GetAddressFromViaCEP(cep string, resultChan chan<- *APIResponse, errChan chan<- error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	var viaCEPResponse models.ViaCEPAPIInformation

	err := utils.FetchJSON(url, &viaCEPResponse)
	if err != nil {
		errChan <- err
		return
	}

	address := models.Address{
		CEP:          viaCEPResponse.CEP,
		Street:       viaCEPResponse.Logradouro,
		Neighborhood: viaCEPResponse.Bairro,
		City:         viaCEPResponse.Localidade,
		State:        viaCEPResponse.UF,
	}

	resultChan <- &APIResponse{Address: &address, Source: "ViaCEP"}
}
