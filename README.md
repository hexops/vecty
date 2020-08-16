<p align="center">
	<img src="https://github.com/vecty/vecty-logo/raw/master/horizontal_color_tagline.png" />
</p>

Vecty lets you build responsive and dynamic web frontends in Go using WebAssembly, competing with modern web frameworks like React & VueJS.

[![Build Status](https://travis-ci.org/hexops/vecty.svg?branch=master)](https://travis-ci.org/hexops/vecty) [![PkgGoDev](https://pkg.go.dev/badge/github.com/hexops/vecty)](https://pkg.go.dev/github.com/hexops/vecty) [![GoDoc](https://godoc.org/github.com/hexops/vecty?status.svg)](https://godoc.org/github.com/hexops/vecty) [![codecov](https://img.shields.io/codecov/c/github/hexops/vecty/master.svg)](https://codecov.io/gh/hexops/vecty)

Benefits
========

- Go developers can be competitive frontend developers.
- Share Go code between your frontend & backend.
- Reusability by sharing components via Go packages so that others can simply import them.

Goals
=====

- _Simple_
	- Designed from the ground up to be easily mastered _by newcomers_ (like Go).
- _Performant_
	- Efficient & understandable performance, small bundle sizes, same performance as raw JS/HTML/CSS.
- _Composable_
	- Nest components to form your entire user interface, seperating them logically as you would any normal Go package.
- _Designed for Go (implicit)_
	- Written from the ground up asking the question _"What is the best way to solve this problem in Go?"_, not simply asking _"How do we translate $POPULAR_LIBRARY to Go?"_

Features
========

- Compiles to WebAssembly (via standard Go compiler).
- Small bundle sizes: 0.5 MB hello world (see section below).
- Fast expectation-based browser DOM diffing ('virtual DOM', but less resource usage).

Current Status
==============

**Vecty is currently considered to be an experimental work-in-progress.** Prior to widespread production use, we must meet our [v1.0.0 milestone](https://github.com/hexops/vecty/issues?q=is%3Aopen+is%3Aissue+milestone%3A1.0.0) goals, which are being completed slowly and steadily as contributors have time (Vecty is over 4 years in the making!).

Early adopters may make use of it for real applications today as long as they are understanding and accepting of the fact that:

- APIs will change (maybe extensively).
- A number of important things are not ready:
	- Extensive documentation, examples and tutorials
	- URL-based component routing
	- Ready-to-use component libraries (e.g. material UI)
	- Server-side rendering
	- And more, see [milestone: v1.0.0 ](https://github.com/hexops/vecty/issues?q=is%3Aopen+is%3Aissue+milestone%3A1.0.0)
- The scope of Vecty is only ~80% defined currently.
- There are a number of important [open issues](https://github.com/hexops/vecty/issues).

For a list of projects currently using Vecty, see the [doc/projects-using-vecty.md](doc/projects-using-vecty.md) file.

Near-zero dependencies
======================

Vecty has nearly zero dependencies, it only relies on `reflect` from the Go stdlib. Through this, it is able to produce the smallest bundle sizes for Go frontend applications out there, limited only by the Go compiler itself:

| Example binary    | Compiler        | uncompressed | `gzip --best` | `brotli` |
|-------------------|-----------------|--------------|---------------|----------|
| `hellovecty.wasm` | `tinygo 0.14.0` | 252K         | 97K           | 81K      |
| `hellovecty.wasm` | `go 1.15`       | 2.1M         | 587K          | 443K     |
| `markdown.wasm`   | `go 1.15`       | 3.6M         | 1010K         | 745K     |
| `todomvc.wasm`    | `go 1.15`       | 2.9M         | 825K          | 617K     |

Note: Bundle sizes above do not scale linearly with more code/dependencies in your Vecty project. `hellovecty` is the smallest base-line bundle that the compiler can produce with just Vecty as a dependency, other examples above pull in more of the Go standard library and you would not e.g. have to pay that total cost again.

Community
=========

- Join us in the [#vecty](https://gophers.slack.com/messages/vecty/) channel on the [Gophers Slack](https://gophersinvite.herokuapp.com/)!

Changelog
=========

See the [doc/CHANGELOG.md](doc/CHANGELOG.md) file.
