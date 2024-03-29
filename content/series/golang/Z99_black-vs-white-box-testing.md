---
title: Black- vs White-box Testing
slug: black-vs-white-box-testing
date: 2020-12-30T16:14:33-08:00
chapter: z
order: 99
tags:
    - golang
    - testing
draft: true
---

If you test code is in the same package as the functions you are testing, then your test code is able to access all the private (non-exported) methods. This allows you to perform blackbox unit testing on every unit in your package.

You can also have some other tests that tests the public methods and interface that it exposes. These tests should be in a different package (usually named `<package>_test`) to the package you are testing because you want to test it as if your test package is a third-party package, without access to the private methods. You place the test code in a different package and import the package to test just like a third-party package would. This will help prevent you from forgetting to export methods.
