package bagListCmd

import (
	bagLabel "xn--gckvb8fzb.com/kopi/bag/label"
	"xn--gckvb8fzb.com/kopi/helpers/formatter"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func mdList(entities []bagLabel.Label, fields []string) {
	out.Put(formatter.ListToMarkdown(entities, fields))
}
