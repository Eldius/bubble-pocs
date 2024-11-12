package purpur

import (
	"fmt"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eldius/bubble-pocs/internal/client/purpur"
	"log/slog"
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

type mineVer struct {
	ver    string
	latest bool
}

type purpurModel struct {
	tea.Model
	msg       string
	err       error
	curr      int
	c         *purpur.Client
	mineVer   string
	purpurVer string

	mineVersPaginated paginator.Model
	mineVersionsList  []mineVer
}

func newModel() *purpurModel {
	c := purpur.NewClient()
	return &purpurModel{
		c:    c,
		curr: 0,
		msg:  "loading data...",
	}
}

func Start() error {
	m := newModel()
	_, err := tea.NewProgram(m).Run()
	if err != nil {
		return err
	}
	return nil
}

func (m *purpurModel) Init() tea.Cmd {
	return fetchMineVers(m)
}

func (m *purpurModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msg = ""
	switch msg := msg.(type) {
	case string:
		m.msg = msg
	case error:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.curr > 0 {
				m.curr = m.curr - 1
			}
		case "down", "j":
			if m.curr < (len(m.mineVersionsList) - 1) {
				m.curr = m.curr + 1
			}
		}
	case *purpur.GetMinecraftVersionsResponse:
		var versions []mineVer
		for _, v := range msg.Versions {
			slog.With("ver", v).Debug("oneMoreVersion")
			versions = append(versions, mineVer{
				ver:    v,
				latest: v == msg.Metadata.Current,
			})
		}
		m.mineVersionsList = versions
	}
	return m, nil
}

//func updateVersions(m *purpurModel, msg *purpur.GetMinecraftVersionsResponse) (tea.Model, tea.Cmd) {
//	var versions []mineVer
//	for _, v := range msg.Versions {
//		slog.With("ver", v).Debug("oneMoreVersion")
//		versions = append(versions, mineVer{
//			ver:    v,
//			latest: v == msg.Metadata.Current,
//		})
//	}
//	m.mineVersionsList = versions
//	m.mineVersPaginated = paginator.New()
//	m.mineVersPaginated.ItemsOnPage()
//
//	return m, nil
//}

func (m *purpurModel) View() string {
	if m.msg != "" {
		return m.msg
	}
	if m.err != nil {
		msg := strings.TrimSpace(m.err.Error())
		bar := "\t" + strings.Repeat("#", len(msg)+4) + "\t"
		return errorStyle.Render(bar) + "\n" +
			errorStyle.Render(fmt.Sprintf("\t# %s #\t", msg)) + "\n" +
			errorStyle.Render(bar) + "\n"
	}
	msg := fmt.Sprintf("---\nlist size: %d\n", len(m.mineVersionsList))
	for i, v := range m.mineVersionsList {
		if i == m.curr {
			msg += activeCurrStyle.Render(fmt.Sprintf("-> %s\t\t(latest: %v)", v.ver, v.latest)) + "\n"
		} else {
			msg += defaultStyle.Render(fmt.Sprintf("-  %s\t\t(latest: %v)", v.ver, v.latest)) + "\n"
		}
	}
	return msg
}

func fetchMineVers(m *purpurModel) tea.Cmd {
	return func() tea.Msg {
		res, err := m.c.GetMinecraftVesions()
		if err != nil {
			return err
		}
		return res
	}
}
