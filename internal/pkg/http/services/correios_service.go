package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"io/ioutil"
	"net/http"
	"time"
)

// ICorreiosService represents an interface of CorreiosService
type ICorreiosService interface {
	FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error)
}

// CorreiosService is the service that calls correios APIs
type CorreiosService struct {
	client  *http.Client
	baseURL string
}

// ProvideCorreiosService provides a new Correios service
var ProvideCorreiosService = func(url string) ICorreiosService {
	return &CorreiosService{&http.Client{
		Timeout: time.Duration(10) * time.Second,
	}, url}
}

// FindOrderByNumber returns the order data or a error
func (cs *CorreiosService) FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(cs.baseURL, orderNumber), nil)
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
	errMsg := responseObject.Objects[0].Message
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}
	return &responseObject, nil
}
