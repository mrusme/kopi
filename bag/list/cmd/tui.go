package bagListCmd

import (
	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(entities []bagLabel.Label, fields []string) {
	out.Put(formatter.ListToTUI(entities, fields))
}
