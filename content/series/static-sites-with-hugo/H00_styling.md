---
title: Adding Styles with Sass
slug: adding-styles-with-sass
date: 2020-12-16T15:24:30-08:00
chapter: h
order: 0
---

We have a working site, but it's not very pretty nor user-friendly. What it needs is some styles to keep the elements aligned and organized.

The simplest way to add styles to the site is to create a CSS stylesheet inside the `static/` directory (e.g. `static/style.css`), and add a `<link rel="stylesheet" href="/style.css">` to the `<head>` element.

<pre><code>
    &lt;title&gt;{{ if not .IsHome }}{{ .Title }} - {{ end }}{{ .Site.Title }}&lt;/title&gt;
    <mark>&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ &quot;style.css&quot; | relURL }}&quot;&gt;</mark>
&lt;/head&gt;
</code></pre>

Next, we should also add a logo to the site. We could place it in `static/img/`, but then we're mixing up the images that we add to posts to the images that are used for styles.

Minify CSS + JavaScript - reduce file size (results in faster transfer time and shorter time to first load/usable) by removing whitespaces, comments, and shortening variable names.

Fingerprinting CSS and JavaScript - https://web.dev/http-cache/#invalidating_and_updating_cached_responses - the name of the file is derived from its content, so that if the content changes, the name changes as well. This will help invalidate the cache.

Typically, you may use a tool like Webpack to handle the minification and fingerprinting. But the extended version of Hugo have all these features already.