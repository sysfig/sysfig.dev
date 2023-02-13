---
title: Security
slug: security
date: 2020-12-15T16:17:54-08:00
chapter: z
order: 99
draft: true
---

All modern major browsers enforces the _same-origin policy_, which ensures any data stored with one of the aforementioned browser storage APIs can only be read from and written to by the same origin. Different origins (e.g. `https://foo.com` and `http://foo.com`) cannot read each other's data.

## XSS

Most storage methods are vulnerable to _Cross-Site Scripting_ (XSS) attacks. XSS attacks are a type of _code injection_ vulnerability, in which a vulnerable web application runs takes user input without validating or sanitizing it, and injects the user input onto the web page the victim is viewing. Alternatively, the site can also include third-party scripts that are malicious or have been compromized. When the user input or malicious script is executed on the victim's browser, because the script is ran within the application, it is able to read the data saved in cookies, Web Storage, and IndexedDB.

"""Cookies, when used with the HttpOnly cookie flag, are not accessible through JavaScript, and are immune to XSS"""(Cookies, when used with the HttpOnly cookie flag, are not accessible through JavaScript, and are immune to XSS)
