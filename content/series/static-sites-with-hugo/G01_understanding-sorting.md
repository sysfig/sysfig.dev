---
title: Understanding Sorting in Lists
slug: list-sorting
date: 2020-12-15T16:07:57-08:00
chapter: g
order: 1
draft: true
---

Currently, we only have a single post in our section list page. But what order will posts appear in when we add more? Ideally, we want our posts to be sorted in reverse chronological order, so that the latest posts appear first.

First, let's create a new post with a different date.

```
hugo new posts/what-is-a-mac-address.md
```

Once again, the content file will be generated from the `posts` archetype.

```
---
title: "What Is a Mac Address"
date: 2020-12-13T16:26:35-08:00
draft: true
tags:
    - web
    - devops
    - linux
    - programming
---
```

Open the content file, add in some dummy content, and modify the front matter, ensuring the `date` front matter variable uses a different date to our other article.

```
---
title: "What is a MAC Address?"
date: 2020-12-13T16:26:35-08:00
draft: false
tags:
    - web
    - networking
---

## Inquit laudemque altissima mirere cum dixit ducere

Lorem markdownum Clymenen. Ista sunt gestare audet, ante, ne delicuit fluctus...
```

If we return to the `/posts/` page, we can see that Hugo has already sorted the posts by date.

![](posts-list-two-posts.png)

This is because Hugo uses some defaults on how it [orders list items](https://gohugo.io/templates/lists/#order-content). It will try to sort by date (reverse chronologically), then if the date is exactly the same, then by link title (alphabetically), then file path (alphabetically).

You can make it more explicit by writing <code>{{ range .Pages<mark>.ByDate.Reverse</mark> }}</code> instead.

You can also override the sort order by giving each page a numeric, non-zero _weight_ (via the `weight` front matter), where the page with the lower weight is ordered first. If the weights of two posts are the same, then Hugo will revert to sorting by date, title, and path.
