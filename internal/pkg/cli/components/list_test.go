package components

import (
	"testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func TestModel_Init(t *testing.T) {
	t.Run("Should return nil when model is initialized", func(t *testing.T) {
		model := model{}
		cmd := model.Init()

		if cmd != nil {
			t.Fatalf("Cmd should be nil when initialized")
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
			t.Fatalf("Should update the model when Ctrl+C on the screen")
		}
	})

	t.Run("Should return resized window when msg is a Window Size Message", func(t *testing.T) {

		model := model{list.NewModel(
			[]list.Item{
				Item{Text: "bla", Time: "2021"},
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
			list.NewModel(
				[]list.Item{
					Item{Text: "bla", Time: "2021"},
				}, list.NewDefaultDelegate(), 20, 20)}

		viewString := model.View()

		if viewString == "" {
			t.Fatalf("Should render a string when View is called")
		}
	})
}
