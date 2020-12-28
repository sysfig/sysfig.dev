---
title: Pulling Out Common Elements into Partials
slug: partials
date: 2020-12-15T15:28:03-08:00
chapter: e
order: 0
---

Now that our home and About pages are working, let's do some refactoring and remove duplication from our code.

Our `layouts/index.html` and `layouts/_default/single.html` shares a lot of code. In fact, they differ only on a couple of lines. Just like you'd extract common logic in code and abstract them into a function or module, we can extract common HTML elements from the layouts template into re-usable HTML component templates called _partials_. Once extracted, we can reinsert the partial back into the layout using the `partial` command.

```html
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
```


```html
<!DOCTYPE html>​
<html lang='{{ .Site.LanguageCode }}'>
    {{- partial "head.html" . -}}
    <body>
        <header>
            {{- partial "header.html" . -}}
        </header>
        <main​>
            {{- block "main" . }}{{- end }}
        </main>
        <footer>
            {{- partial "footer.html" . -}}
        </footer>
    </body>
</html>
```
