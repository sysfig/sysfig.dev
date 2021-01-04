---
title: Pulling Out the Head Element
slug: head-partial
date: 2020-12-15T16:04:24-08:00
chapter: e
order: 2
draft: true
---

Next, let's extract the `<head>` element from both layouts into its own partial. This is slightly more complicated than with the header and footer because the `<title>` element within the `<head>` is different for different layouts. With `layouts/index.html`, it's `<title>{{ .Site.Title }}</title>`; with `layouts/_default/single.html`, it's `<title>{{ .Title }} - {{ .Site.Title }}</title>`.

To make it easier for us, let's assume that the `<title>` element will be `{{ .Site.Title }}` only for the home page, and `{{ .Title }} - {{ .Site.Title }}` for all other pages. This allows us to use `.isHome` page variable in the context, which has the value `true` only when it's the home page.

Create a new partial at `layouts/partials/head.html` with the following content:

```html
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ if not .IsHome }}{{ .Title }} - {{ end }}{{ .Site.Title }}</title>
</head>
```

It uses the `.isHome` variable to display the page's title only when it's not the home page. Now, make the change to the regular layout templates to use this partial.

<pre><code>
&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;{{ .Site.LanguageCode }}&quot;&gt;
<mark>{{- partial &quot;head.html&quot; . -}}</mark>
&lt;body&gt;
...
</code></pre>
