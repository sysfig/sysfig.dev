---
title: Creating an About Page
slug: creating-about-page
date: 2020-12-15T15:24:46-08:00
chapter: d
order: 0
draft: true
---

We'll soon polish up the home page by adding some styling and _pazazz_ to it. But before that, let's finish up laying out the structure of our site by adding more pages to the site. We'll start with the 'About' page.

Just as with the home page, we start by creating a content file. But instead of manually creating the file yourself, let's use the `hugo new` command.

```
hugo new about.md
```


```
/Users/sysfig/projects/sysfig.dev/content/about.md created
```

Because we used the `hugo new` command, Hugo will use one of the archetypes as the basis of our new content file (remember that archetypes are templates for new content). Just like layout selection, Hugo has a well-defined set of rules for selecting archetypes, which we will cover later. For now, just know that our `archetypes/default.md` archetype can be used as the default archectype for all content files.

Our newly-generated `content/about.md` looks like this:

```
---
title: "About"
date: 2020-12-09T20:48:46-08:00
draft: true
---
```

You can see that Hugo ran the `replace` command in `{{ replace .Name "-" " " | title }}` to generate the `title` from the file name. It also replaced `.Date` with the actual date and time when the `hugo new` comman was ran.

Let's add to that file by adding a line or two about yourself.

<pre><code>
---
title: "About"
date: 2020-12-09T20:48:46-08:00
draft: true
---
<mark>Hi! My name is X. I am a DevOps engineer. I specialize in Y and Z.</mark>
</code></pre>

The `content/about.md` file needs to be matched with a layout. The `layouts/index.html` layout only works for the home page; we need another layout file for the About page.

But what should we name the layout file? Where should we place it? With these questions in mind, now is a good time to dive a little deeper into Hugo's layout lookup order.

When Hugo generates the site, each content file gets turned into a web page. After processing the content file to HTML, Hugo needs to inject that content into a HTML layout file to create the final page. We've already seen that the `content/_index.md` content file can be matched with the `layouts/index.html` layout file, but what are the general rules for matching layouts with content?

To understand this, we must first understand what a content section is.
