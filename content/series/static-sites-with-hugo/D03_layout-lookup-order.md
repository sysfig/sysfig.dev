---
title: Understanding Hugo's Lookup Order
slug: hugo-lookup-order
date: 2020-12-15T15:27:30-08:00
chapter: d
order: 3
---

Circling back to creating our About page, we need to create a layout file that can match with the `content/about.md` content file. We need to create that file at a location that satisfies Hugo's layout lookup rules. You can find the full lookup order at [_Hugo's Lookup Order_](https://gohugo.io/templates/lookup-order/), but the abridged version can be summed up as:

- `layouts/<type>/<kind>.html`
- `layouts/_default/<kind>.html` - for pages that are not part of a section, or when the section has not defined its own layout

So to create a layout for the About page (which is not part of any sections), we can create a file at `layouts/_default/single.html` with the following content:

```html
<!DOCTYPE html>
<html lang="{{ .Site.LanguageCode }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ .Title }} - {{ .Site.Title }}</title>
</head>
<body>
    <header>
        <nav>
            <a href="/">Home</a>
            <a href="/series">Series</a>
            <a href="/posts">Posts</a>
            <a href="/projects">Projects</a>
            <a href="/about">About</a>
        </nav>
    </header>
    <main>
        <h1 class="title">{{ .Title }}</h1>
        <div class="page-content">{{ .Content }}</div>
    </main>
    <footer>sysfig.dev &copy; {{ time.Now.Year }}</footer>
</body>
</html>
```

This HTML layout is very similar to the `layouts/index.html` file - with a header, a body, and a footer. The only difference is the addition of a `h1` element and modification of the `title` element. The `{{ .Title }}` comes from the `title` front matter variable in the content file (you can use all predefined variables this way).

We will shortly look at how to avoid this duplication and keep our code [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself), but we now have the content and the layout for the About page, let's make sure everything works before we look at optimizations.
