---
title: Notes
slug: notes
date: 2020-12-15T16:21:45-08:00
chapter: z
order: 00
draft: true
---

In fact, the underscore (`_`) prefix in the file name tells Hugo that this content file is for a list page.

---

But what if I need to process the images? Creating different sizes of images and compressing them is a common requirement.

Create a new directory called `assets/` at the root of the project. Files stored here are processed by [Hugo Pipes](https://gohugo.io/hugo-pipes/). The processed assets can be inserted into ??layouts and content?? by using `.Permalink` or `.RelPermalink`. If an asset is not linked via `.Permalink` or `.RelPermalink`, it does not get added to the `public/` directory.
