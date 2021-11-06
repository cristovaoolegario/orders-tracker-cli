package mock_services

import (
	"orders-tracker-cli/internal/pkg/http/dto"
	"orders-tracker-cli/internal/pkg/http/services"
)

var _ services.ICorreiosService = (*CorreiosServiceMock)(nil)

var CorreiosServiceMockFindOrderByNumber func(orderNumber string) (*dto.CorreiosResponse, error)

type CorreiosServiceMock struct {}

func(cs *CorreiosServiceMock) FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error){
	return CorreiosServiceMockFindOrderByNumber(orderNumber)
}