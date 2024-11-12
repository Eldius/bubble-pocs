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
	if m, ok := m1.(*purpurVersionsModel); ok {
		if m.err != nil {
			return err
		}
		opts := m.GetInstallOpts()
		if opts.MineVer == "" {
			fmt.Println("\n", defaultStyle.Render("no Minecraft version selected"))
			fmt.Println()
			return nil
		}
		if opts.PurpurVer == "" {
			fmt.Println("\n", defaultStyle.Render("no Purpur version selected"))
			fmt.Println()
			return nil
		}
		msgMineVer := "Minecraft selected version: " + opts.MineVer

		msgPurpurVer := "Purpur selected version:  " + opts.PurpurVer

		mineMsgLength := len(msgMineVer)
		purpurMsgLength := len(msgPurpurVer)

		boxSize := mineMsgLength
		if boxSize < purpurMsgLength {
			boxSize = purpurMsgLength
		}

		msgMineVer = fmt.Sprintf("\t# %s #\t", msgMineVer+strings.Repeat(" ", boxSize-mineMsgLength))
		msgPurpurVer = fmt.Sprintf("\t# %s #\t", msgPurpurVer+strings.Repeat(" ", boxSize-purpurMsgLength))

		bar := "\t" + strings.Repeat("#", boxSize+4) + "\t"
		fmt.Println("\n",
			defaultStyle.Render(bar), "\n",
			defaultStyle.Render(msgMineVer), "\n",
			defaultStyle.Render(msgPurpurVer), "\n",
			defaultStyle.Render(bar))
		fmt.Println()
		fmt.Println()
	}
	return nil
}
