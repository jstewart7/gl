// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gl is a Go cross-platform binding for OpenGL, with an OpenGL ES 2-like API.

It currently supports OS X, Linux via OpenGL ES 2 backend, and all modern desktop
and mobile browsers via WebGL 1 backend. Windows support is upcoming.

This is a fork of golang.org/x/mobile/gl/... packages with [CL 8793](https://go-review.googlesource.com/8793)
merged in. This package may change as that CL is reviewed, and hopefully eventually deleted once
the CL is merged and golang.org/x/mobile/gl/... can be used.
*/
package gl // import "github.com/goxjs/gl"

//go:generate go run gendebug.go -o gldebug.go
