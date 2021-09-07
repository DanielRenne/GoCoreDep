// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bettermd5_test

import (
	"fmt"
	"github.com/koofr/go-cryptoutils/bettermd5"
	"io"
)

func ExampleNew() {
	h := bettermd5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
	// Output: e2c569be17396eca2a2e3c11578123ed
}

func ExampleNewFromState() {
	h := bettermd5.New()
	io.WriteString(h, "The fog is getting thicker!")
	h.Sum(nil)
	h1 := bettermd5.NewFromState(h.GetState())
	io.WriteString(h1, "And Leon's getting laaarger!")
	fmt.Printf("%x", h1.Sum(nil))
	// Output: e2c569be17396eca2a2e3c11578123ed
}

func ExampleSum() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", bettermd5.Sum(data))
	// Output: b0804ec967f48520697662a204f5fe72
}
