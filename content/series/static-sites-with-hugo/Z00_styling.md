---
title: Adding Styles
slug: styles
date: 2020-12-15T16:18:23-08:00
chapter: z
order: 00
draft: true
---

Add a link tag

```html
<link rel="stylesheet" href="{{ css/style.css | relURL }}">
```

[`relURL`](https://gohugo.io/functions/relurl/) is one of Hugo's built-in functions that constructs a relative URL. So `css/style.css | relURL` would be `/css/style.css`. One benefit of writing it this way (as opposed to just using `href="/css/style.css"`) is that you can quickly switch between `relURL` and [`absURL`](https://gohugo.io/functions/absurl/, which will build the URL to include the base URL (the string specified with the site configuration variable `baseURL`).
