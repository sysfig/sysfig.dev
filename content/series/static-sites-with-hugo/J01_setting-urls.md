---
title: Setting URLs
slug: setting-urls
date: 2020-12-15T16:52:29-08:00
chapter: i
order: 1
draft: true
---

There are multiple ways to fulfil our first constraint - Posts within a series should have the URL path `/series/<series-name>/<post-slug>`:

1. **Approach**: Use filename directory structure

   By default, a page's URL path reflects its filename and the directory structure it is found in. Therefore, we can simply place posts that should be within a series inside the `content/series/<series-name>/` directory.

   **Advantages**:

   1. It's simple to implement.
   2. Doing so will fulfil our second constraint (Posts in a series should not appear on the `/posts` list page). Because series post pages are not in the `content/posts/` directory, the `{{ .Pages }}` block in the layout will not include them.

   **Drawbacks**:

   1. Because it's now within the `content/series/` directory, it will belong to the section `series` and defaults to using the `series` archetype and layouts, which is not what we want. The `series` archetype and layouts should be used for the series page (i.e. `/series/<series-name>`), not the series post pages.

2. **Approach**: Use custom naming convention

   Treat series posts as a normal post with a specific naming convention. Here's how it can work:

   1. Prefix the filenames of series posts with the kebab-case name of the series followed by and underscore (e.g. `installing-hugo.md` becomes `static-sites-with-hugo_installing-hugo.md`)
   2. Use Hugo's `replace` functions to replace the underscore with a forward slash (i.e. `{{ replace .Name "_" "/" }}`)
   3. Use the `url` front matter variable to set the URL path (i.e. `{{if in .Name "_" -}} url: /series/{{ replace .Name "_" "/" }} {{- end }}`)
   4. Update the `title` front matter template to remove the underscore and everything before it before piping to the `title` function (i.e. `title: {{ replace .Name "-" " " | replaceRE "^.*_(.*)" "$1" | title }}`)

   **Advantages**:

   1. It is in the `posts` section and so no extra work needs to be done for it to use the `posts` archetype and layouts.

   **Drawbacks**:

   1. The standalone posts and series posts are all going to be located in the same directory. If you have a lot of posts, then this can make things hard to find.
   2. The file can be lengthy if both the series name and the post title are long

We are going with the first approach for its simplicity, plus there are easy workarounds to address its drawbacks. So let's create a new post at `content/series/<series-name/<series-post-title>` (e.g. `content/series/generate-static-sites-with-hugo/test.md`) by running `hugo new`.

```
hugo new series/generate-static-sites-with-hugo/test.md
```

Hugo will create the file using an archetype. It will first try to find an archetype called `series`, but since it doesn't exist, it'll default to the `_default` archetype.

```yaml
---
title: Test
date: 2020-12-15T20:01:47-08:00
draft: true
---
```

However, we want series posts to use the `posts` archetype and layout, so we need to explicitly give each series post a content type of `posts`. So how can we set the `type` property for every post within every `content/series/*/` directory?
