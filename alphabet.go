// Copyright (c) 2015, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package gibberish

type Alphabet struct {
	Upper string
	Lower string
	Punct string
}

var ASCII = &Alphabet{
	Upper: "ABBCCDDEFFGGHHIJJKKLLMMNNOPPQRRSSTTUVWWXYZ",
	Lower: "aaaaabcdeeeeefghiiiiijklmnoooopqrstuuuuuvwxyz",
	Punct: ",,,,,,,,;;;:",
}
