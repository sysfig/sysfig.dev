---
title: Creating Content
slug: creating-content
date: 2020-12-14T20:23:00-08:00
chapter: c
order: 1
draft: true
---

Next, we are going to create some content to put on our home page. The content goes in the `content/` directory and the name of the file will become the path on the final site. For example, if we add an `content/about.md` file, then a page at `/about` will be created.

You can create new content by running `hugo new [contentPath]`, which will create a new content file based on a specific archetype (or the default one if none are specified). However, for the home page, we are not going to use the `new` sub-command; instead, we will manually create a new file at `content/_index.md` with the following content.

```md
Welcome to `sysfig.dev`
```

The `content/_index.md` file name is Hugo's convention for naming the home page. If Hugo were to follow the aforementioned naming convention, then the home page content file will be named `.md` (since the home page has an empty path), which looks more like a _dotfile_. Prefixing the home page content file with an underscore also means it'll appear first in most file listings, making it easier to find.

> `_index.md` is the name of the content file for page with an empty path.

But having just content is not enough to create a web page. Remember that a web page is a composite of the content and a layout.

Later on, when you run `hugo` to generate the web page, Hugo will first parse the content file with the appropriate parser ([Goldmark](https://github.com/yuin/goldmark) for Markdown files). The parser generates some HTML, which is then inserted into a slot inside a layout template.
