# Building examples

Vecty fully supports two major Go <-> Web compilers thanks to its minimal dependencies:

- Go 1.14+ WebAssembly support
- [GopherJS](https://github.com/gopherjs/gopherjs) (Go to JavaScript transpiler)

If you are just getting started, we suggest using the Go 1.14+ WebAssembly support.

## Go 1.14+ WebAssembly support

**Ensure you are running Go 1.14 or higher.** Vecty requires Go 1.14+ as it makes use of improvements to the `syscall/js` package which are not present in earlier versions of Go.

### Running examples

The easiest way to run the examples as WebAssembly is via [`wasmserve`](https://github.com/hajimehoshi/wasmserve).

Install it (**using Go 1.14+**):

```bash
go get -u github.com/hajimehoshi/wasmserve
```

Then run an example:

```bash
cd markdown/
wasmserve
```

And navigate to http://localhost:8080/

# GopherJS

### Running examples

Install [GopherJS](https://github.com/gopherjs/gopherjs#installation-and-usage) then run an example:

```bash
cd markdown
gopherjs serve
```

And navigate to http://localhost:8080/

## TinyGo WebAssembly support

[TinyGo](https://github.com/tinygo-org/tinygo) WebAssembly support is [being actively worked on](https://github.com/tinygo-org/tinygo/issues/93).
