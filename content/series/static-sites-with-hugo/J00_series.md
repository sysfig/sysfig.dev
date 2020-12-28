---
title: Series
slug: series
date: 2020-12-15T10:34:49-08:00
chapter: i
order: 0
---

The term 'series' may have different meanings on different sites or blog, but for this blog, a series is an ordered selection of posts.

When implementing series, I imposed some additional constraints (i.e. personal preferences based on my writing habits):

1. Posts within a series should have the URL path `/series/<series-name>/<post-slug>`
2. Posts in a series should not appear on the `/posts` list page
3. Every series should have its own _series page_ at `/series/<series-name>/` and it should show a list of all posts that belong to it
4. Series posts should display in the correct order on the site and in the editor
5. `/series/` should list out a list of series pages, but not the series posts

With these constraints in mind, let's figure out how to implement series.
