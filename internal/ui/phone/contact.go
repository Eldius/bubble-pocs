package phone

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	//cursorStyle  = focusedStyle
	//noStyle      = lipgloss.NewStyle()
)

type contactModel struct {
	tea.Model
	name  textinput.Model
	email textinput.Model
	curr  int
	//currMode cursor.Mode
}

func newContactModel() *contactModel {
	name := textinput.New()
	name.Prompt = "Name"
	name.Placeholder = "My Name"
	name.Cursor.Style = focusedStyle
	email := textinput.New()
	email.Prompt = "Email"
	email.Placeholder = "me@me.com"
	email.Cursor.Style = blurredStyle

	return &contactModel{
		name:  name,
		email: email,
	}
}

func (c *contactModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (c *contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return c, tea.Quit
		}
	}
	return c, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (c *contactModel) View() string {
	if c.curr == 0 {
		c.name.Cursor.Style = focusedStyle
	} else {
		c.name.Cursor.Style = blurredStyle
	}

	if c.curr == 1 {
		c.email.Cursor.Style = focusedStyle
	} else {
		c.email.Cursor.Style = blurredStyle
	}

	var b strings.Builder

	b.WriteString(c.name.View() + "\n")
	b.WriteString(c.email.View() + "\n")
	return b.String()
}

func (c *contactModel) String() string {
	return fmt.Sprintf(`----------------------------
Name:   %s
Email:  %s
******************************`, c.name.Value(), c.email.Value())
}
