---
title: Generating a Project Scaffold
slug: generating-a-project-scaffold
date: 2020-12-14T17:52:52-08:00
chapter: b
order: 0
---

Let's create a scaffold for a new site using the `hugo new site <path>` command.

```
$ hugo new site sysfig.dev
```


```
Congratulations! Your new Hugo site is created in /Users/sysfig/projects/sysfig.dev.

Just a few more steps and you're ready to go:

1. Download a theme into the same-named folder.
   Choose a theme from https://themes.gohugo.io/ or
   create your own with the "hugo new theme <THEMENAME>" command.
2. Perhaps you want to add some content. You can add single files
   with "hugo new <SECTIONNAME>/<FILENAME>.<FORMAT>".
3. Start the built-in live server via "hugo server".

Visit https://gohugo.io/ for quickstart guide and full documentation.
```

`hugo` will create a new directory called `sysfig.dev` and place a skeleton boilerplate inside.

```
$ tree sysfig.dev/
sysfig.dev/
├── archetypes
│   └── default.md
├── config.toml
├── content
├── data
├── layouts
├── static
└── themes
```

Let's go through each of these files and directories, and along the way, explain important concepts you'd need to know to work with Hugo.
