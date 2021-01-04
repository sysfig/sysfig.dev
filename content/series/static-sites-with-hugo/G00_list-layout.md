---
title: Adding Posts List Layout
slug: posts-list-layout
date: 2020-12-15T16:07:31-08:00
chapter: g
order: 0
draft: true
---

We can reach individual posts by going to `http://localhost:1313/posts/<name>` but what if want to see a list of all posts? If you try `http://localhost:1313/posts/`, you'll be served a blank page because there is no content file or layout file for the section list page (`layouts/index.html` only works for the home page and `layouts/_default/single.html` only works for single pages).

We can create a list layout at `layouts/_default/list.html` that acts as the default list layout. If we need a different layout later (e.g. for series), we can create a more specific layout then.

```html
{{ define "main" }}
  <h1 class="title">{{ .Title }}</h1>
  <div class="page-content">{{ .Content }}</div>
  {{ range .Pages }}
    <div class="post-list-item">
      <div class="post-list-item__meta">// {{ .Date.Format "2 Jan 06" }} &bullet; {{ range .Params.tags }}
          <div class="tag">{{ . }}</div>
        {{ end }}
      </div>
      <a class="post-list-item__title" href="{{ .RelPermalink }}">{{ .Title }}</a>
    </div>
  {{ end }}
{{ end }}
```

List layouts are given a different context to regular pages. With a section list page layout, the page context of every page in the section are available under `.Pages`. This allows us to iterate over `.Pages` using the [`range`](https://gohugo.io/functions/range/) function to generate links to all pages in that section.

Within the `{{ range .Pages }}` block, the context changes to the page's context. That's why the top-level `{{ .Title }}` displays the title of the list page (i.e. `posts`), whilst the `{{ .Title }}` within the `{{ range .Pages }}` block displays the title of each individual post.

Within each page, we are using the [`.Format`](https://gohugo.io/functions/format/) function of `.Date` to format the date (taken from the `date` front matter field); without formatting, the date would read something like `2020-12-12 16:58:14 -0800 PST`.

An additional `range` block iterates over the posts' tags. We have to use `.Params.tags` instead of `.Tags` because `tags` is not a [predefined](https://gohugo.io/content-management/front-matter/#predefined) front matter variable.

Let's head over to `http://localhost:1313/posts/` and see that our section list page is now displaying details for a single post.

![](posts-list-single-post.png)
