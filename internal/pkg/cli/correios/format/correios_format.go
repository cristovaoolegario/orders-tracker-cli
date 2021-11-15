package format

import (
	"fmt"
	"time"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
)

// FormatEventByEventTypeAndEventStatus formats a string based on the dto.Event props
func FormatEventByEventTypeAndEventStatus(event dto.Event) string {
	searchString := fmt.Sprintf("%s%s", event.Type, event.Status)
	return fmt.Sprintf("%s\t%s", pkg.IconDictionary[searchString], event.Description)
}

// FormatDateTimeCreated formats the date
func FormatDateTimeCreated(date string) string {
	formattedTime, err := time.Parse("02012006150405", date)
	if err != nil {
		return "⏱\tNo time registered for operation"
	}
	return fmt.Sprintf("⏱\t%s", formattedTime.Format("02 Jan 06 15:04"))
}
