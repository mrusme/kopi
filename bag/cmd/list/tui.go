package bagListCmd

import (
	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/out"
)

func tuiList(labels []bagLabel.Label, accessible bool) {
	for _, label := range labels {
		out.Put("%s", label.Name)
	}
}
