package mojang

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eldius/bubble-pocs/internal/client/mojang"
)

type usersModel struct {
	users  mojang.MojangUsers
	params []string
	c      *mojang.Mojang
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
			msg += fmt.Sprintf("-> %s\t\t(%s)\n", user.Name, user.ID)
		} else {
			msg += fmt.Sprintf("-  %s\n", user.Name)
		}

	}

	return msg + "\n\nPress Ctrl+C or q to exit\n"
}

func (u usersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
		u.users = msg
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
