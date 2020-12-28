---
title: Using Base Templates and Blocks
slug: base-templates-blocks
date: 2020-12-15T16:05:18-08:00
chapter: e
order: 3
---

We've componentized most of the layouts into partials, but the layout themselves look almost identical to each other.

`layouts/index.html`
```html
<!DOCTYPE html>
<html lang="{{ .Site.LanguageCode }}">
{{- partial "head.html" . -}}
<body>
    {{- partial "header.html" . -}}
    <main>
        <div class="page-content">{{ .Content }}</div>
    </main>
    {{- partial "footer.html" . -}}
</body>
</html>
```

`layouts/_default/single.html`
```html
<!DOCTYPE html>
<html lang="{{ .Site.LanguageCode }}">
{{- partial "head.html" . -}}
<body>
    {{- partial "header.html" . -}}
    <main>
        <h1 class="title">{{ .Title }}</h1>
        <div class="page-content">{{ .Content }}</div>
    </main>
    {{- partial "footer.html" . -}}
</body>
</html>
```

Can we do more to remove this last bit of duplication? In fact, we can - with a _base template_ and _blocks_.

Create a new base template at `layouts/_default/baseof.html` with the following content:

```html
<!DOCTYPE html>
<html lang="{{ .Site.LanguageCode }}">
{{- partial "head.html" . -}}
<body>
    {{- partial "header.html" . -}}
    <main>
        {{- block "main" . }}{{- end }}
    </main>
    {{- partial "footer.html" . -}}
</body>
</html>
```

A base template allows us to create a common template for every page. Within the base template are placeholder for named blocks. Here, we are defining a placeholder for a block named `main`. Now, when Hugo constructs a page, it will start with the base template and try to find a block named `main` from the regular layout template that matched with the content file; if found, it will inject that block into the base template and use the result as the layout. If the `main` block cannot be found, then it will revert back to using the regular layout template (if found).

With that in mind, we now need to define a block named `main` within our regular layout files.

Replace the entirety of the `layouts/index.html` with:

```go
{{ define "main" }}
    <div class="page-content">{{ .Content }}</div>
{{ end }}
```

And do the same with `layouts/_default/single.html`

```go
{{ define "main" }}
    <h1 class="title">{{ .Title }}</h1>
    <div class="page-content">{{ .Content }}</div>
{{ end }}
```

Now go back to `http://localhost:1313` and `http://localhost:1313/about`, you should see no difference, which means our base templates and blocks are working correctly.

Using blocks is similar to partials, but the important difference is that the `partial` command injects the same partial for every page, but `block` injects a different block based on the matched layout.

This concludes the refactoring of our code. In the next chapter, we will look into creating blog posts.
