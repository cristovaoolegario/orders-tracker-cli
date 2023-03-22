package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
)

// ICorreiosService represents an interface of CorreiosService
type ICorreiosService interface {
	FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error)
}

// CorreiosService is the service that calls correios APIs
type CorreiosService struct {
	client        *http.Client
	baseURL       string
	validationURL string
}

// ProvideCorreiosService provides a new Correios service
var ProvideCorreiosService = func(url string, validationUrl string) ICorreiosService {
	return &CorreiosService{&http.Client{
		Timeout: time.Duration(10) * time.Second,
	}, url, validationUrl}
}

func (cs *CorreiosService) getValidationToken() (*dto.TokenType, error) {
	r, err := http.NewRequest("POST", cs.validationURL, bytes.NewBuffer(pkg.ValidationData))
	if err != nil {
		return nil, errors.New("cannot craft request")
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Dart/2.18 (dart:io)")
	resp, err := cs.client.Do(r)
	if err != nil {
		return nil, errors.New("cannot send request")
	}
	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("cannot decode body")
	}
	res := dto.TokenType{}
	json.Unmarshal(_body, &res)
	return &res, nil
}

// FindOrderByNumber returns the order data or a error
func (cs *CorreiosService) FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error) {
	token, err := cs.getValidationToken()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(cs.baseURL, orderNumber), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Dart/2.18 (dart:io)")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("app-check-token", token.Token)
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
