package out

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Put(s string, v ...any) {
	fmt.Printf(s+"\n", v...)
}

func Pit(s string, v ...any) {
	fmt.Printf(s, v...)
}

func Debug(s string, v ...any) {
	if viper.GetBool("Debug") {
		fmt.Fprintf(os.Stderr, s+"\n", v...)
	}
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

func NilOrErr(err error, v ...string) bool {
	if err != nil {
		if len(v) == 1 {
			Err("%s: %s", v[0], err)
		} else {
			Err("%s", err)
		}
		return true
	}

	return false
}
