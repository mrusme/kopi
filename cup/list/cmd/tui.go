package cupListCmd

import (
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []cup.Cup, fields []string) {
	out.Put(formatter.ListToTUI(entities, fields))
}
