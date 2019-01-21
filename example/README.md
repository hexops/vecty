# WebAssembly

## WebAssembly: Go 1.12+ requirement

Vecty makes use of Go [synchronous callback support](https://go-review.googlesource.com/c/go/+/142004) which is only present in Go 1.12+.

**Go 1.12 is not yet released, so you must use Go tip until then**

Mac:

```
brew unlink go
brew install go --HEAD
```

## WebAssembly: Running examples

The easiest way to run the examples as WebAssembly is via [`wasmserve`](https://github.com/hajimehoshi/wasmserve).

Install it (**using Go 1.12+**):

```bash
go get -u github.com/hajimehoshi/wasmserve
```

Then run an example:

```bash
cd markdown/
wasm-server
```

And navigate to http://localhost:8080/
