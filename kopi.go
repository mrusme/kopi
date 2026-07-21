package main

import (
	"xn--gckvb8fzb.com/kopi/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Execute(version)
}
