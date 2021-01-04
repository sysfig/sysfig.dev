---
title: Creating Layout
slug: creating-layout
date: 2020-12-14T21:10:57-08:00
chapter: c
order: 3
draft: true
---

After Hugo has parsed your content file, it will then try to match the content with a layout from the `layouts/` directory.

Hugo matches content with layout by a set of well-defined rules, documented on the [_Hugo's Lookup Order_](https://gohugo.io/templates/lookup-order/) page. We will examine these lookup rules more closely shortly; for now, just know that a layout file with the name `index.html` will match our Home page content and, together, generates a complete web page.

So let's now create that layout file at `layouts/index.html` with the following content:

```html
<!DOCTYPE html>
<html lang="{{ .Site.LanguageCode }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ .Site.Title }}</title>
</head>
<body>
    <header>
        <nav>
            <a href="/">Home</a>
            <a href="/series">Series</a>
            <a href="/posts">Posts</a>
            <a href="/projects">Projects</a>
            <a href="/now">Now</a>
        </nav>
    </header>
    <main>
        <div class="page-content">{{ .Content }}</div>
    </main>
    <footer>sysfig.dev &copy; {{ time.Now.Year }}</footer>
</body>
</html>
```

Most of this HTML template is just boilerplate for a HTML document. We have a header navigation, some content, and a footer.

Hugo uses Go's `http/templates` library for templating, where the templating syntax are enclosed in a pair of braces `{{ }}`. You can use certain Go functions like we are doing with `{{ time.Now.Year }}`, which runs the Go function `time.Now().Year()` to retrieve the current year.

Data from various sources are also injected into the template via a _context_ object, represented by the starting period (`.`). You can find all properties available within the page context in _[Page Variables](https://gohugo.io/variables/page/)_. For instance, in that article, you'll find the `.Site` variable, which makes available everything that was specified inside the configuration file (`config.yaml`). Here, we are using `{{ .Site.LanguageCode }}` and `{{ .Site.Title }}` to inject the values of `languageCode` and `title` into the HTML file.

Within the context object, `.Content` refers to the unstructured content of the web page, which, for our home page, is the HTML-equivalent of ``Welcome to `sysfig.dev` ``.
