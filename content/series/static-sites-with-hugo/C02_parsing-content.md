---
title: Parsing Content
slug: parsing-content
date: 2020-12-14T20:23:25-08:00
chapter: c
order: 2
---

For any file that is not HTML, Hugo uses a parser to parse the content and convert it to HTML. For Hugo to be able to parse your content, you must use a supported content format (like Markdown). Hugo supports content written in Markdown or HTML out-of-the-box. Additional content formats are supported with the installation of additional tools on your machine (Hugo calls these tools _external helpers_).

Hugo determines the content format of a file based on a set of clues:

1. The `markup` front matter value (front matter is the structured metadata of a content file)
2. The file extension (e.g. `.md` or `.html`)

You can find the most up-to-date list of valid values at _[List of content formats](https://gohugo.io/content-management/formats/#list-of-content-formats)_.
