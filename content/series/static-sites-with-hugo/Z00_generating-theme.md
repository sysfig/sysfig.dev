---
title: Generating a Theme
slug: theme
date: 2020-12-15T16:19:40-08:00
chapter: z
order: 00
draft: true
---

You can generate a new theme by running `hugo new theme <name>`.

```
hugo new theme sysfig
```

???`tree` output??

You should be familiar with the `themes/<name>` directory structure, as it is just a stripped down version of our site directories.

Start by moving our archetypes and layout files into the `themes/<name>` directory.

We must tell `hugo` which theme to use for the site by adding the `theme` property to the site configuration.

<pre><code>
baseURL: https://sysfig.dev
languageCode: en-US
title: sysfig
<mark>theme: sysfig</mark>
</code></pre>
