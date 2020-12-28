---
title: Deployment Preparations
slug: deployment-preparations
date: 2020-12-15T16:21:05-08:00
chapter: z
order: 00
draft: true
---

By default, Hugo does not clean up the `public/` directory when you run `hugo`. It will update files that needs updating, but it will not remove files that shouldn't be there. We can instruct `hugo` to start with a clean directory each time by specifying the `--cleanDestinationDir​` flag.

`--minify​` - removes whitespace characters
