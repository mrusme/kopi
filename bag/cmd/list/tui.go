package bagListCmd

import (
	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []bagLabel.Label, accessible bool) {
	for _, entity := range entities {
		out.Put("%s", entity.Name)
	}
}
