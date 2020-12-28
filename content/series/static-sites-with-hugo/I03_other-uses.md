---
title: Other Uses of Taxonomy
slug: taxonomy-use-cases
date: 2020-12-15T16:11:04-08:00
chapter: h
order: 3
---

We are only using one taxonomy here (`tags`), so let's update our site configuration (`config.yaml`) to make that explicit.

```yaml
...
taxonomies:
  tag: tags
```

This will prevent categories from becoming a taxonomy, and prevents Hugo from generating a `/categories/` page.

Taxonomies are useful, you may want to consider using taxonomies to group:

- different categories of posts (e.g. tutorial, analysis, announcement, review)
- by author
