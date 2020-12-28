---
title: Overriding Default List Layout
slug: overriding-default-list-layout
date: 2020-12-15T16:09:29-08:00
chapter: h
order: 1
---

We can also go to the taxonomy list page (i.e. `/tags/`) to see a list of all tags, but it doesn't look quite right. It's also using the `layouts/_default/list.html` layout, but that layout is meant for posts, not tags (or any type of taxonomy terms), so let's create a more specific layout for our tags.

If we look at the _[Hugo's Lookup Order](https://gohugo.io/templates/lookup-order/)_ documentation, we can see that the `layouts/_default/list.html` can be overriden on the taxonomy list page with a file at `layouts/tags/terms.html`. So create a file at `layouts/tags/terms.html` with the following content:

```html
{{ define "main" }}
  <h1 class="title">{{ .Title }}</h1>
  <div class="page-content">{{ .Content }}</div>
  {{ range .Pages }}
    <a class="tag-list-item" href="{{ .RelPermalink }}">{{ .Title }}</a>
  {{ end }}
{{ end }}
```

![](/img/taxonomy-tags-list.png)
