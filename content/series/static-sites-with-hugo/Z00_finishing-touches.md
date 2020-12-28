---
title: Adding Permalink Copy Button
slug: permalink-copy-button
date: 2020-12-15T16:20:06-08:00
chapter: z
order: 00
draft: true
---

`.Permalink`

compared with .RelPermalink

https://gohugo.io/content-management/urls/

# Adding a Last Updated Date

`.Lastmod` 

"""If lastmod is not set, and .GitInfo feature is enabled, .GitInfo.AuthorDate will be used instead."""

# Adding Table of Contents to Post

`.TableOfContents`
https://gohugo.io/getting-started/configuration-markup/#table-of-contents

# Adding Prev/Next Links to Post

{{with .Next}}{{.Permalink}}{{end}}

""".NextInSection
    Points up to the next regular page below the same top level section (e.g. in /blog)). Pages are sorted by Hugo’s default sort. Example: {{with .NextInSection}}{{.Permalink}}{{end}}. Calling .NextInSection from the first page returns nil."""

.Prev
    Points down to the previous regular page (sorted by Hugo’s default sort). Example: {{if .Prev}}{{.Prev.Permalink}}{{end}}. Calling .Prev from the last page returns nil.
.PrevInSection
    Points down to the previous regular page below the same top level section (e.g. /blog). Pages are sorted by Hugo’s default sort. Example: {{if .PrevInSection}}{{.PrevInSection.Permalink}}{{end}}. Calling .PrevInSection from the last page returns nil.

# Adding Word Count `.FuzzyWordCount` / `.WordCount`/`.ReadingTime`
# 404

# Showing Last Commit on bottom of each page