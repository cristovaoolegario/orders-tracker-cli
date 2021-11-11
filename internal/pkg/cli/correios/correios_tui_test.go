package correios

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/components"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	mock "github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/mock/mock_services"
)

func TestProvideNewModel(t *testing.T) {
	t.Run("Should provide model when other dependencies are ok", func(t *testing.T) {
		model := ProvideNewModel("order_number", "http://example_url.com")

		if model.orderNumber == "" || model.service == nil {
			t.Fatalf("Dependencies weren't provided correctly")
		}
	})
}

func TestFormatListToListItem(t *testing.T) {
	t.Run("Should add item when theres a elegible item", func(t *testing.T) {
		jsonFile, _ := os.Open("../../mock/mock_data/valid_code_with_one_event.json")

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		responseObject := dto.CorreiosResponse{}
		json.Unmarshal(byteValue, &responseObject)

		orderData := FormatListToListItem(&responseObject, nil)

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

	t.Run("Should add error item when theres no elegible itens", func(t *testing.T) {

		orderData := FormatListToListItem(nil, errors.New("Test error"))

		expected := "❌\tTest error"
		if orderData[0].(components.Item).Text != expected {
			t.Errorf("Expected %q, got %q", expected, orderData[0].(components.Item).Text)
		}
	})
}

func TestMountList(t *testing.T) {
	t.Run("Should mount list component with the right properties", func(t *testing.T) {
		model := MountList("test")

		expectedTitle := "test"
		if model.list.Title != expectedTitle {
			t.Fatalf("Should've set the correct title to the list. Expected '%s', Got '%s'", expectedTitle, model.list.Title)
		}
	})
}

func TestModel_Init(t *testing.T) {
	t.Run("Should return nil when model is initialized", func(t *testing.T) {
		model := model{}
		cmd := model.Init()

		if cmd == nil {
			t.Fatalf("Theres a problem Initializing the model")
		}
	})
}

func TestModel_Update(t *testing.T) {
	t.Run("Should return model when msg is a Ctrl+C key message", func(t *testing.T) {
		model := model{}

		updatedModel, err := model.Update(tea.KeyMsg{
			Type: tea.KeyCtrlC,
		})

		if updatedModel == nil && err != nil {
			t.Fatalf("Should had updated the model when Ctrl+C on the terminal")
		}
	})

	t.Run("Should set items when msg is a []list.Item", func(t *testing.T) {
		model := model{
			list: list.NewModel(
				[]list.Item{
					components.Item{Text: "bla", Time: "2021"},
				}, list.NewDefaultDelegate(), 20, 20)}

		updatedModel, err := model.Update([]list.Item{})

		if updatedModel == nil && err == nil {
			t.Fatalf("Should've return the updated model and a cmd when list is updated")
		}
	})

	t.Run("Should return resized window when msg is a Window Size Message", func(t *testing.T) {
		model := model{
			list: list.NewModel(
				[]list.Item{
					components.Item{Text: "bla", Time: "2021"},
				}, list.NewDefaultDelegate(), 20, 20)}

		updatedModel, err := model.Update(tea.WindowSizeMsg{Width: 10, Height: 10})

		if updatedModel == nil && err == nil {
			t.Fatalf("Should return the updated model and a cmd when window resizes")
		}
	})
}

func TestModel_View(t *testing.T) {
	t.Run("Should return string when view Model", func(t *testing.T) {
		model := model{
			list: list.NewModel(
				[]list.Item{
					components.Item{Text: "bla", Time: "2021"},
				}, list.NewDefaultDelegate(), 20, 20)}

		viewString := model.View()

		if viewString == "" {
			t.Fatalf("Should render a string when View is called")
		}
	})
}

func TestModel_LoadCmd(t *testing.T) {
	t.Run("Should return a tea.Cmd of the type []list.Item", func(t *testing.T) {
		model := model{}
		model.service = &mock.CorreiosServiceMock{}

		mock.CorreiosServiceMockFindOrderByNumber = func(orderNumber string) (*dto.CorreiosResponse, error) {
			return nil, errors.New("Test error")
		}

		cmd := model.LoadCmd()

		if cmd == nil {
			t.Fatalf("Should've load the cmd")
		}
	})
}
