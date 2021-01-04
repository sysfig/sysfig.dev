---
title: Creating Posts Archetype
slug: posts-archetypes
date: 2020-12-15T16:05:54-08:00
chapter: f
order: 0
draft: true
---

We should create an archetype for posts because we want all our posts to have the same set of fields in the front matter. Create a new archetype at `archetypes/posts.md` with the following content:

```md
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
tags:
    - web
    - devops
    - linux
    - programming
---
```

This is similar to `archetypes/default.md` with an extra field called `tags`. We've also provided a list of 4 common tags as suggestions, which the author can delete and add to as appropriate.

You can also add an `author` field to the archetype, but since this will be my own personal blog, and I don't plan on having guest posts, I have left that field out.
