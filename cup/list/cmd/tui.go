package cupListCmd

import (
	"xn--gckvb8fzb.com/kopi/cup"
	"xn--gckvb8fzb.com/kopi/helpers/formatter"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func tuiList(entities []cup.Cup, fields []string) {
	out.Put(formatter.ListToTUI(entities, fields))
}
