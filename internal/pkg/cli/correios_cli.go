package cli

import (
	"fmt"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
	"time"
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
	if err == nil {
		for _, event := range response.Objects[0].Events {
			fmt.Println(FormatEventByEventCodeAndEventType(event))
			fmt.Println(FormatDateTimeCreated(event.DateTimeCreated))
			fmt.Println()
		}
	} else {
		fmt.Printf("❌ %s\n", err.Error())
	}
}

// FormatEventByEventCodeAndEventType formats a string based on the dto.Event props
func FormatEventByEventCodeAndEventType(event dto.Event) string {
	searchString := fmt.Sprintf("%s%s", event.Code, event.Type)
	return fmt.Sprintf("%s %s", pkg.IconDictionary[searchString], event.Description)
}

// FormatDateTimeCreated formats the date
func FormatDateTimeCreated(date string) string {
	formattedTime, err := time.Parse("2006-01-02T15:04:05", date)
	if err != nil{
		return "⏱ No time registered for operation"
	}
	return fmt.Sprintf("⏱ %s", formattedTime.Format("02 Jan 06 15:04"))
}
