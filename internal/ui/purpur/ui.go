package purpur

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#0a0100")).
			Background(lipgloss.Color("#eb4034"))

	defaultStyle = lipgloss.NewStyle().
			Bold(false).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#000000"))

	activeCurrStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#03fc03")).
			Background(lipgloss.Color("#7f8085"))
)

func Start() error {
	m := newModel()
	m1, err := tea.NewProgram(m).Run()
	if err != nil {
		return err
	}
	if m, ok := m1.(*mineVersionsModel); ok {
		if m.err != nil {
			return err
		}
		msg := m.SelectedVersion()
		if msg == "" {
			fmt.Println("\n", defaultStyle.Render("no version selected"))
			fmt.Println()
			return nil
		}
		msg = "selected version: " + msg
		bar := "\t" + strings.Repeat("#", len(msg)+4) + "\t"
		msg = fmt.Sprintf("\t# %s #\t", msg)
		fmt.Println("\n",
			defaultStyle.Render(bar), "\n",
			defaultStyle.Render(msg), "\n",
			defaultStyle.Render(bar))
		fmt.Println()
		fmt.Println()
	}
	return nil
}

func displayErrorDetails(err error) string {
	msg := err.Error()
	bar := "\t" + strings.Repeat("#", len(msg)+4) + "\t"
	msg = fmt.Sprintf("\t# %s #\t", msg)
	return fmt.Sprintf(`
%s
%s
%s

`, errorStyle.Render(bar), errorStyle.Render(msg), errorStyle.Render(bar))
}
