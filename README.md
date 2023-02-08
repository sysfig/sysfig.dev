# sysfig.dev

## Quickstart

1. [Install Hugo](https://gohugo.io/installation/)
2. Run `hugo server`

## Walkthrough

Hugo separates content with representation. We are taking things one step further and separated everything that is not content (not in `content/` or `data/`) and encapsulating it into a theme called `sysfig`.

We start by looking at `themes/sysfig/layouts/_default/baseof.html`, which holds the base template used as the outermost shell of our site.

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
    <script src="{{ "js/main.js" | relURL }}"></script>
</body>
</html>
```

The base template includes partials for:

- the `<head>` (`themes/sysfig/layouts/partials/head.html`) - set `<title>` and stylesheets
- header (`themes/sysfig/layouts/partials/header.html`) - navigation
- footer (`themes/sysfig/layouts/partials/footer.html`) - purely decorative

At the bottom is a link to a static JavaScript file at `themes/sysfig/static/js/main.js`, whose purpose is to enable the menu button to work on mobile.

```js
document.getElementById("menu-button").addEventListener('click', () => {
  document.getElementById("main-nav").classList.toggle("mobile");
});
```

## TODOs

- Add [Open Graph](https://ogp.me/) metadata
- Add created as well as last-updated date
- Add maturity levels (e.g. sprout, bush, tree)
