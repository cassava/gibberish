// Copyright (c) 2015, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Command gibberize goes through all files in the current directory,
// and replaces their contents with meaningless text.
package main

import (
	"fmt"

	"github.com/cassava/gibberish"
	flag "github.com/ogier/pflag"
)

var (
	force   = flag.BoolP("force", "f", false, "overwrite existing files")
	recurse = flag.BoolP("recurse", "r", false, "recurse into directories")
)

func help() {
	fmt.Println(`gibberize version 0.1

Gibberize creates files filled with gibberish for test purposes.
`)
}

func main() {
	flag.Parse()

	g := gibberish.NewGenerator()
	fmt.Println(g.Document())
}
