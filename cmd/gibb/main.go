// Command gibberize goes through all files in the current directory,
// and replaces their contents with meaningless text.
package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	force   = flag.BoolP("f", "force", false, "overwrite existing files")
	recurse = flag.BoolP("r", "recurse", false, "recurse into directories")

	mean      = flag.BoolP("m", "mean", 10, "mean number of lines")
	deviation = flag.BoolP("s", "deviation", 5, "standard deviation")
)

func help() {
	fmt.Println(`gibberize version 0.1

Gibberize creates files filled with gibberish for test purposes.
`)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		help()
		os.Exit(1)
	}

}
