---
title: Overriding Archetype
slug: overriding-archetype
date: 2020-12-15T20:51:16-08:00
chapter: i
order: 2
---

With picking the right archetype, we can use the `--kind` option to specify the content type to create.

> There's some inconsistency here with how the `--kind` option is named. It really should be named `--type` since it is used to specify content type. Page kinds should only used to mean one of home, regular page, section, taxonomy, taxonomy term).

Delete the test file we created earlier and use the `hugo new` command with the `--kind` option to recreate it.

```
rm -rf content/series/generate-static-sites-with-hugo
hugo new series/generate-static-sites-with-hugo/test.md --kind posts
/Users/sysfig/projects/sysfig.dev/content/series/generate-static-sites-with-hugo/test.md created
```

Now, it will use the correct archetype.

```yaml
---
title: Test
date: 2020-12-15T20:39:43-08:00
draft: true
tags:
    - web
    - devops
    - linux
    - programming
---
```

This solves the issue of setting the correct type when picking an archetype; but if we build the site now, the test page will still try to use the layout at `layouts/series/single.html`, and defaulting to `layouts/_default/single.html` if it doesn't exists.

This is because a page's content type defaults to the root section it is in; we used `--kind` to override it when calling the `hugo new` command, but that is only temporary.
