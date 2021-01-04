---
title: Implementing Custom Filename Conventions
slug: custom-filename-conventions
date: 2020-12-15T11:02:02-08:00
chapter: i
order: 5
draft: true
---

Currently, the series posts are sorted alphabetically in my editor by their filenames, whilst on the site, they are sorted by date.

The easiest way to achieve the fourth constraint (Series posts should display in the correct order on the site and in the editor) is implement a filenaming convention where series posts that should come first are first alphabetically. We can then order posts on the site using the underlying file's name.

I have my own system of ordering series posts by chapter (letter) and intrachapter order (a 2-digit number, starting at 0). This is reflected in the names of each post; for example, `B07_themes.md` would be the eighth post in the second chapter.

## Overriding Title and Slug

This ensures series post files are ordered properly on my editor, but its URL path would not be so good for SEO.

By default, the URL of a page is determined by its name and location within the `content/` directory. Therefore, a content file at `content/series/generate-static-sites-with-hugo/I00_series.md` would, by default, appear in `<host>/series/generate-static-sites-with-hugo/I00_series` However, my file-naming convention won't look too good in the URL path (nor would it be good for SEO). Therefore, we need to use the `slug` predefined front matter variable to override the path for the page.

To make it easier, we can use the [`replaceRE`](https://gohugo.io/functions/replacere/) to automatically remove the chapter and order information from the title and slug.

Since this logic doesn't need to apply to standalone posts, we should create a new archetype for series posts. Create a file at `archetypes/series-posts.md` and fill it with the following lines:

```yaml
---
title: {{ replace .Name "-" " " | replaceRE "^.+_(.*)" "$1" | title }}
slug: {{ replaceRE "^.+_(.*)" "$1" .Name | lower }}
date: {{ .Date }}
draft: true
tags:
    - web
    - devops
    - linux
    - programming
---
```

> `.Type`, `.Date`, `.Site`, `.Name`, the target content file, and all Hugo functions, are available to use within an archetype template.

In addition to `replaceRE`, we are also using the [`lower`](https://gohugo.io/functions/lower/) function to ensure the slug only contain lowercase letters.

We don't need to update `content/series/_index.md` since series posts can still use the `posts/single.html` layout.

## Ordering Series Posts in Layout

With 

https://github.com/gohugoio/hugo/blob/21fa1e86f2aa929fb0983a0cc3dc4e271ea1cc54/resources/page/pages_sort.go

`ByWeight`, `ByTitle`, `ByLinkTitle`, `ByDate`, `ByPublishDate`, `ByExpiryDate`, `ByLastmod`, `ByLength`, `ByLanguage`, `ByParam`, 

There is no `ByFilePath`. To get around that, we can add some logic in the `series-posts` archetype to add two new params (`chapter` and `order`) whose value is derived from the filename.

<pre><code>
---
title: {{ replace .Name "-" " " | replaceRE "^.+_(.*)" "$1" | title }}
slug: {{ replaceRE "^.+_(.*)" "$1" .Name | lower }}
date: {{ .Date }}
<mark>chapter: {{ replaceRE "^([[:alpha:]]+).*" "$1" .Name | lower }}
order: {{ int (replaceRE "^[[:alpha:]]+(\\d+)_.*" "$1" .Name | strings.TrimLeft "0" | default 0) }}</mark>
tags:
    - web
    - devops
    - linux
    - programming
draft: true
---
</code></pre>

We could have used a single front matter field that combines `chapter` and `order` together, but spliting it in two helps maintain their semantics (ultimately, it's a personal preference).

In addition to `replaceRE`, we are also using [`int`](https://gohugo.io/functions/int/) to parse the order string as an integer. This isn't strictly necessary, but it feels cleaner to have a numeric value be represented as a numeric type.

The `int` function does not work with strings with leading 0's, so we are also using [`strings.TrimLeft`](https://gohugo.io/functions/strings.trimleft/) to trim leading 0's. But this means `00` will be trimmed to an empty string (`""`), which `int` also rejects. Therefore, we need to use the [`default`](https://gohugo.io/functions/default/) function to set the value to `0` if `strings.TrimLeft` returns with an empty string.

We can now update the `layouts/series/single.html` layout to use the new `chapter` and `order` params to order our series posts.

<pre><code>
  <h2>Syllabus</h2>
  {{ range <mark>(.Pages.ByParam "order").ByParam "chapter"</mark> }}
    {{ if .IsPage}}
</code></pre>
