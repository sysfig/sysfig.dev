---
title: Installing Hugo
slug: installing-hugo
date: 2020-12-14T17:52:37-08:00
chapter: a
order: 4
draft: true
---

Hugo comes in two versions - regular and extended. The extended version provides support for:

- Asset management
- CSS preprocessing - Sass

Although Hugo is available as a single binary, the recommended approach for normal, human-centric usage is still to use a package manager (like Homebrew for macOS). This is because it makes it easier to upgrade. With a binary, you'd basically have to follow the same installation process and replace the old binary with the new. With a package manager, however, it's as simple as `brew upgrade hugo`.

In an automated pipeline, however, using the binary is recommended as you can avoid the bloat of having to download and install a package manager when all you really need to do is build your site.

To confirm that `hugo` is installed properly, run `hugo version`.

```
$ hugo version
Hugo Static Site Generator v0.79.0/extended darwin/amd64 BuildDate: unknown
```

"""LibSASS is currently the only reason we build an extended version"""(https://discourse.gohugo.io/t/question-are-there-plans-to-support-dart-sass-and-its-newly-introduced-use-modular-system/21882/13)
