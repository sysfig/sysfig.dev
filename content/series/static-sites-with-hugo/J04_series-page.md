---
title: Creating Series Page
slug: series-page
date: 2020-12-15T21:24:44-08:00
chapter: i
order: 4
---

Let's tackle our third constraint next - every series should have its own _series page_ at `/series/<series-name>/` and it should show a list of all posts that belong to it.

I'd imagine a series page is going to look a lot different from a post page, with different kind of metadata, so let's create a new archetype for series. Create a new archetype at `archetypes/series.md` with the following content:

```yaml
---
title: {{ replace .Name "-" " " | title }}
date: {{ .Date }}
description:
prerequisites:
  required:
    - x
  recommended:
    - x
learning-goals:
  - x
out-of-scope:
  - x
what-you-will-build:
  - name: x
    description: x
    image: /img/x.jpg
software:
  supported:
    - name: x
      versionStart: 0
      versionEnd: 1
  tested:
    - name: x
      version: 1
tags:
    - web
    - devops
    - linux
    - programming
draft: true
---
```

`hugo new series/static-sites-with-hugo/_index.md`

Because the target content file would be inside the `content/series/` directory, Hugo will automatically use the `series` archetype. Open the file and make changes to tailor it to your specific series.

We also need a new layout for this new content type. Create a new layout at `layouts/series/single.html` with the following content:

```html
{{ define "main" }}
  <h1 class="title">{{ .Title }}</h1>
  <h2>About this Series</h2>
  <div class="page-content">{{ .Content }}</div>
  <h2>Prerequisites</h2>
  <ul class="prerequisites">
    {{ range .Params.prerequisites.required }}
      <li class="prerequisites__required-item">{{ . }}</li>
    {{ end }}
    {{ range .Params.prerequisites.recommended }}
      <li class="prerequisites__recommended-item">(optional) {{ . }}</li>
    {{ end }}
  </ul>
  <h2>What You'll Learn</h2>
  <ul class="learning-goals">
    {{ range (index .Params "learning-goals") }}
      <li class="learning-goals__item">{{ . }}</li>
    {{ end }}
  </ul>
  <h2>What You'll Build</h2>
  <div class="what-you-will-build">
    {{ range (index .Params "what-you-will-build") }}
      <div class="what-you-will-build__item">
        <img class="what-you-will-build__item__image" src="{{ .image }}" alt="">
        <div class="what-you-will-build__item__name">{{ .name }}</div>
        <div class="what-you-will-build__item__description">{{ .description }}</div>
      </div>
    {{ end }}
  </div>
  <h2>Syllabus</h2>
  {{ range .Pages }}
    {{ if .IsPage}}
        <a class="series-item" href="{{ .RelPermalink }}">
          <span class="series-item__title">{{ .Title }}</span>
          <span class="series-item__reading-time">{{ .ReadingTime }}m</span>
        </a>
    {{ end }}
  {{ end }}
  <h3>Software Used</h3>
  <p>This series should work for the following version ranges of software:</p>
  <ul class="software__supported">
    {{ range .Params.software.supported }}
      <li class="software__supported__item">{{ .name }} v{{ .versionStart }} - {{ if .versionEnd -}} v{{- .versionEnd -}} {{- else }} current {{- end -}}</li>
    {{ end }}
  </ul>
  <p>This series has been tested against:</p>
  <ul class="software__tested">
    {{ range .Params.software.tested }}
      <li class="software__tested__item">{{ .name }} v{{ .version }}</li>
    {{ end }}
  </ul>
  <h3>What's Out of Scope</h3>
  <ul class="out-of-scope">
    {{ range (index .Params "out-of-scope") }}
      <li class="out-of-scope__item">{{ . }}</li>
    {{ end }}
  </ul>
{{ end }}
```

Now, if you go to `http://localhost:1313/series/static-sites-with-hugo` you'll see this:

![](/img/series-page.png)

We can now check the third constraint off our list.
