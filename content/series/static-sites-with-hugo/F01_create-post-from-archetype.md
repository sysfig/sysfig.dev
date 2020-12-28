---
title: Creating Post from Archetype
slug: post-from-archetypes
date: 2020-12-15T16:06:19-08:00
chapter: f
order: 1
---

Let's create a our first post about password hashing. You can do that by running `hugo new posts/what-is-password-hashing.md`. Because the path starts with `posts/`, Hugo will know to use the `posts` archetype.

This will create a content file at `content/posts/what-is-password-hashing.md` with the following content:

```md
---
title: "What Is Password Hashing"
date: 2020-12-12T16:58:14-08:00
draft: <mark>false</mark>
tags:
    - web
    - devops
    - linux
    - programming
---
```

Let's make a few small changes to the front matter and add our actual blog post content. For demonstration purposes, we are just going to use some dummy _lorem ipsum_ texts, but feel free to add any content you want here.

<pre><code>
---
title: "What <mark>i</mark>s Password Hashing<mark>?</mark>"
date: 2020-12-12T16:58:14-08:00
draft: true
tags:
    - web
    - programming
    <mark>- security</mark>
---

## Quotiens adsuetam subit Clymeneia innixus iuncta pertimuit

Lorem markdownum, partem _ille_ spectasse, herbas ille omnia promere. Pugnat meo
terrae origo clamant hac, sed quam Thebas. Ait riget discite epota nullum
commisisse ferrum!
...
</code></pre>

We are using [Lorem Markdownum](https://github.com/jaspervdj/lorem-markdownum) by [Jasper Van der Jeugt](https://jaspervdj.be/) to generate more realistic-looking content, rather than just plain text.

Now start the development server if you haven't already, and go to `http://localhost:1313/posts/what-is-password-hashing/`. You should see your post rendered.

![](/img/blog-post.png)
