package format

import (
	"fmt"
	"testing"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
)

func TestFormatEventByEventCode(t *testing.T) {
	t.Run("Should return correct icon and description when event code exists", func(t *testing.T) {
		testItems := []struct {
			Code string
			Type string
			Icon string
		}{
			{"BDE", "01", "🎁"},
			{"BDE", "20", "📪"},
			{"OEC", "01", "🙌"},
			{"DO", "01", "🚚"},
			{"RO", "01", "🚚"},
			{"PO", "01", "📦"},
			{"PO", "09", "💤"},
			{"PAR", "10", "✅"},
			{"PAR", "16", "🛬"},
			{"PAR", "17", "💸"},
			{"PAR", "18", "🗺"},
			{"PAR", "21", "🔎"},
			{"PAR", "24", "🔙"},
			{"PAR", "26", "🙅"},
			{"PAR", "31", "🤑"},
			{"", "", "🚧"},
		}

		for _, item := range testItems {
			formattedString := FormatEventByEventTypeAndEventStatus(dto.Event{
				Code:        item.Code,
				Type:        item.Type,
				Description: "test",
			})

			expected := fmt.Sprintf("%s\ttest", item.Icon)
			if formattedString != expected {
				t.Errorf("Expected: %s Got: %s", expected, formattedString)
			}
		}
	})
}

func TestFormatDateTimeCreated(t *testing.T) {
	t.Run("Should return No time registered when theres an error formatting", func(t *testing.T) {
		result := FormatDateTimeCreated("")
		expected := "⏱\tNo time registered for operation"
		if result != expected {
			t.Fatalf("Expected: %s, Received: %s", expected, result)
		}
	})

	t.Run("Should return time formatted in DD/MM/YYYY when theres no error formatting", func(t *testing.T) {
		items := []struct {
			input    string
			expected string
		}{
			{"2021-11-04T15:25:08", "04 Nov 21 15:25"},
			{"1990-12-15T09:13:08", "15 Dec 90 09:13"},
		}

		for _, item := range items {
			result := FormatDateTimeCreated(item.input)
			expected := fmt.Sprintf("⏱\t%s", item.expected)
			if result != expected {
				t.Fatalf("Expected: %s, Received: %s", expected, result)
			}
		}
	})
}
