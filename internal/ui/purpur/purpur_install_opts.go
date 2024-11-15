package purpur

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eldius/bubble-pocs/internal/client/purpur"
	"github.com/eldius/bubble-pocs/internal/ui/utils"
	"log/slog"
	"strings"
)

const (
	mineVersionScreen   = "mine"
	purpurVersionScreen = "purpur"
)

type VersionInfo struct {
	ver    string
	latest bool
}

type InstallOpts struct {
	MineVer   string
	PurpurVer string
}

func (o InstallOpts) String() string {
	msgMineVer := "Minecraft selected version: " + o.MineVer

	msgPurpurVer := "Purpur selected version:  " + o.PurpurVer

	mineMsgLength := len(msgMineVer)
	purpurMsgLength := len(msgPurpurVer)

	boxSize := mineMsgLength
	if boxSize < purpurMsgLength {
		boxSize = purpurMsgLength
	}

	msgMineVer = fmt.Sprintf("\t# %s #\t", msgMineVer+strings.Repeat(" ", boxSize-mineMsgLength))
	msgPurpurVer = fmt.Sprintf("\t# %s #\t", msgPurpurVer+strings.Repeat(" ", boxSize-purpurMsgLength))

	bar := "\t" + strings.Repeat("#", boxSize+4) + "\t"

	emptyLine := strings.Repeat(" ", boxSize)
	return fmt.Sprintf(`
%s
%s
%s
%s
%s
%s
`, emptyLine, defaultStyle.Render(bar), defaultStyle.Render(msgMineVer), defaultStyle.Render(msgPurpurVer), defaultStyle.Render(bar), emptyLine)

}

func (m *purpurVersionsModel) selectedMineVersion() string {
	return m.mineVerList[m.mineVerPaginatedModel.Page*m.mineVerPaginatedModel.PerPage+m.curr].ver
}

func (m *purpurVersionsModel) selectedPurpurVersion() string {
	return m.purpurVerList[m.mineVerPaginatedModel.Page*m.mineVerPaginatedModel.PerPage+m.curr].ver
}

func (m *purpurVersionsModel) GetInstallOpts() InstallOpts {
	return InstallOpts{
		MineVer:   m.mineVer,
		PurpurVer: m.purpurVer,
	}
}

type purpurVersionsModel struct {
	tea.Model
	msg  string
	err  error
	curr int
	c    *purpur.Client

	mineVer               string
	mineVerPaginatedModel paginator.Model
	mineVerList           []VersionInfo

	purpurVer               string
	purpurVerPaginatedModel paginator.Model
	purpurVerList           []VersionInfo

	screen string
}

func (m *purpurVersionsModel) isMineVersionScreen() bool {
	return m.screen == mineVersionScreen
}

func (m *purpurVersionsModel) isPurpurVersionScreen() bool {
	return m.screen == purpurVersionScreen
}

func newModel(c *purpur.Client) *purpurVersionsModel {
	pm := paginator.New(
		paginator.WithPerPage(10),
		paginator.WithTotalPages(0),
	)
	pp := paginator.New(
		paginator.WithPerPage(10),
		paginator.WithTotalPages(0),
	)
	return &purpurVersionsModel{
		c:                       c,
		curr:                    0,
		msg:                     "loading data...",
		mineVerPaginatedModel:   pm,
		purpurVerPaginatedModel: pp,
		screen:                  mineVersionScreen,
	}
}

func (m *purpurVersionsModel) Init() tea.Cmd {
	return fetchMineVers(m)
}

func (m *purpurVersionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msg = ""
	switch msg := msg.(type) {
	case string:
		m.msg = msg
		return m, nil
	case error:
		m.err = msg
		return m, nil
	}
	if m.isMineVersionScreen() {
		return m.updateMineVerScreen(msg)
	}

	if m.isPurpurVersionScreen() {
		return m.updatePurpurVerScreen(msg)
	}
	return m, nil
}

func (m *purpurVersionsModel) updateMineVerScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.curr < (m.mineVerPaginatedModel.PerPage - 1) {
				m.curr = m.curr + 1
			}
		case "enter":
			m.mineVer = m.selectedMineVersion()
			//m.screen = purpurVersionScreen
			return m, fetchPurpurVers(m)
		}

	case *purpur.GetMinecraftVersionsResponse:
		msg.SortVersions()
		return setMineVersionsList(m, msg)
	}
	var cmd tea.Cmd
	m.mineVerPaginatedModel, cmd = m.mineVerPaginatedModel.Update(msg)
	return m, cmd
}

