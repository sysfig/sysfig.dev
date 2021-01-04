---
title: Hugo vs Other Static Site Generators
slug: hugo-vs-other-static-site-generators
date: 2020-12-14T17:07:12-08:00
chapter: a
order: 2
tags:
    - web
    - hugo
draft: true
---

This book is about Hugo. But Hugo is not the only static site generator out there. So what makes Hugo so different? Before we go into that, let's first talk about the features you can expect from _any_ decent static site generators:

- Templating - HTML templates that includes placeholders for styles and content. You may have a template for blog posts and another for static pages, but both rendered inside a main template that includes a header and footer.
  Templates allows you to keep your code [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) by factoring out the common, re-usable elements.
- Development server - a background process that watches for changes in your code, rebuilds the sites, and serves the built site on a local port (which you can reach on an address like `localhost:1313`)

In this section, we compare Hugo with other static site generators.

- [Jekyll](https://jekyllrb.com/) - one of the first and most popular static site generator. Created by [Tom Preston-Werner](https://tom.preston-werner.com/), a co-founder of GitHub. Requires Ruby and installs as a Gem
- [11ty](https://www.11ty.dev/) - requires Node.js v8+ and installs as a (usually global) npm package
- [Zola]() - 

For me, the feature I like most about Hugo is the fact that it is available as a static binary. This means all of Hugo's dependencies are packaged into a single binary, which means you only need to download _one file_ and place it in your `PATH` and it's ready to use. This is unlike installing as a gem or npm package, which have dependencies on the language run time (e.g. Ruby and Node.js) and other gems/packages.

The only other popular static site generator that has this feature is Zola, written in Rust. The only reason why I picked Hugo for my site over Zola is simply because Hugo is more popular. All other things equal, a more popular tool typically means better support, either from the maintainers or the community.
