package equipmentListCmd

import (
	"xn--gckvb8fzb.com/kopi/equipment"
	"xn--gckvb8fzb.com/kopi/helpers/formatter"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

func mdList(entities []equipment.Equipment, fields []string) {
	out.Put(formatter.ListToMarkdown(entities, fields))
}
