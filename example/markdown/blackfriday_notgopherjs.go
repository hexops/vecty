// +build wasm

// Please ignore this file it is not a useful part of the example. It works
// around some nasty import path stuff where GopherJS does not have
// module-aware build support (https://github.com/gopherjs/gopherjs/issues/855)
// and as such blackfriday must be imported under a different import path
// there. In practice this won't be something you face in your application
// since you'll pick just one build system, but we support both here.
//
// Note: This is tricky due to the overlap in GOARCH/GOOS on each target:
//
// GopherJS: GOARCH=js GOOS=<host OS>
// WebAssembly: GOOS=js GOARCH=wasm

package main

import "github.com/russross/blackfriday/v2"

var blackfridayRun = blackfriday.Run