package cupListCmd

import (
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func mdList(entities []cup.Cup, fields []string) {
	out.Put(formatter.ListToMarkdown(entities, fields))
}
