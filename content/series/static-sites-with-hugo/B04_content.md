---
title: content/
slug: content
date: 2020-12-14T20:21:14-08:00
chapter: b
order: 4
---

Whilst the `layouts/` directory houses _HTML_ elements that can be reused across multiple web pages, the `content/` directory houses content that are unique for that page. Examples include the content of an 'About Us' page, or the content of a blog post.

A file in `content/` consists of two parts - the first is a structured, YAML-formatted metadata block (called _front matter_), and the second is some unstructured content.

The YAML metadata block is enclosed in a pair of triple hyphens (`---`) and can be used to store named pieces of data like `title`, `draft`, `author`, etc. The values specified here can be used to fill in placeholders found in the layouts.

Hugo supports unstructured content written in HTML and Markdown out-of-the-box, but support for additional content formats (Emacs Org-Mode, AsciiDoc, RST, Pandoc) is possible if you install additional libraries.

The content gets converted into HTML before being embedded into the layout template to produce the final web page.
