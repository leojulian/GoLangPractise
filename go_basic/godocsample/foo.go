// Copyright 2021 The Go Authers. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package foo implements a ser of simple mathematical functions. These comments are for
demonstration purpose only. Nothing more.
*/
package foo

func foo(a, b int) (ret int, err error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
	return 0, nil
}

// BUG(leo): #1: I'm sorry but this code has an issue to be solved
// BUG(jack): #2: An issue assigned to another person.
