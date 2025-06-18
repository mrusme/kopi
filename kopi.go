package main

import (
	"fmt"

	"github.com/mrusme/kopi/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Execute(version)
}
