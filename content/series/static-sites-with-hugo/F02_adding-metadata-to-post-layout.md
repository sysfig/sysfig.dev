---
title: Adding Metadata to Post Page
slug: post-page-metadata
date: 2020-12-15T16:07:02-08:00
chapter: f
order: 2
draft: true
---

The post page uses the `layouts/_default/single.html` page, which displays only the title and content. Some information from the post is lost, such as the date and tags.

So let's create a more specific layout for single pages in the `posts` section by creating a new layout file at `layouts/posts/single.html` with the following content:

```html
{{ define "main" }}
    <h1 class="title">{{ .Title }}</h1>
    <div class="post__meta">// {{ .Date.Format "2 Jan 06" }} &bullet; {{ range .Params.tags }}
        <div class="tag">{{ . }}</div>
      {{ end }}
    </div>
    <div class="page-content">{{ .Content }}</div>
{{ end }}
```

Then, go back to `http://localhost:1313/posts/what-is-a-mac-address/`, and you should see the tags and date being displayed.

![](/img/post-page-with-meta.png)
