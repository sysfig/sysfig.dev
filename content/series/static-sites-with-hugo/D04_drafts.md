---
title: Drafts
slug: drafts
date: 2020-12-15T15:27:30-08:00
chapter: d
order: 4
draft: true
---

Start the Hugo development server by running `hugo server` and go to `http://localhost:1313/about/`. Contrary to our expectations, it's displaying a page that reads `404 page not found`. So what's wrong?

If we look at our `content/about.md` content file, we'll find a front matter value of `draft: true`. `draft` is one of the many _[predefined variables](https://gohugo.io/content-management/front-matter/#predefined)_ that carries a special meaning with Hugo (`title` and `date` are also predefined variables). All content with `draft` set to `true` will not be rendered (unless overriden using the `--buildDrafts` option to the `hugo` command).

So to make our About page appear, change the `draft` value from `true` to `false`.

<pre><code>
---
title: "About"
date: 2020-12-09T20:48:46-08:00
draft: <mark>false</mark>
---

Hi! My name is X. I am a DevOps engineer. I specialize in Y and Z.
</code></pre>

Now when we go to `http://localhost:1313/about/` we can see the rendered About page.

![](/img/about-page.png)
