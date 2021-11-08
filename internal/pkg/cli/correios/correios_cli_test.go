package correios

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

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
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		cli.RetrieveOrder("not_an_order_number")

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		expected := "\n❌\tTest error\n"
		if string(out) != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})

	t.Run("Should print events description and date when there's an order", func(t *testing.T) {
		cli := CorreiosCLI{}
		cli.service = &mock.CorreiosServiceMock{}

		mock.CorreiosServiceMockFindOrderByNumber = func(orderNumber string) (*dto.CorreiosResponse, error) {
			jsonFile, _ := os.Open("../../mock/mock_data/valid_code_with_one_event.json")

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

		expected := "\n🎁\tObjeto entregue ao destinatário\n⏱\t06 Sep 21 15:58\n\n"
		if string(out) != expected {
			t.Errorf("Expected %q, got %q", expected, out)
		}
	})
}
