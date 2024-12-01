package purpur

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eldius/bubble-pocs/internal/client/purpur"
	"github.com/eldius/bubble-pocs/internal/service"
	"github.com/erikgeiser/promptkit/confirmation"
	"os"
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
			f, err := c.DownloadPurpur(opts.MineVer, opts.PurpurVer, ".")
			if err != nil {
				err = fmt.Errorf("downloading file: %w", err)
				return err
			}

			fmt.Printf("output file: %s\n\n\n", f)

			if err := startScript(f); err != nil {
				return err
			}
		}

	}
	return nil
}

func startScript(serverFile string) error {
	startScript, err := service.Generate(
		service.WithHeadless(true),
		service.WithMemLimit("2g"),
		service.WithJDKPath("java/jdk/bin"),
		service.WithServerFile(serverFile),
	)
	if err != nil {
		return fmt.Errorf("generating start script: %w", err)
	}

	f, err := os.OpenFile("start.sh", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("opening start.sh: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	if _, err := f.WriteString(startScript); err != nil {
		return fmt.Errorf("writing to start_script.sh: %w", err)
	}

	return nil
}
