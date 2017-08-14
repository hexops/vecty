Changelog
=========

Although v1.0.0 [is not yet out](https://github.com/gopherjs/vecty/milestone/1), we do not expect many breaking changes. When there is one, however, it is documented clearly here.

Pre-v1.0.0 Breaking Changes
---------------------------

- Aug 6, 2017: The `Restorer` interface has been removed, component instances are now persistent. Properties should be denoted via ``` `vecty:"prop"` ``` struct field tags. ([PR #130](https://github.com/gopherjs/vecty/pull/130))
- Jun 17, 2017: `(*HTML).Restore` is no longer exported, this method was not generally used externally. ([PR #117](https://github.com/gopherjs/vecty/pull/117))
- May 11, 2017: `(*HTML).Node` is now a function instead of a struct field. ([PR #108](https://github.com/gopherjs/vecty/pull/108))
