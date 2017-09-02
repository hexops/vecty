Changelog
=========

Although v1.0.0 [is not yet out](https://github.com/gopherjs/vecty/milestone/1), we do not expect many breaking changes. When there is one, however, it is documented clearly here.

Pre-v1.0.0 Breaking Changes
---------------------------


## Aug 6, 2017 ([PR #130](https://github.com/gopherjs/vecty/pull/130)): minor breaking change

The `Restorer` interface has been removed, component instances are now persistent. Properties should be denoted via ``` `vecty:"prop"` ``` struct field tags.


## Jun 17, 2017 ([PR #117](https://github.com/gopherjs/vecty/pull/117)): minor breaking change

`(*HTML).Restore` is no longer exported, this method was not generally used externally. 


## May 11, 2017 ([PR #108](https://github.com/gopherjs/vecty/pull/108)): minor breaking change

`(*HTML).Node` is now a function instead of a struct field.
