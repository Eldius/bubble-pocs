package utils

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func DisplayMessage(msg string, style lipgloss.Style) string {
	bar := "\t" + strings.Repeat("#", len(msg)+4) + "\t"
	msg = fmt.Sprintf("\t# %s #\t", msg)
	return fmt.Sprintf(`
%s
%s
%s

`, style.Render(bar), style.Render(msg), style.Render(bar))
}
