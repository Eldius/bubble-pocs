package styled

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eldius/bubble-pocs/internal/client/mojang"
)

var (
	defaultStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#000000"))

	activeCursorStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#03fc03")).
				Background(lipgloss.Color("#7f8085"))
)

type usersModel struct {
	users  mojang.MojangUsers
	params []string
	c      *mojang.Mojang
	err    error
	curr   int
}

func newUsersModel(users ...string) *usersModel {
	return &usersModel{
		c:      mojang.NewMojang(),
		params: users,
	}
}

func (u usersModel) Init() tea.Cmd {
	return func() tea.Msg {
		users, err := u.c.FetchUsers(u.params...)
		if err != nil {
			return err
		}
		return users
	}
}

func (u usersModel) View() string {
	msg := "---\n"
	for i, user := range u.users {
		if i == u.curr {
			msg += activeCursorStyle.Render(fmt.Sprintf("-> %s\t\t(%s)", user.Name, user.ID)) + "\n"
		} else {
			msg += defaultStyle.Render(fmt.Sprintf("-  %s", user.Name)) + "\n"
		}
	}

	return msg + defaultStyle.Render("\n\nPress Ctrl+C or q to exit") + "\n"
}

func (u usersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		//fmt.Println(msg.(tea.KeyMsg).String())
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return u, tea.Quit
		case "q":
			return u, tea.Quit
		case "up":
			if u.curr > 0 {
				u.curr--
			}
			return u, nil
		case "down":
			if u.curr < len(u.users)-1 {
				u.curr++
			}
			return u, nil
		case "enter":
			return u, tea.Quit
		}
	case mojang.MojangUsers:
		u.users = msg.(mojang.MojangUsers)
		return u, nil
	}
	return u, nil
}

func Start(params ...string) {
	p := tea.NewProgram(
		newUsersModel(params...),
	)
	m, err := p.Run()
	if err != nil {
		panic(err)
	}

	if m, ok := m.(usersModel); ok {
		u := m.users[m.curr]
		fmt.Printf(" => selected user: %s (%s)\n\n", u.Name, u.ID)
	}
}
