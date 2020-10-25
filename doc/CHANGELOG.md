Changelog
=========

Although v1.0.0 [is not yet out](https://github.com/hexops/vecty/milestone/1), we do not expect many breaking changes. When there is one, however, it is documented clearly here.

Pre-v1.0.0 Breaking Changes
---------------------------

## October 25, 2020

* The `master` branch has been renamed to `main`.
* `v0.6.0` has been tagged/released.

## August 15, 2020 ([PR #266](https://github.com/hexops/vecty/pull/266)): minor breaking change

Vecty has moved to the [github.com/hexops](https://github.com/hexops) organization. Update your import paths:

```diff
-import "github.com/gopherjs/vecty"
+import "github.com/hexops/vecty"
```

And update your `go.mod` as required.

For more information see [issue #230](https://github.com/hexops/vecty/issues/230#issuecomment-674474753).

## August 15, 2020 ([PR #265](https://github.com/hexops/vecty/pull/265)): minor breaking change

Deprecated and removed official support for GopherJS.

New versions of Vecty _may_ compile with GopherJS, but we do not officially support it and it is dependent on GopherJS being compatible with the official Go compiler.

If your application cannot compile without GopherJS, you may continue to use the tag `v0.5.0` which is the last version of Vecty which officially supported GopherJS at the time.

For more information please see [issue #264](https://github.com/hexops/vecty/issues/264).

## February 28, 2020 ([PR #256](https://github.com/hexops/vecty/pull/256)): indirect breaking change

- Go 1.14+ is now required by Vecty. Users of older Go versions and/or GopherJS (until https://github.com/gopherjs/gopherjs/issues/962 is fixed) may wish to continue using commit `6a0a25ee5a96ce029e684c7da6333aa1f34f8f96`.

## Nov 30, 2019 ([PR #249](https://github.com/hexops/vecty/pull/249)): minor breaking change

- `vecty.RenderBody(comp)` is now a blocking function call. Users that rely on it being non-blocking can instead now use `if err := vecty.RenderInto("body", comp); err != nil { panic(err) }`

## June 30, 2019 ([PR #232](https://github.com/hexops/vecty/pull/232)): major breaking change

- `(*HTML).Node` now returns a `syscall/js.Value` instead of `*gopherjs/js.Object`. Users will need to update to the new `syscall/js` API in their applications.
- Go 1.12+ is now required by Vecty, as we make use of [synchronous callback support](https://go-review.googlesource.com/c/go/+/142004) not present in earlier versions.

## May 25, 2019 ([PR #235](https://github.com/hexops/vecty/pull/235)): minor breaking change

- `prop.TypeUrl` has been renamed to `prop.TypeURL`.

## Nov 4, 2017 ([PR #158](https://github.com/hexops/vecty/pull/158)): major breaking change

All `Component`s must now have a `Render` method which returns `vecty.ComponentOrHTML` instead of the prior `*vecty.HTML` type.

This change allows for higher order components (components that themselves render components), which is useful for many more advanced uses of Vecty.

### Upgrading

Upgrading most codebases should be trivial with a find-and-replace across all files.

From your editor:
* Find `) Render() *vecty.HTML` and replace with `) Render() vecty.ComponentOrHTML`.

From the __Linux__ command line:
```bash
git grep -l ') Render() \*vecty.HTML' | xargs sed -i 's/) Render() \*vecty.HTML/) Render() vecty.ComponentOrHTML/g'
```

From the __Mac__ command line:
```bash
git grep -l ') Render() \*vecty.HTML' | xargs sed -i '' -e 's/) Render() \*vecty.HTML/) Render() vecty.ComponentOrHTML/g'
```

Obviously, you'll still need to verify that this only modifies your `Component` implementations. No other changes are needed, and no behavior change is expected for components that return `*vecty.HTML` (as the new `vecty.ComponentOrHTML` interface return type).

## Oct 14, 2017 ([PR #155](https://github.com/hexops/vecty/pull/155)): major breaking change

The function `prop.Class(string)` has been removed and replaced with `vecty.Class(...string)`.  Migrating users must use the new function and split their classes into separate strings, rather than a single space-separated string.

## Oct 1, 2017 ([PR #147](https://github.com/hexops/vecty/pull/147)): minor breaking change

`MarkupOrChild` and `ComponentOrHTML` can both now contain `KeyedList` (a new type that has been added)

## Sept 5, 2017 ([PR #140](https://github.com/hexops/vecty/pull/140)): minor breaking change

Package `storeutil` has been moved to `github.com/hexops/vecty/example/todomvc/store/storeutil` import path.


## Sept 2, 2017 ([PR #134](https://github.com/hexops/vecty/pull/134)): major breaking change

Several breaking changes have been made. Below, we describe how to upgrade your Vecty code to reflect each of these changes.

On the surface, these changes _may_ appear to be needless or simple API changes, however when combined they in fact resolve one of the last major open issues about how Vecty fundamentally operates. With this change, Vecty now ensures that the persistent pointer to your component instances remain the same regardless of e.g. the styles that you pass into element constructors.

### constructors no longer accept markup directly

`Tag`, `Text`, and `elem.Foo` constructors no longer accept markup (styles, properties, etc.) directly. You must now specify them via `vecty.Markup`. For example, this code:

```Go
func (p *PageView) Render() *vecty.HTML {
 	return elem.Body(
 		vecty.Style("background", "red"),
	 	vecty.Text("Hello World"),
 	)
}
```

Must now be written as:

```Go
func (p *PageView) Render() *vecty.HTML {
 	return elem.Body(
 		vecty.Markup(
	 		vecty.Style("background", "red"),
 		),
	 	vecty.Text("Hello World"),
 	)
}
```

### If no longer works for markup

`If` now only accepts `ComponentOrHTML` (meaning `Component`, `*HTML`, `List` or `nil`). It does not accept markup anymore (styles, properties, etc). A new `MarkupIf` function is added for this purpose. For example you would need to make a change like this to your code:

```diff
func (p *PageView) Render() *vecty.HTML {
 	return elem.Body(
 		vecty.Markup(
-			vecty.If(isBackgroundRed, vecty.Style("background", "red")),
+			vecty.MarkupIf(isBackgroundRed, vecty.Style("background", "red")),
 		),
 		vecty.Text("Hello World"),
 	)
}
```

### Other breaking changes

- `ComponentOrHTML` now includes `nil` and the new `List` type, rather than just `Component` and `*HTML`.
- `MarkupOrComponentOrHTML` has been renamed to `MarkupOrChild`, and now includes `nil` and the new `List` and `MarkupList` (instead of `Markup`, see below) types.
- The `Markup` _interface_ has been renamed to `Applyer`, and a `Markup` _function_ has been added to create a `MarkupList`.


## Aug 6, 2017 ([PR #130](https://github.com/hexops/vecty/pull/130)): minor breaking change

The `Restorer` interface has been removed, component instances are now persistent. Properties should be denoted via ``` `vecty:"prop"` ``` struct field tags.


## Jun 17, 2017 ([PR #117](https://github.com/hexops/vecty/pull/117)): minor breaking change

`(*HTML).Restore` is no longer exported, this method was not generally used externally.


## May 11, 2017 ([PR #108](https://github.com/hexops/vecty/pull/108)): minor breaking change

`(*HTML).Node` is now a function instead of a struct field.
