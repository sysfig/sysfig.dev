---
title: Creating Series List Page
slug: series-list-page
date: 2020-12-16T14:44:02-08:00
chapter: i
order: 6
draft: true
---

Our fifth constraint is that '`/series/` should list out a list of series pages, but not the series posts'. So let's create a layout for list pages within the `series` section.

```html
{{ define "main" }}
  <h1 class="title">{{ .Title }}</h1>
  <div class="page-content">{{ .Content }}</div>
  {{ range .Pages }}
    {{ if and .IsSection (not (eq .CurrentSection.Title "Series")) }}
      <div class="series-list-item">
        <div class="series-list-item__meta">// {{ .Date.Format "2 Jan 06" -}}<span class="separator">&bullet;</span>{{- range .Params.tags }}
            <div class="tag"><a href="/tags/{{ . }}/">{{ . }}</a></div>
            {{ end }}
        </div>
        <a class="series-list-item__title" href="{{ .RelPermalink }}">{{ .Title }}</a>
      </div>
    {{ end }}
  {{ end }}
{{ end }}
```

We are using the `.IsSection` filter here to only keep section list pages (not single pages) and `(not (eq .CurrentSection.Title "Series"))` is used to exclude the `/series/` page itself.

![](/img/series-list.png)
