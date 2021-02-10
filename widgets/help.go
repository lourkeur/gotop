package widgets

import (
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	lingo "github.com/xxxserxxx/lingo"
)

var tr lingo.Translations
var keyBinds string

type HelpMenu struct {
	widgets.Paragraph
}

func NewHelpMenu(tra lingo.Translations) (help *HelpMenu) {
	tr = tra
	keyBinds = tr.Value("help.help")
	help = &HelpMenu{
		Paragraph: *widgets.NewParagraph(),
	}
	help.Paragraph.Text = keyBinds
	return
}

func (help *HelpMenu) Resize(termWidth, termHeight int) {
	textWidth := 53
	for _, line := range strings.Split(keyBinds, "\n") {
		if textWidth < len(line) {
			textWidth = len(line) + 2
		}
	}
	textHeight := strings.Count(keyBinds, "\n") + 1
	x := (termWidth - textWidth) / 2
	y := (termHeight - textHeight) / 2

	help.Paragraph.SetRect(x, y, textWidth+x, textHeight+y)
}

func (help *HelpMenu) Draw(buf *ui.Buffer) {
	help.Paragraph.Draw(buf)
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
