package cli

import (
	"fmt"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
)

// CorreiosCLI provides the cli validation for correios command
type CorreiosCLI struct {
	service services.ICorreiosService
}

// ProvideCorreiosCLI provides a CorreiosCLI
func ProvideCorreiosCLI(baseURL string) CorreiosCLI {
	return CorreiosCLI{service: services.ProvideCorreiosService(baseURL)}
}

// RetrieveOrder prints the order data on the terminal
func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	if err == nil {
		for _, event := range response.Objects[0].Events {
			fmt.Println(FormatEventByEventCodeAndEventType(event))
			fmt.Println(FormatDateTimeCreated(event.DateTimeCreated))
			fmt.Println()
		}
	} else {
		fmt.Printf("❌ - %s\n", err.Error())
	}
}

// FormatEventByEventCodeAndEventType formats a string based on the dto.Event props
func FormatEventByEventCodeAndEventType(event dto.Event) string {
	searchString := fmt.Sprintf("%s%s", event.Code, event.Type)
	return fmt.Sprintf("%s - %s", pkg.IconDictionary[searchString], event.Description)
}

// FormatDateTimeCreated formats the date
func FormatDateTimeCreated(date string) string {
	return fmt.Sprintf("⏱ - " + date)
}
