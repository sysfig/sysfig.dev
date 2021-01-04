---
title: Adding Links to Tag Pages
slug: linking-tag-pages
date: 2020-12-15T16:10:37-08:00
chapter: h
order: 2
draft: true
---

Now we that there is a taxonomy term listing page for for each tag, let's go back to our list and single page layouts and add links to them.

Go to `layouts/_default/list.html` and `layouts/posts/single.html` and surround the `{{ . }}` block with an `<a>` tag.

<pre><code>
&lt;div class=&quot;tag&quot;&gt;<mark>&lt;a href=&quot;/tags/{{ . }}/&quot;&gt;</mark>{{ . }}<mark>&lt;/a&gt;</mark>&lt;/div&gt;
</code></pre>
