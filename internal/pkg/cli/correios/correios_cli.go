package correios

import (
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/correios/format"
)

// CorreiosCLI provides the cli validation for correios command
type CorreiosCLI struct {
	service services.ICorreiosService
}

// ProvideCorreiosCLI provides a CorreiosCLI
var ProvideCorreiosCLI = func(baseURL string) CorreiosCLI {
	return CorreiosCLI{service: services.ProvideCorreiosService(baseURL)}
}

// RetrieveOrder prints the order data on the terminal
func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	format.FormatListToTerminal(response, err)
}
