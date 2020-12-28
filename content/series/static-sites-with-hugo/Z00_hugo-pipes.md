---
title: Preprocessing Sass with Hugo Pipes
slug: sass-hugo-pipes
date: 2020-12-21T11:46:39-08:00
chapter: z
order: 00
draft: true
---

## Maintainability vs Performance

As you add more elements and pages to your site, the size of your stylesheet is likely to increase. At some point, it may become too unwieldy to have all styles for all elements in one big file. At this point, you may want to split the stylesheet into multiple ones, and use either multiple `<link>` elements in the HTML or the `@import` at-rule inside the CSS file. Both requires fetching stylesheets separately and is not good for performance. With HTTP/1.x, each of those fetches requires a new TCP connection, which takes time to establish and costs resources for the server to keep alive. But even with HTTP 2.0 multiplex connections, where multiple requests and responses can reuse the same TCP connection, the server still have to process multiple requests.

So it's better for server performance if all our styles are in a single stylesheets, but it's better for maintainability if our styles are separated into multiple, more manageable ones.

The solution is to split our stylesheets into multiple files, but combine them into one during the building process. This way, it's maintainable and performant.

## Webpack vs Hugo Pipes

There are tools like Webpack which are designed for bundling assets and building websites from source code, but if you don't need the whole repertoire that these tools provide, Hugo also comes with its own special set of asset processing functions called _[Hugo Pipes](https://gohugo.io/hugo-pipes/)_.

The benefit of Hugo Pipes over tools like Webpack is that it is much easier to learn - it's simply a set of functions that you add to the layout templates. You put assets you want to process in the assets directory (defaults to `assets/` but can be overriden with the `assetDir` site configuration variable), and use one or more of Hugo Pipes' functions to process them.

## Setting Up Sass Build Pipeline with Hugo Pipes

There are a few options for splitting CSS stylesheets into multiple files and combine them at build time; here, we are going to use a popular CSS preprocessors called Sass. Sass has the [`@import`](https://sass-lang.com/documentation/at-rules/import) and [`@use`](https://sass-lang.com/documentation/at-rules/use)) functions that allows us to combine stylesheets.

There are two syntax you can use with Sass (the preprocessor) - SASS (the syntax) and Sassy CSS (SCSS). SCSS is the newer syntax that's also an extension of CSS. This means that any valid CSS is also valid SCSS. Thus, we are going to use SCSS syntax, as this makes migration from CSS to SCSS much easier.

First, let's set up the pipeline before we make any changes to our stylesheet. Create the directories `assets/`, move the `style.css` into the directory, and rename it `style.scss`.

```
mkdir assets
mv static/style.css assets/style.scss
```

Then, replace our previous `<link>` element with these two lines.

<pre><code>
    <mark>{{ $style := resources.Get &quot;style.scss&quot; | resources.ToCSS }}</mark>
    &lt;link rel=&quot;stylesheet&quot; href=&quot;<mark>{{ $style.Permalink }}</mark>&quot;&gt;
&lt;/head&gt;
</code></pre>

We are first using `resources.Get` to retrieve the source file from `assets/`, and then we are piping the output (using [pipes from Go templating](https://gohugo.io/templates/introduction/#pipes)) to the Hugo Pipes function [`resources.ToCSS`](https://gohugo.io/hugo-pipes/scss-sass/), which will use the Sass preprocessor to process it to CSS.

> Some Hugo Pipes functions have sorted aliases. For example, `resources.ToCSS` is aliased to `toCSS`. So we could have also used `{{ $style := resources.Get "style.scss" | toCSS }}`.

The result is saved in a variable called `$style`. In the next line, we are replacing the `{{ "style.css" | relURL }}` portion (which gets the `style.css` file from `static/`) with a reference to `$style.Permalink`.

If we build the site now with `hugo`, you'll find the processed stylesheet in `public/` with the same name as the original from `assets/` (you can override this using the `targetPath` option to `toCSS`).

```
$ hugo
Start building sites … 

                   | EN  
-------------------+-----
  Pages            | 75  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     | 12  
  Processed images |  0  
  Aliases          |  0  
  Sitemaps         |  1  
  Cleaned          |  0  

Total in 71 ms

$ ls public 
about           img             index.xml       series          style.css
categories      index.html      posts           sitemap.xml     tags
```

> Note that to ensure build time is kept to a minimum, if there are no references to an asset in the layout templates (i.e. using `.Permalink` or `.RelPermalink`), then that asset will not be processed since it will not be used.

Now we can split out stylesheets into multiple ones and import them into `style.scss`.

> Note that Hugo prior to v? uses libsass (C/C++ port of the original [Ruby Sass](https://sass-lang.com/ruby-sass)), which has been [deprecated](https://sass-lang.com/blog/libsass-is-deprecated) by the Sass team as of 26 Oct 2020. Hugo post-v? uses [Dart Sass](https://sass-lang.com/dart-sass) as the default. Dart Sass supports the [Sass module system](https://sass-lang.com/blog/the-module-system-is-launched)
