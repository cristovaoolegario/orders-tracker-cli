package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
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
	body := fmt.Sprintf(`{ "code": "%s", "type": "LS" }`, orderNumber)
	req, err := http.NewRequest("POST", cs.baseURL, strings.NewReader(body))
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
	if len(responseObject.Objects[0].Events) == 0 {
		errMsg := responseObject.Objects[0].Category
		return nil, errors.New(errMsg)
	}
	return &responseObject, nil
}
