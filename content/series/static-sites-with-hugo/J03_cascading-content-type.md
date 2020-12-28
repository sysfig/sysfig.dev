---
title: Cascading Content Type
slug: cascading-content-type
date: 2020-12-15T20:49:54-08:00
chapter: i
order: 3
---

We could manually add `type: posts` to the archetype but this makes maintainance difficult - imagine if we want to change the type for thousands of series pages (even though we could do a directory-global find-and-replace, it'll look real bad when doing code reviews).

> Avoid adding `type` to the archetype. A page's content type should, most of the time, be determined from its file location; use `type` to override only in special cases.

Instead, use Hugo's [Front Matter Cascade](https://gohugo.io/content-management/front-matter#front-matter-cascade) feature, which allows any node (anything page that is not a single page) to pass front matter for its descendents to inherit. To do so, specify the properties to pass down under the `cascade` front matter key. So for us to set the `type` property for all pages in `content/series/*/`, we can create a file at `content/series/_index.md` with the following content:

```yaml
---
cascade:
  type: posts
---
```

This is much cleaner as the overriding logic is encapsulated in to 2 lines.
