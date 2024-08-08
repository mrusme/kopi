package out

import (
	"fmt"
	"os"
)

func Put(s string, v ...any) {
	fmt.Printf(s+"\n", v...)
}

func Err(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}

func Die(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
	os.Exit(1)
}

func NilOrDie(err error, v ...string) {
	if err != nil {
		if len(v) == 1 {
			Die("%s: %s", v[0], err)
		} else {
			Die("%s", err)
		}
	}
}
