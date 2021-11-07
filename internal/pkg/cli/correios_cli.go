package cli

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/components"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
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
func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) []list.Item {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	renderList := []list.Item{}
	if err == nil {
		for _, event := range response.Objects[0].Events {
			item := components.Item{
				Text: FormatEventByEventCodeAndEventType(event),
				Time: FormatDateTimeCreated(event.DateTimeCreated),
			}
			renderList = append(renderList, []list.Item{item}...)
		}
	} else {
		renderList = []list.Item{
			components.Item{
				Text: fmt.Sprintf("❌\t%s", err.Error()),
				Time: "",
			},
		}
	}

	return renderList
}

// RetrieveOrderAsList renders A list on the terminal
func (cli *CorreiosCLI) RetrieveOrderAsList(orderNumber string) {
	items := cli.RetrieveOrder(orderNumber)
	components.RenderList(orderNumber, items)
}

// FormatEventByEventCodeAndEventType formats a string based on the dto.Event props
func FormatEventByEventCodeAndEventType(event dto.Event) string {
	searchString := fmt.Sprintf("%s%s", event.Code, event.Type)
	return fmt.Sprintf("%s\t%s", pkg.IconDictionary[searchString], event.Description)
}

// FormatDateTimeCreated formats the date
func FormatDateTimeCreated(date string) string {
	formattedTime, err := time.Parse("2006-01-02T15:04:05", date)
	if err != nil {
		return "⏱\tNo time registered for operation"
	}
	return fmt.Sprintf("⏱\t%s", formattedTime.Format("02 Jan 06 15:04"))
}
