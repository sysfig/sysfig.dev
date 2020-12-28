---
title: Previewing the Site with the Development Server
slug: previewing-site-with-development-server
date: 2020-12-15T15:24:24-08:00
chapter: c
order: 5
---

It's not possible to serve our site by running a web server (like NGINX or Apache HTTP server) with `public/` as the site root. However, running a web server in the background and re-running `hugo` on every change can tiresome.

Hugo provides a fast development server that will watch for changes in your code, automatically rebuild the site, store the result in memory, and server it at the local host address (defaults to port `1313`). It also has a hot-reload feature, where it will inject some code into web page that'll allow the page to update itself without refreshing whenever the code changes. Generally, when developing a Hugo site, you should use this development server.

Start the development server by running `hugo server`.

```
hugo server
```

```
Start building sites …

                   | EN  
-------------------+-----
  Pages            |  4  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     |  0  
  Processed images |  0  
  Aliases          |  0  
  Sitemaps         |  1  
  Cleaned          |  0  

Built in 2 ms
Watching for changes in /Users/sysfig/projects/sysfig.dev/{archetypes,content,data,layouts,static}
Watching for config changes in /Users/sysfig/projects/sysfig.dev/config.yaml
Environment: "development"
Serving pages from memory
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop
```

Our site is now available to view at http://localhost:1313/.

![](/img/hello-world-home-page-color.png)

Although it's not a very useful home page, this 'Hello World' exercise has taught you how Hugo works and how content and layouts come together to form the page.
