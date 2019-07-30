package externalinputservice

import (
    "os"

    "github.com/azillion/edgcm-converter/climate"
)



type ExternalInputService struct {

}

func NewExternalInputService() *ExternalInputService {
	service := new(ExternalInputService)
	return service
}