---
title: Introducing Taxonomies
slug: taxonomies
date: 2020-12-15T16:08:54-08:00
chapter: h
order: 0
---

We are now able to display a single blog post as well as a listing of blog posts. But what if I am only interested in blog posts about the web? Would I be able to only get a listing of posts with the `web` tag?

Yes, you can - with _[taxonomies](https://gohugo.io/content-management/taxonomies/)_.

Taxonomies are user-defined groupings of content. For our use case, we'd like to group our posts by their `tags` front matter property. Therefore, we need to make `tags` into a taxonomy.

We can specify which front matter properties we want to use as taxonomies in the site configuration using the `taxonomies` key.

```yaml
taxonomies:
  category: categories
  series: series
  tag: tags
```

The `taxonomies` object takes the singular labels as keys, and the plural label as values. If you do not specify the `taxonomies` site configuration, Hugo defaults to:

```yaml
taxonomies:
  category: categories
  tag: tags
```

Therefore, without doing anything `tags` is already considered a taxonomy by Hugo.

Taxonomies are grouped by taxonomy _terms_. If we look at our MAC address post, we can see that there are two items in the `tags` list.

```md
---
title: "What is a MAC Address?"
date: 2020-12-13T16:26:35-08:00
draft: false
tags:
    - web
    - networking
---
...
```

Each of the list items (`web` and `networking`) is a _term_ within the `tags` taxonomy. And because the MAC address post is tagged with `web` and `networking`, the MAC address post itself is a _value_ within the `web` and `networking` terms.

> _Taxonomies_ contain _terms_. Pages that belongs to a term are _values_ of that term.

Hugo will automatically create list pages for each taxonomy (i.e. `/tags/`), as well as each distinct taxonomy term (e.g. `/tags/web/`).

This means we can go to `/tags/web/` and find a list of all posts that carry the `web` tag.

![](/img/taxonomy-tags-web.png)

Because this is a list page, it can use the `layouts/_default/list.html` layout we defined earlier. If we want, we can also define a more specific layout only for taxonomy term list pages at `layouts/tags/term.html`, or any other files specified in the _[Hugo's Lookup Order](https://gohugo.io/templates/lookup-order/)_ documentation.
