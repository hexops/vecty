# Vecty Tutorial

You will need to install gopherjs and vecty.

This tutorial assumes you have Go installed, and that you will follow along by
writing code in a directory somewhere on your GOPATH.

## Hello, World

This tutorial will aim to introduce you to basic Vecty concepts. To begin, we 
will render a simple web page. Create a file **main.go** and paste the following
contents. 

```go
package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.SetTitle("Vecty Tutorial")
	c := &MyComponent{}
	vecty.RenderBody(c)
}

type MyComponent struct {
	vecty.Core
}

func (c *MyComponent) Render() *vecty.HTML {
	return elem.Body(
		elem.Heading1(vecty.Text("Hello, World")),
	)
}
```

Since Vecty is built on gopherjs, we can view it right away with the gopherjs 
development server. Open a terminal and run `gopherjs server`, and navigate to 
the appropriate URL. The gopherjs server will render and serve this file 
directly from your **GOPATH**. For instance, if you saved **main.go** here:

```
$GOPATH/src/vecty-example/main.go
```

The development server should serve this page at http://localhost:8080/vecty-example

This code is very simple. We declare one custom Component, and this component
embeds the `Body`, `Heading1`, and `Text` elements built in to Vecty. Using 
Components let's us create high-level, reusable containers for our UI. 

At a minimum, a valid Component must be a `struct` that

* declares a `Render` function
* embeds `vecty.Core`

The `Render` function is where we define the look and feel of our component. The
`vecty.Core` bit uses Go's struct embedding to help our component implement the
`vecty.Component` interface.

So far, we haven't written any HTML. All we have done is write our component and
pass it to Vecty's `RenderBody` function. This function is special: you **must**
pass it a Component that returns an `elem.Body`. Fortunately, our code satisfies
this.

## More Complex Example

_TODO_

* use some custom, arbitrary "data" fields in our component
* add some basic styling
* render a custom component in a custom component
* model the example off of go-kit's stringsvc example?

