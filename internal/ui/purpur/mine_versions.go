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

func (m *mineVersionsModel) SelectedVersion() string {
	return m.mineVersList[m.paginated.Page*m.paginated.PerPage+m.curr].ver
}

type mineVersionsModel struct {
	tea.Model
	msg  string
	err  error
	curr int
	c    *purpur.Client

	paginated    paginator.Model
	mineVersList []mineVer
}

func newModel() *mineVersionsModel {
	c := purpur.NewClient()
	p := paginator.New(
		paginator.WithPerPage(10),
		paginator.WithTotalPages(0),
	)
	return &mineVersionsModel{
		c:         c,
		curr:      0,
		msg:       "loading data...",
		paginated: p,
	}
}

func (m *mineVersionsModel) Init() tea.Cmd {
	return fetchMineVers(m)
}

func (m *mineVersionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.curr < (m.paginated.PerPage - 1) {
				m.curr = m.curr + 1
			}
		case "enter":
			return m, tea.Quit
		}

	case *purpur.GetMinecraftVersionsResponse:
		msg.SortVersions()
		var versions []mineVer
		for _, v := range msg.Versions {
			slog.With("ver", v).Debug("oneMoreVersion")
			versions = append(versions, mineVer{
				ver:    v,
				latest: v == msg.Metadata.Current,
			})
		}

		return updateVersions(m, msg)
	}
	var cmd tea.Cmd
	m.paginated, cmd = m.paginated.Update(msg)
	return m, cmd
}

func updateVersions(m *mineVersionsModel, msg *purpur.GetMinecraftVersionsResponse) (tea.Model, tea.Cmd) {
	var versions []mineVer
	for _, v := range msg.Versions {
		slog.With("ver", v).Debug("oneMoreVersion")
		versions = append(versions, mineVer{
			ver:    v,
			latest: v == msg.Metadata.Current,
		})
	}
	m.mineVersList = versions
	m.paginated.TotalPages = len(versions) / m.paginated.PerPage

	return m, nil
}

func (m *mineVersionsModel) View() string {
	if m.msg != "" {
		return m.msg
	}
	if m.err != nil {
		msg := m.err.Error()
		bar := "\t" + strings.Repeat("#", len(msg)+4) + "\t"
		msg = fmt.Sprintf("\t# %s #\t", msg)
		return fmt.Sprintf(`
%s
%s
%s

`, errorStyle.Render(bar), errorStyle.Render(msg), errorStyle.Render(bar))
	}
	var b strings.Builder
	b.WriteString("\n  Select mine version\n\n")
	start, end := m.paginated.GetSliceBounds(len(m.mineVersList))
	for i, item := range m.mineVersList[start:end] {
		if i == m.curr {
			b.WriteString(activeCurrStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		} else {
			b.WriteString(defaultStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		}
	}
	b.WriteString("  " + m.paginated.View())
	b.WriteString("\n\n  h/l ←/→ page • 'space' select/unselect mod • 'enter' to select • q: quit\n")
	return b.String()
}

func fetchMineVers(m *mineVersionsModel) tea.Cmd {
	return func() tea.Msg {
		res, err := m.c.GetMinecraftVesions()
		if err != nil {
			return err
		}
		return res
	}
}
