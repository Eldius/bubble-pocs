package purpur

import (
	"fmt"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eldius/bubble-pocs/internal/client/purpur"
	"log/slog"
	"strings"
)

type purpurVersionsModel struct {
	tea.Model
	msg  string
	err  error
	curr int
	c    *purpur.Client

	mineVer   string
	paginated paginator.Model
	versList  []VersionInfo
}

func newPurpurVersionsModel(c *purpur.Client, mineVer string) (*purpurVersionsModel, tea.Cmd) {
	p := paginator.New(
		paginator.WithPerPage(10),
		paginator.WithTotalPages(0),
	)
	slog.Debug("newPurpurVersionsModel")
	m := &purpurVersionsModel{
		c:         c,
		mineVer:   mineVer,
		paginated: p,
	}
	return m, nil
}

func (m *purpurVersionsModel) Init() tea.Cmd {
	return fetchPurpurVersion(m.c, m.mineVer)
}

func (m *purpurVersionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *purpur.GetPurpurVersionsResponse:
		var versions []VersionInfo
		for _, v := range msg.Builds.All {
			versions = append(versions, VersionInfo{
				ver:    v,
				latest: v == msg.Builds.Latest,
			})
		}
		m.versList = versions
		m.paginated.TotalPages = len(versions) / m.paginated.PerPage

	case tea.KeyMsg:

	}
	return m, nil
}

func (m *purpurVersionsModel) View() string {
	if m.msg != "" {
		return m.msg
	}
	if m.err != nil {
		return displayErrorDetails(m.err)
	}
	var b strings.Builder
	b.WriteString("\n  Select Purpur build\n\n")
	start, end := m.paginated.GetSliceBounds(len(m.versList))
	for i, item := range m.versList[start:end] {
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

func fetchPurpurVersion(c *purpur.Client, mineVer string) tea.Cmd {
	return func() tea.Msg {
		res, err := c.GetBuildsByMineVersion(mineVer)
		if err != nil {
			return err
		}
		res.SortVersions()
		return res
	}
}
