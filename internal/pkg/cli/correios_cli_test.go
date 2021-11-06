package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"orders-tracker-cli/internal/pkg/http/dto"
	mock "orders-tracker-cli/internal/pkg/mock/mock_services"
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

		if string(out) != fmt.Sprintf("🚚 - Objeto entregue ao destinatário\n⏱ - 2021-09-06T15:58:08\n\n") {
			t.Errorf("Expected %q, got %q", fmt.Sprintf("🚚 - Objeto entregue ao destinatário\n⏱ - 2021-09-06T15:58:08\n"), out)
		}
	})
}