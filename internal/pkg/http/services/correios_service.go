package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"orders-tracker-cli/internal/pkg/http/dto"
	"time"
)

type ICorreiosService interface {
	FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error)
}

// CorreiosService is the service that calls correios APIs
type CorreiosService struct {
	client  *http.Client
	baseUrl string
}

// ProvideCorreiosService provides a new Correios service
var ProvideCorreiosService = func(url string) ICorreiosService {
	return &CorreiosService{&http.Client{
		Timeout: time.Duration(10) * time.Second,
	}, url}
}

// FindOrderByNumber returns the order data or a error
func (cs *CorreiosService) FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(cs.baseUrl, orderNumber), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseObject := dto.CorreiosResponse{}
	json.Unmarshal(bodyBytes, &responseObject)
	errMsg := responseObject.Objetos[0].Mensagem
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}
	return &responseObject, nil
}
