---
title: Adding Content to Section List Pages
slug: adding-content-to-section-list-pages
date: 2020-12-15T16:08:36-08:00
chapter: g
order: 2
---

Remember that `_index.md` is the name of the content file for page with an empty path; so to add some free form content to the `/posts/` list page, we can create a new content file at `content/posts/_index.md` with the following disclaimer.

```md
---
title: Posts
---

The list below shows only standalone posts; posts which are part of a series are not listed here.
```

Then, to display it, add a `<div class="page-content">{{ .Content }}</div>` block to `layouts/_default_/list.html`.

<pre><code>
{{ define "main" }}
  &lt;h1&gt;{{ .Title }}&lt;/h1&gt;
  <mark>&lt;div class="page-content"&gt;{{ .Content }}&lt;/div&gt;</mark>
  {{ range .Pages }}
</code></pre>

![](/img/posts-list-content.png)
