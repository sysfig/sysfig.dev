---
title: Adding Lists to Home Page
slug: home-page-list
date: 2020-12-15T16:20:36-08:00
chapter: z
order: 00
draft: true
---

range first 3 (where .Site.RegularPages "Type" "in" "posts").ByDate.Reverse

range first 3 (where .Site.RegularPages "Type" "in" "series").ByDate.Reverse

range first 3 (where .Site.RegularPages "Type" "in" "projects").ByDate.Reverse