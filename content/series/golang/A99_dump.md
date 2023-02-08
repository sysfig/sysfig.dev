---
title: Dump
slug: dump
date: 2020-12-30T00:00:57-08:00
chapter: a
order: 99
tags:
    - golang
---

Package naming:

- short
- no underscore_case or camelCase
- If you want to rename your packages, you can use [`gofmt`](https://golang.org/cmd/gofmt/)'s `-r` flag to rewrite package Go expressions, including package names. This is preferred to using the Find and Replace function in your code editor because `gofmt` is aware of the Go syntax and will only replace Go expressions, and not, say, the value of a string.

??gocode, godef, and go-outline??


---


In package repositories like [`npm`](https://www.npmjs.com/) (for JavaScript packages), they have a policy where once a package is published for a certain version, it cannot be unpublished again. They learnt this the hard way after the author of the `left-pad` package unpublished his package, which led to many major tools that depend on it failing to build due to a missing dependency. Although the disruption only lasted 2.5 hours (they ended up republishing the package), its effect was widely felt. You can read more about this story on the npm blog titled _[kik, left-pad, and npm](https://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm)_


goreportcard.com
