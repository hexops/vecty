# Components

There are two ways to define something that will render on a page.

1. Create a `struct` that embeds `vecty.Core` and implements `Render`
2. Write a function with the signature `func(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML`

## Component Structs

Component structs look like this

```go
type SomeComponent struct {
    vecty.Core
}

func (sc *SomeComponent) Render() *vecty.HTML {
    return elem.Body(vecty.Text("some text in the body"))
}
```

Component structs are regular Go structs that adhere to the `vecty.Component` 
interface. 

```go
type Component interface {
	Render() *HTML
	Context() *Core
}
```

If you need some background on Go interfaces, please read [Effective Go](https://golang.org/doc/effective_go.html#interfaces_and_types).
Note that embedding `vecty.Core` gives your components the `Context` method,
and it is up to you to provide the `Render` method. The `Render` method must
be implemented on a pointer receiver.

Beyond implementing the `Component` interface, component structs are just regular
Go structs. We can add custom data fields and make use of them during rendering.
Here we add a string field for our message.

```go
type SomeComponent2 struct {
    vecty.Core
    Message string
}

func (sc *SomeComponent2) Render() *vecty.HTML {
    return elem.Body(vecty.Text(sc.Message))
}
```

## Component Functions

If you do not need a component to maintain state, it is possible to write
regular functions that will render to the page. These functions have this
signature

```go
func(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML
```

We can define our own like this

```go
func NavWrapper(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	markup = append(markup, prop.Class("nav"))
	return elem.Div(markup...)
}
```

All we are doing is embedding a CSS class in our list of markup items that
will be **applied** by Vecty. Since this function returns `*vecty.HTML`, 
we can compose it into another component's `Render` function.

```go
// updated

func (sc *SomeComponent2) Render() *vecty.HTML {
    return elem.Body(
        NavWrapper(
            elem.UnorderedList(
                elem.ListItem(vecty.Text("First")),
                elem.ListItem(vecty.Text("Second")),
                elem.ListItem(vecty.Text("Third")),
            )
        )
    )
}
```

Note that this is the same signature used by the HTML helper functions 
in `package elem`. 
