---
title: data/
slug: data
date: 2020-12-14T20:20:26-08:00
chapter: b
order: 3
draft: true
---

There may be portions of your layout that has a list structure. For example, for a developer's personal blog, their `projects/` page may be a list of GitHub repositories. Whilst you can copy and paste the same HTML template for every list item, it'll be more DRY to have a single HTML snippet and use a loop to iterate over the list of repositories.

This is what the `data/` directory is for - it allows you to store data in JSON, YAML, or TOML, and inject the data into layouts.