func setMineVersionsList(m *purpurVersionsModel, msg *purpur.GetMinecraftVersionsResponse) (tea.Model, tea.Cmd) {
	var versions []VersionInfo
	for _, v := range msg.Versions {
		versions = append(versions, VersionInfo{
			ver:    v,
			latest: v == msg.Metadata.Current,
		})
	}
	verCount := len(versions)
	if verCount < m.mineVerPaginatedModel.PerPage {
		m.mineVerPaginatedModel.PerPage = verCount
	}

	m.mineVerList = versions
	m.mineVerPaginatedModel.TotalPages = len(versions) / m.mineVerPaginatedModel.PerPage

	return m, nil
}

func (m *purpurVersionsModel) updatePurpurVerScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.curr < (m.purpurVerPaginatedModel.PerPage - 1) {
				m.curr = m.curr + 1
			}
		case "enter":
			m.purpurVer = m.selectedPurpurVersion()
			return m, tea.Quit
		}

	case *purpur.GetPurpurVersionsResponse:
		m.curr = 0
		msg.SortVersions()
		return setPurpurVersionsList(m, msg)
	}
	var cmd tea.Cmd
	m.purpurVerPaginatedModel, cmd = m.purpurVerPaginatedModel.Update(msg)
	return m, cmd
}

func setPurpurVersionsList(m *purpurVersionsModel, msg *purpur.GetPurpurVersionsResponse) (tea.Model, tea.Cmd) {
	var versions []VersionInfo
	for _, v := range msg.Builds.All {
		slog.With("purpur_ver", v).Debug("oneMoreVer")
		versions = append(versions, VersionInfo{
			ver:    v,
			latest: v == msg.Builds.Latest,
		})
	}
	m.purpurVerList = versions
	verCount := len(versions)
	if verCount < m.purpurVerPaginatedModel.PerPage {
		m.purpurVerPaginatedModel.PerPage = verCount
	}
	m.purpurVerPaginatedModel.TotalPages = verCount / m.purpurVerPaginatedModel.PerPage

	return m, nil
}

func (m *purpurVersionsModel) View() string {
	if m.msg != "" {
		return utils.DisplayMessage(m.msg, defaultStyle)
	}
	if m.err != nil {
		return utils.DisplayMessage(m.err.Error(), errorStyle)
	}

	if m.isMineVersionScreen() {
		return m.viewMineVerScreen()
	}

	if m.isPurpurVersionScreen() {
		return m.viewPurpurVerScreen()
	}

	return ""
}

func (m *purpurVersionsModel) viewMineVerScreen() string {
	var b strings.Builder
	b.WriteString("\n  Select mine version\n\n")
	start, end := m.mineVerPaginatedModel.GetSliceBounds(len(m.mineVerList))
	for i, item := range m.mineVerList[start:end] {
		if i == m.curr {
			b.WriteString(activeCurrStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		} else {
			b.WriteString(defaultStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		}
	}
	b.WriteString("  " + m.mineVerPaginatedModel.View())
	b.WriteString("\n\n  h/l ←/→ page • • 'enter' to select version • q: quit\n")

	return b.String()
}

func (m *purpurVersionsModel) viewPurpurVerScreen() string {
	var b strings.Builder
	b.WriteString("\n  Select mine version\n\n")
	start, end := m.purpurVerPaginatedModel.GetSliceBounds(len(m.mineVerList))
	for i, item := range m.purpurVerList[start:end] {
		if i == m.curr {
			b.WriteString(activeCurrStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		} else {
			b.WriteString(defaultStyle.Render(fmt.Sprintf("- %s (latest: %v)", item.ver, item.latest)) + "\n")
		}
	}
	b.WriteString("  " + m.purpurVerPaginatedModel.View())
	b.WriteString("\n\n  h/l ←/→ page • • 'enter' to select version • q: quit\n")

	return b.String()
}

func fetchMineVers(m *purpurVersionsModel) tea.Cmd {
	return func() tea.Msg {
		res, err := m.c.GetPurpurMinecraftVesions()
		if err != nil {
			return err
		}
		return res
	}
}

func fetchPurpurVers(m *purpurVersionsModel) tea.Cmd {
	return func() tea.Msg {
		ver := m.mineVer
		if ver == "" {
			return errors.New("could not fetch purpur version without a mine version")
		}
		res, err := m.c.GetPurpurBuildsByMineVersion(ver)
		if err != nil {
			return err
		}
		m.screen = purpurVersionScreen
		return res
	}
}
