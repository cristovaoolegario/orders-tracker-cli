package correios

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/components"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/correios/format"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/dto"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
)

type model struct {
	orderNumber string
	service     services.ICorreiosService
	list        list.Model
}

var ProvideNewModel = func(orderNumber, baseURL string) *model {
	return &model{
		orderNumber: orderNumber,
		service:     services.ProvideCorreiosService(baseURL),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		spinner.Tick,
		m.LoadCmd(),
	)
}

func (m model) LoadCmd() tea.Cmd {
	return func() tea.Msg {
		data, err := m.service.FindOrderByNumber(m.orderNumber)
		return FormatListToListItem(data, err)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, nil
		}
	case []list.Item:
		m.list.SetShowStatusBar(true)
		spinnerCmd := m.list.ToggleSpinner()
		cmd := m.list.SetItems(msg)
		return m, tea.Batch(cmd, spinnerCmd)
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	return m, tea.Batch(cmd)
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func RenderBubbleTeaList(orderNumber string) {
	m := MountList(orderNumber)

	p := tea.NewProgram(m)
	p.EnterAltScreen()

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func MountList(orderNumber string) *model {
	m := ProvideNewModel(orderNumber, pkg.CorreiosBaseURL)
	m.list = list.NewModel([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	m.list.Styles.Title = titleStyle
	m.list.SetShowStatusBar(false)
	m.list.SetSpinner(spinner.Pulse)
	m.list.StartSpinner()
	m.list.Title = m.orderNumber
	return m
}

func FormatListToListItem(response *dto.CorreiosResponse, err error) []list.Item {
	renderList := []list.Item{}
	if err == nil {
		for _, event := range response.Objects[0].Events {
			item := components.Item{
				Text: format.FormatEventByEventTypeAndEventStatus(event),
				Time: format.FormatDateTimeCreated(event.DateTimeCreated),
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
