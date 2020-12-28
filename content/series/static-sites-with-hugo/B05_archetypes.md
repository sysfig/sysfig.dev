---
title: archetypes/
slug: archetypes
date: 2020-12-14T20:21:32-08:00
chapter: b
order: 5
---

Chances are, your website or blog would have a limited number of content types. For a professional blog, you may have tutorials, opinion/discussion/analysis posts, announcements, etc. For a personal blog, it may be photography, food, travel, etc.

Posts of a similar type tend to have a similar structure (i.e. the same set of metadata and content structure). For example, all photography posts may share the metadata fields `date`, `time`, `location`, `lat`, `lon`, `photographer`. Food recipe posts may share the metadata fields `name`, `cuisines`, `ingredients`, `prep_time`, `cook_time`, `servings`, `calories`, `photos`; their content may also usually include the headings `Introduction`, `Instructions`, and `Notes`.

Instead of writing out the same metadata and content skeleton for each new post, you can define a prototypic content file called an _archetype_. A archetype defines all the common metadata fields and content skeleton, and may pre-fill metadata values such as `date` (set to current time) and `draft` (defaults to `true`).

Whenever you use `hugo new [contentPath]` to create new content, Hugo will bootstrap that content file with a copy of an archetype file.

Archetypes provides the skeleton for new content files. Archetypes exists because writing out the same set of metadata and content skeleton for every content file is tiresome.

> Archetypes are not schemas - once you generated a content file from an archetype, you are free to modify the content file as much as you like. If you come from the programming world, an archetype is more akin to a _prototype_.

## `archetypes/default.md`

The `hugo new site <path>` command creates a minimal archetype at `archetypes/default.md`. It consists of a metadata block with 3 properties, and no unstructured content.

```
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
---
```

The `{{}}` is ??Go templating syntax?? which allows you to access variables that are injected into the file (e.g. `Name` and `.Date`), and call Go functions (like `replace` and `title`).

Let's now run `hugo new [contentPath]` to create a new content file.

```
$ hugo new test.md
/Users/sysfig/projects/sysfig.dev/content/test.md created
```

If we look into that file, we can see the same 3-property metadata block but its values have been evaluated and filled in for us.

```
---
title: "Test"
date: 2020-12-05T17:16:51-08:00
draft: true
---
```
