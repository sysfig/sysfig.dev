---
title: Different Kinds of Pages
slug: page-kinds
date: 2020-12-15T15:25:25-08:00
chapter: d
order: 2
draft: true
---

But what happens when a user hits the URL `<host>/posts/`? Typically, this page would be a chronological  index/listing of all posts on the site. Hugo makes a distinction between these _kinds_ of _list pages_ as opposed to the regular _single page_ (e.g. About, Terms, or a blog post).

There are different kinds of list pages - home page, section listing, taxonomy lists, and taxonomy terms. The page at `<host>/posts/` would be a type of section listing.

---


All pages belong to one of 5 _kinds_:

- home
- page
- section
- taxonomy
- (taxnomy) term

Regular pages also have a _[content type](https://gohugo.io/content-management/types/_, which defaults to the root section (first directory within `content/`) but can be overriden using the `type` front matter variable.

The content type is used to determine which archetype and layout to use.
