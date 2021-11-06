package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	mock "github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/mock/mock_services"
	"io/ioutil"
	"os"
	"testing"
)

func TestCorreiosCLI_RetrieveOrder(t *testing.T) {

	t.Run("Should print error message when don't find order", func(t *testing.T) {
		cli := CorreiosCLI{}
		cli.service = &mock.CorreiosServiceMock{}

		mock.CorreiosServiceMockFindOrderByNumber = func(orderNumber string) (*dto.CorreiosResponse, error) {
			return nil, errors.New("Test error")
		}

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		cli.RetrieveOrder("not_an_order_number")

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		if string(out) != fmt.Sprintf("❌ - Test error\n") {
			t.Errorf("Expected %q, got %q", "❌ - Test error\n", out)
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

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		cli.RetrieveOrder("an_order_number")

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		if string(out) != fmt.Sprintf("🎁 - Objeto entregue ao destinatário\n⏱ - 2021-09-06T15:58:08\n\n") {
			t.Errorf("Expected %q, got %q", fmt.Sprintf("🎁 - Objeto entregue ao destinatário\n⏱ - 2021-09-06T15:58:08\n"), out)
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
			{"PAR", "10", "🔎✅"},
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

			if formattedString != fmt.Sprintf("%s - test", item.Icon) {
				t.Errorf("Expected: %s Got: %s", fmt.Sprintf("%q - test\n", item.Icon), formattedString)
			}
		}
	})
}
