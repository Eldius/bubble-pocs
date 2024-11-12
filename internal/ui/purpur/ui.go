package purpur

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
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
