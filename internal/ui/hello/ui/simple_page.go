package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// MODEL DATA

type simplePage struct{ text string }

func newSimplePage(text string) simplePage {
	return simplePage{text: text}
}

func (s simplePage) Init() tea.Cmd { return nil }

// VIEW

func (s simplePage) View() string {
	textLen := len(s.text)
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf(
		"%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
		topAndBottomBar, s.text, topAndBottomBar,
	)
}

// UPDATE

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit
		default:
			return s, nil
		}
	}
	return s, nil
}
