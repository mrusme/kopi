package bagListCmd

import (
	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/formatter"
	"github.com/mrusme/kopi/helpers/out"
)

func mdList(entities []bagLabel.Label, fields []string) {
	out.Put(formatter.ListToMarkdown(entities, fields))
}
