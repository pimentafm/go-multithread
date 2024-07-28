package api

import "github.com/pimentafm/go-multithread/models"

type APIResponse struct {
	Address *models.Address
	Source  string
}
