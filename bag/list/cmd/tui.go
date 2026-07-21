package bagListCmd

import (
	bagLabel "xn--gckvb8fzb.com/kopi/bag/label"
	"xn--gckvb8fzb.com/kopi/helpers/formatter"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func tuiList(entities []bagLabel.Label, fields []string) {
	out.Put(formatter.ListToTUI(entities, fields))
}
