package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eldius/bubble-pocs/internal/client/mojang"
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
			msg += fmt.Sprintf("-> %s\n", user.Name)
		} else {
			msg += fmt.Sprintf("-  %s\n", user.Name)
		}

	}

	return msg + "\n\nPress Ctrl+C or q to exit\n"
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
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
