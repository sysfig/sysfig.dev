---
title: rss
slug: rss
date: 2020-12-15T16:19:05-08:00
chapter: z
order: 00
draft: true
---

```
{{ with .OutputFormats.Get "RSS" }}{{ .RelPermalink }}{{ end }}
```