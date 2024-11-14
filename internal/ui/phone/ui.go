package phone

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#0a0100")).
			Background(lipgloss.Color("#eb4034"))

	defaultInputStyle = lipgloss.NewStyle().
				Bold(false).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#000000"))

	activeInputStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#03fc03")).
				Background(lipgloss.Color("#7f8085"))
)

func Start() error {
	m := newContactModel()
	m1, err := tea.NewProgram(m).Run()
	if err != nil {
		return err
	}
	if m, ok := m1.(*contactModel); ok {
		fmt.Println(m.String())
	}
	return nil
}
