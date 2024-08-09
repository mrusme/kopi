package cupListCmd

import (
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []cup.Cup, accessible bool) {
	for _, entity := range entities {
		out.Put("%s", entity.ID)
	}
}
