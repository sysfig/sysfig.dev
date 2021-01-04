---
title: config.toml
slug: config-toml
date: 2020-12-14T20:20:48-08:00
chapter: b
order: 1
draft: true
---

To start off, we have the `config.toml` configuration file.

```toml
baseURL = "http://example.org/"
languageCode = "en-us"
title = "My New Hugo Site"
```

`config.toml` holds variables that can be used in all templates. Typically, these would be site-specific information like the site's name, canonical domain name, language code, theme, etc.

By default, the configuration is written in _[Tom's Obvious, Minimal Language](https://toml.io/en/)_ (TOML), named after Tom Preston-Werner, the co-founder of GitHub and Jekyll. But if you're like me and prefer to write configuration files in _[YAML Ain't Markup Language](https://yaml.org/)_ (YAML), then you can delete the whole directory and re-create the site using the `--format` option in the `hugo new site` command.

```
hugo new site sysfig.dev --format yaml
```

The rest of this book will use YAML for the configuration file. Now let's look at what these three configuration variables mean:

- `baseURL` - used to construct absolute URLs. This should be the website's canonical domain name.
- `languageCode` - For accessibility to allow screen readers to pronounce the contents of the page as well as the page's `<title>` element appropriately. It should be a [valid IETF identifying language tag](https://www.ietf.org/rfc/bcp/bcp47.txt)

You can find a list of all valid configuration settings on the [_Configure Hugo_](https://gohugo.io/getting-started/configuration/#all-variables-yaml) page.
