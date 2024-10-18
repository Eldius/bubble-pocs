package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(
		newSimplePage("This app is under construction"),
	)
	if m, err := p.Run(); err != nil {
		panic(err)
	} else {
		fmt.Printf("m: %+v\n", m)
	}
}
