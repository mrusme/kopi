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
