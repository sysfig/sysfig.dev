---
title: Understanding Content Section
slug: understanding-content-section
date: 2020-12-15T15:27:41-08:00
chapter: d
order: 1
---

A [_section_](https://gohugo.io/content-management/sections/ in Hugo is a collection of content pages grouped by directory structure. For example, if we have a directory at `content/posts/`, then `posts` would be a section. First-level directories under `content/` are called _root sections_.

You can have sections within sections (i.e. [_nested sections_](https://gohugo.io/content-management/sections/#nested-sections)) by creating additional directories within the first-level directories, but these non-root sections need to contain an `_index.md` file.

Remember that the structure of the `content/` directory reflects the structure of the URL path in the site. So if we have a `posts` root section, then content within that section is going to have the URL path prefix `/posts/` (i.e. content file at `content/posts/how-dns-works` corresponds to the URL `<host>/posts/how-dns-works`).
