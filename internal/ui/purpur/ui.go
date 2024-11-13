package purpur

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eldius/bubble-pocs/internal/client/purpur"
	"github.com/erikgeiser/promptkit/confirmation"
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
	c := purpur.NewClient()
	m := newModel(c)
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
		fmt.Println(opts.String())

		cfbox := confirmation.New("\tDo you want to download it?", confirmation.NewValue(false))
		answer, err := cfbox.RunPrompt()
		if err != nil {
			return err
		}
		if answer {
			f, err := c.Download(opts.MineVer, opts.PurpurVer)
			if err != nil {
				err = fmt.Errorf("downloading file: %w", err)
				return err
			}

			fmt.Printf("output file: %s\n\n\n", f)
		}
	}
	return nil
}
