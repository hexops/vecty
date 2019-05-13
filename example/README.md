# Building examples

Vecty fully supports all three major Go <-> Web compilers thanks to its minimal dependencies:

- Go 1.12+ WebAssembly support
- [GopherJS](https://github.com/gopherjs/gopherjs) (Go to JavaScript transpiler)
- [TinyGo](https://github.com/tinygo-org/tinygo) WebAssembly support (produces tiny binaries, but with tradeoffs)

If you are just getting started, we suggest using the Go 1.12+ WebAssembly support.


## Go 1.12+ WebAssembly support

Vecty requires Go 1.12+ as it makes use of [synchronous callback support](https://go-review.googlesource.com/c/go/+/142004) which is not present in earlier versions of Go. **Ensure you are running Go 1.12 or higher.**

### Running examples

The easiest way to run the examples as WebAssembly is via [`wasmserve`](https://github.com/hajimehoshi/wasmserve).

Install it (**using Go 1.12+**):

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

(Not recommended for beginners.)

TODO(slimsag)
