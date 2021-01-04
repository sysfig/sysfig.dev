---
title: Pulling Out the Header and Footer
slug: header-footer-partials
date: 2020-12-15T16:04:49-08:00
chapter: e
order: 1
draft: true
---

Let's take our `layouts/index.html`.

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
            <a href="/about">About</a>
        </nav>
    </header>
    <main>
        <div class="page-content">{{ .Content }}</div>
    </main>
    <footer>sysfig.dev &copy; {{ time.Now.Year }}</footer>
</body>
</html>
```

The header and footer are the same for all pages, so let's extract them into a partial.

Create a new file at `layouts/partials/header.html` and copy in the entire `<header>` element.

```html
<header>
    <nav>
        <a href="/">Home</a>
        <a href="/series">Series</a>
        <a href="/posts">Posts</a>
        <a href="/projects">Projects</a>
        <a href="/about">About</a>
    </nav>
</header>
```

Then, replace the `<header>` element in both `layouts/index.html` and `layouts/_default/single.html` with this markup:

```go
{{- partial "header.html" . -}}
```

Everything between the pair of braces (`{{ }}`) is processed using Go's `template` package. The `{{- ` syntax is the `template` package's way of saying "trim all whitespaces preceding the opening braces"; likewise, ` -}}` means trim all whitespaces after the closing braces. This means `a b c {{- d }} e` will be outputted as `a b cd e`.

> Note that the `{{-`/`-}}` syntax does not trim whitespaces _within_ the expression, but on either side of the expression.

Within the braces are three directives:

- `partial` - the command that imports the partial.
- `"header.html"` - the filename of the partial within the `layouts/partials` directory.
- `.` - the third directive is the context you pass into the partial. By default, all layouts are passed the page's context, which is represented by the period `.`. By specifying the `.` here, we are passing in the page's context into the partial.

The last few lines of the `layouts/index.html` files should now look like this:

<pre><code>
&lt;body&gt;
    <mark>{{- partial &quot;header.html&quot; . -}}</mark>
    &lt;main&gt;
        <div class="page-content">{{ .Content }}</div>
    &lt;/main&gt;
    &lt;footer&gt;sysfig.dev &copy; {{ time.Now.Year }}&lt;/footer&gt;
&lt;/body&gt;
&lt;/html&gt;
</code></pre>

Let's go back to `http://localhost:1313` and see that our page content is still the same as before, which means our partials are being injected properly.

Now carry out the same exercise with the footer. You should end up with a file at `layouts/partials/footer.html` with the following content:

```html
<footer>sysfig.dev &copy; {{ time.Now.Year }}</footer>
```

And the `<footer>` element in both layouts replaced with the `{{- partial -}}` block.

<pre><code>
&lt;body&gt;
    {{- partial &quot;header.html&quot; . -}}
    &lt;main&gt;
        <div class="page-content">{{ .Content }}</div>
    &lt;/main&gt;
    <mark>{{- partial &quot;footer.html&quot; . -}}</mark>
&lt;/body&gt;
&lt;/html&gt;
</code></pre>

Go once more to `http://localhost:1313` to confirm the page is the same as before.
