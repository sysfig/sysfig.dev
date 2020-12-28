---
title: Tips
slug: tips
date: 2020-12-15T16:17:54-08:00
chapter: z
order: 00
draft: true
---

## Put as much information as possible in the front matter

You should put as much data as possible in the front matter block since each value is named and thus conveys meaning. The unstructured content that follows can only be treated as a single string - it'll be much harder to extract information from that string later on.

As an example - instead of writing the title and author in the unstructured content portion, extract that information out and put it in the metadata block. So not this:

```
---
draft: true
---
How to Build a Blog with Hugo by Daniel Li

...
```

But this:

```
---
title: How to Build a Blog with Hugo
author: Daniel Li
draft: true
---

...
```
