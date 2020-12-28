---
title: Static vs Dynamic Sites
slug: static-vs-dynamic-sites
date: 2020-12-14T17:07:12-08:00
chapter: a
order: 1
tags:
    - web
    - hugo
    - wordpress
    - ghost
---

Some open-source blogging software like [WordPress](https://wordpress.org/) and [Ghost](https://ghost.org/) store post data (content and metadata) as database entries, and retrieve them _dynamically_ at the time clients request them. For example, if a visitor navigates to a post on a Ghost blog using the URL `https://blog.example.com/explaining-dns`, the Ghost backend would take the path (`/explaining-dns`) and try to match it to an entry in the database. Once found, it will either return the post data to the front-end to display, or generate the HTML page server-side and respond with the dynamically-generated HTML page.

The alternative to this approach is to create _static sites_, where the post data are stored as text files. To create the site, a tool like [Hugo](https://gohugo.io/), [Jekyll](https://jekyllrb.com/), [Gatsby](https://www.gatsbyjs.com/), or [Zola](https://www.getzola.org/) is used to read the post data and generate a set of HTML pages that make up the website. You'd then need to deploy a web server (e.g. [NGINX](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/) or [Apache HTTP Server](https://httpd.apache.org/)) to serve these static HTML pages.

The difference, then, between a dynamic site and a static site is on when the HTML page presented to the client is generated. With a static site, the HTML page is pre-generated; with a dynamic site, the HTML page is generated on-the-fly at the client's request.

## Response Time

With a static site, the web server can respond to a client's request immediately with the _pre-generated_ HTML page. With a dynamic site, the backend server must first send an internal request to the database to retrieve the post data before responding to the client. As such, a static site usually have a much shorter response time than a dynamic site.

A dynamic site can improve its response time by adding a caching layer in front of the backend server. If the most recent HTML responses from the backend is cached, the next time a client makes an identical request, the caching server would serve the cached copy instead of passing the request onto the backend, saving a lot of time.

Similarly, a static site can improve its response time by 'caching' its HTML pages and static assets (e.g. images, videos) using a _Content Delivery Network_ (CDN). A CDN is a large network of servers that sits between the client and the _origin server_ (i.e. the blog backend server), and serves content that would otherwise have been served by the origin server. By reducing the physical distance between the client and the content, it can minimize the response time.

For example, if our origin server is deployed in Japan, without CDN, a request from a client in Iceland must go across the globe to Japan, and the response times may be long. However, if our web pages are cached using a CDN that has _edge servers_ in Amsterdam, then the request only need to travel a (relatively) short distance, and response times are likely to be much shorter.

## Resource Requirements

Dynamic sites requires you, at a minimum, to deploy a database and a backend API. A static site, on the other hand, requires only a web server to serve the HTML files. When idle (no client requests), running a web server consumes far fewer system resources (CPU and memory) than running a database and an API server.

When not idle, each (uncached) request on a dynamic site requires at least a database read, which may require network I/O and some data processing. Compared with a static site, where the web server only needs to read a file and send its content in the response message. Thus, static sites also use fewer system resources per requests.

## Maintainability

With static sites, any changes to the blog configuration, theme, or posts requires the entire site to be rebuilt and redeployed. This can be a slow and tiresome process if done manually; thus, static site maintainers typically invest time upfront to set up a pipeline to automate the process. For example, one may set up a [GitHub Actions](https://docs.github.com/en/free-pro-team@latest/actions) action that generates and deploys the static site from source code each time new commits are pushed.

Most dynamic blogging platforms, on the other hand, provide a web user interface (UI) where authors, editors, and administrators can make changes to the blog without re-deploying the site. System administrators must still make an upfront time investment to set up the database and backend, but there are less ongoing costs.

The database and backends of dynamic sites may need to be upgraded from time to time to take advantage of security fixes and new features. But this is similar to static site generators, where the command-line tool used to generate the site may also need to be upgraded.

## Usability

Using static site generators usually requires the author to use a code editor, and write in a language like Markdown. In contrast, many dynamic blogging platforms, like WordPress, provide an online web UI with a _What You See Is What You Get_ (WYSIWYG) editor, which removes the requirement for authors to know HTML or Markdown.

A benefit of using a static site generator is that it's much easier to work offline - you can write in Markdown files and regenerate the entire site offline, once you're happy and back online, you simply have to upload the new set of files.

To make configuration changes with a dynamic framework, you either have to make live changes on the production site, or run a development/test site, make the changes there, and replicate the changes to the live site. Most dynamic framework also does not allow you to work offline (without a non-trivial risk of losing your work).

However, dynamic platforms often have features that allows you to schedule the publication of posts. To implement this feature with a static site would require external tooling that supports running scripts on a schedule. The script is used to change the draft status of the content file, regenerate the site, and upload the files to the server.

Which one is more usable depends on who the author is - a software engineer may prefer to write using their editor and are familiar with Markdown, a journalist may prefer the WYSIWYG editor and the web UI because they can work from any machine with access to the Internet.

## Permissions

For sites built using a static site generator, controlling who can change the blog configuration, write a new article, and publish the changes depends on its underlying technologies - its version control system (e.g. Git), code repository host (e.g. GitHub), and file upload (e.g. SFTP).

For example, you can control who can change blog configuration and make new posts by controlling who can push to the main branch in Git. For those without such permissions, they may push to a new branch and submit a pull request (PR) to the main branch. Someone who has push permissions to the main branch can then review the changes and merge it into the main branch. It's more complicated, however, to allow someone to write new posts but not make changes to the site configuration (you may have place all content in a separate Git repository and use Git submodules to include it in the main repository when building the site).

Likewise, you can control permissions for deploying to production by controlling SSH access to the server.

With dynamic sites, setting permissions is much more flexible. You can break permissions granularly, and create roles like author, editor, and administrators. Any permission scheme is possible as long as you can program it.

## Conclusion

The decision of whether to use a dynamic blogging platform, like WordPress or Ghost, or static site generators, like Hugo or Jekyll, depends largely on who the author(s) are. Dynamic platforms probably the more suitable for the majority of cases - they provide a more sophisticated permissions model suitable for larger publications, and writers don't need to know Markdown. But for the portfolio sites, blogs, and small company sites maintained by a lone software engineer or small team, then the edge in performance afforded by static sites makes it the preferred choice.
