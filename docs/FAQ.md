# FAQ

- [Does Vecty support server-side rendering?](#does-vecty-support-server-side-rendering)
- [Is Vecty being ued in production?](#is-vecty-being-used-in-production)

## Does Vecty support server-side rendering?

No. While Vecty _does_ enable you to share your Go code across backend and frontend, it does not support server-side rendering.

There are two primary reasons why server-side rendering is desired:

1. Quicker page loads, i.e. so the user doesn't need to download a large JavaScript bundle before seeing your webpage for the first time.
2. Compatability with web scrapers such as search engines, etc.

While these are two very serious and very real problems: we think that server-side rendering is not the way to resolve them.

- Code splitting (i.e. small-in-scope javascript bundles, instead of 'one big app bundle') is generally just as quick as server-side rendering is, for a user trying to view your webpage for the first time.
- More and more web scrapers, like Google, are running JavaScript in order to index websites.
- Even when web crawlers do scrape HTML, it is often hard for them to find relevant information. Most professional websites use proper meta-tags or serve separate HTML pages to web crawlers for exactly this reason.
- Server-side rendering introduces a large amount of 'platform'-specific (backend vs. frontend) code which becomes a non-negligable logical overhead when trying to reason about code.

We are open to hearing opposing thoughts, though, so please [get in touch](index.md#get-in-touch)!

## Is Vecty being used in production?

Due to the project's [current experimental status](/README.md#current-status), we're not aware of anyone using Vecty in production yet.
