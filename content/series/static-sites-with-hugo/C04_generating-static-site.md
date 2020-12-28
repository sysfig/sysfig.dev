---
title: Generating the Static Site
slug: generating-static-site
date: 2020-12-15T15:23:55-08:00
chapter: c
order: 4
---

With the `layouts/index.html` and `content/_index.md` files in place, we are now ready to combine them to create the home page. At the root of your repository, run `hugo`.

```console
hugo
```

```output
Start building sites …

                   | EN  
-------------------+-----
  Pages            |  4  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     |  2  
  Processed images |  0  
  Aliases          |  0  
  Sitemaps         |  1  
  Cleaned          |  0  

Total in 12 ms
```

You'll find a new `public/` directory containing the contents of the generated static site.

```
$ tree public/
public/
├── categories
│   └── index.xml
├── hello-world-home-page-color.png
├── index.html
├── index.xml
├── sitemap.xml
└── tags
    └── index.xml

2 directories, 6 files
```

Look inside `public/index.html`, and you'll find the generated home page - a composite of the layout and our content.

```html
<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta name="generator" content="Hugo 0.79.0" />
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>My New Hugo Site</title>
</head>
<body>
    <header>
        <nav>
            <a href="/">Home</a>
            <a href="/series">Series</a>
            <a href="/posts">Posts</a>
            <a href="/projects">Projects</a>
            <a href="/now">Now</a>
        </nav>
    </header>
    <main>
        <p>Welcome to <code>sysfig.dev</code></p>
    </main>
    <footer>sysfig.dev &copy; 2020</footer>
</body>
</html>
```
