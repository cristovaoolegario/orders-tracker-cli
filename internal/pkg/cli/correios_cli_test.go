package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/components"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	mock "github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/mock/mock_services"
)

func TestProvideCorreiosCLI(t *testing.T) {
	t.Run("Should return a CorreiosCLI instance", func(t *testing.T) {
		baseURL := "test"
		cli := ProvideCorreiosCLI(baseURL)

		if cli.service == nil {
			t.Errorf("Should have provided a service.")
		}
	})
}

func TestCorreiosCLI_RetrieveOrder(t *testing.T) {

	t.Run("Should print error message when don't find order", func(t *testing.T) {
		cli := CorreiosCLI{}
		cli.service = &mock.CorreiosServiceMock{}

		mock.CorreiosServiceMockFindOrderByNumber = func(orderNumber string) (*dto.CorreiosResponse, error) {
			return nil, errors.New("Test error")
		}

		orderData := cli.RetrieveOrder("not_an_order_number")

		expected := "❌ \tTest error"
		if len(orderData) == 0 && orderData[0].(components.Item).Text != expected {
			t.Errorf("Expected %q, got %q", expected, orderData)
		}
	})

	t.Run("Should print events description and date when there's an order", func(t *testing.T) {
		cli := CorreiosCLI{}
		cli.service = &mock.CorreiosServiceMock{}

		mock.CorreiosServiceMockFindOrderByNumber = func(orderNumber string) (*dto.CorreiosResponse, error) {
			jsonFile, _ := os.Open("../mock/mock_data/valid_code_with_one_event.json")

			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			responseObject := dto.CorreiosResponse{}
			json.Unmarshal(byteValue, &responseObject)
			return &responseObject, nil
		}

		orderData := cli.RetrieveOrder("an_order_number")

		expectedDescription := "🎁\tObjeto entregue ao destinatário"
		expectedTime := "⏱\t06 Sep 21 15:58"

		if len(orderData) >= 0 {
			if orderData[0].(components.Item).Text != expectedDescription {
				t.Errorf("Expected %q, got %q", expectedDescription, orderData[0].(components.Item).Text)
			}
			if orderData[0].(components.Item).Time != expectedTime {
				t.Errorf("Expected %q, got %q", expectedTime, orderData[0].(components.Item).Time)
			}
		}
	})
}

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
			formattedString := FormatEventByEventCodeAndEventType(dto.Event{
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
