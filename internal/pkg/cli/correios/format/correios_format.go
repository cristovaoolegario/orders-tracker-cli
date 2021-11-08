package format

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/components"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
)

func FormatListToTerminal(response *dto.CorreiosResponse, err error) {
	if err == nil {
		for _, event := range response.Objects[0].Events {
			fmt.Println(FormatEventByEventCodeAndEventType(event))
			fmt.Println(FormatDateTimeCreated(event.DateTimeCreated))
			fmt.Println()
		}
	} else {
		fmt.Printf("❌\t%s\n", err.Error())
	}
}

// FormatListToListItem formats the service.CorreiosService to a []list.Item output
func FormatListToListItem(response *dto.CorreiosResponse, err error) []list.Item {
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
