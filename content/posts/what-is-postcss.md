---
title: What is PostCSS?
date: 2020-12-25T15:11:08-08:00
draft: true
tags:
    - web
    - programming
    - frontend
---

## Short History on CSS Pre-Processors and Post-Processors

Many front-end developers are familiar with CSS preprocessors like [Sass](https://sass-lang.com) and [Less](http://lesscss.org/), which allows you to write in a more feature-rich superset of CSS (e.g. Sassy CSS, or SCSS), and the preprocessor will convert it back to standard-compliant CSS during build time.

When preprocessors first came out (Sass in 2006), they introduced variables, loops, mixins, etc. The preprocessors' impact on front-end development are substantial. Some features, like variables, are now part of the CSS standard (e.g. variables are implemented as [CSS custom properties](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties)).

Post-Processors arrived later, with [PostCSS](https://postcss.org/) in 2013. Whereas preprocessors compiles a language (e.g. SCSS) to CSS, postprocessors takes plain CSS and optimizes it further. Postprocessors were used to automate repetitive tasks such as:

- extending class selectors
- auto-appending prefixes

## Blurring the Lines between Pre- and Post-Processors

However, as preprocessors became more developed, its features overlapped with much of those performed by the postprocessors (and if they don't, there's probably plugins/extensions available which do).

Similarly, postprocessors like PostCSS also started to support future CSS syntax, stepping into the territories of the preprocessors. Furthermore, whilst postprocessors were originally named because they usually join the pipeline after preprocessors, they can work on CSS even if it didn't come from a preprocessor.

All this is to say that the terms 'preprocessors' and 'postprocessors' are quickly becoming irrelevant - both should just be called CSS processors.

Authoring and Optimizing - https://medium.com/@ddprrt/deconfusing-pre-and-post-processing-d68e3bd078a3

## PostCSS

https://www.smashingmagazine.com/2015/12/introduction-to-postcss/
