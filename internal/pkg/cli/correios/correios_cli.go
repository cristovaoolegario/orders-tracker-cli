package correios

import (
	"fmt"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/correios/format"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
)

// CorreiosCLI provides the cli validation for correios command
type CorreiosCLI struct {
	service services.ICorreiosService
}

// ProvideCorreiosCLI provides a CorreiosCLI
var ProvideCorreiosCLI = func(baseURL, validationURL string) CorreiosCLI {
	return CorreiosCLI{service: services.ProvideCorreiosService(baseURL, validationURL)}
}

// RetrieveOrder prints the order data on the terminal
func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	fmt.Print(formatListToString(response, err))
}

func formatListToString(response *dto.CorreiosResponse, err error) string {
	var resultString = "\n"
	if err == nil {
		for _, event := range response.Objects[0].Events {
			resultString += fmt.Sprintf("%s\n%s\n\n", format.FormatEventByEventTypeAndEventStatus(event), format.FormatDateTimeCreated(event.DateTimeCreated))
		}
	} else {
		resultString += fmt.Sprintf("❌\t%s\n", err.Error())
	}

	return resultString
}
