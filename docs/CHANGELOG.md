Changelog
=========

Although v1.0.0 [is not yet out](https://github.com/gopherjs/vecty/milestone/1), we do not expect many breaking changes. When there is one, however, it is documented clearly here.

Pre-v1.0.0 Breaking Changes
-------------------------

- Jun 17, 2017: `(*HTML).Restore` is no longer exported, this method was not generally used externally. ([PR #117](https://github.com/gopherjs/vecty/pull/117))
- May 11, 2017: `(*HTML).Node` is now a function instead of a struct field. ([PR #108](https://github.com/gopherjs/vecty/pull/108))
