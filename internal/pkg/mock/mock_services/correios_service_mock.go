package mock

import (
	"orders-tracker-cli/internal/pkg/http/dto"
	"orders-tracker-cli/internal/pkg/http/services"
)

var _ services.ICorreiosService = (*CorreiosServiceMock)(nil)

// CorreiosServiceMockFindOrderByNumber represents the delegate function
var CorreiosServiceMockFindOrderByNumber func(orderNumber string) (*dto.CorreiosResponse, error)

// CorreiosServiceMock is a mock service of CorreiosService
type CorreiosServiceMock struct {}

// FindOrderByNumber is a mock function for CorreiosService.FindOrderByNumber
func(cs *CorreiosServiceMock) FindOrderByNumber(orderNumber string) (*dto.CorreiosResponse, error){
	return CorreiosServiceMockFindOrderByNumber(orderNumber)
}