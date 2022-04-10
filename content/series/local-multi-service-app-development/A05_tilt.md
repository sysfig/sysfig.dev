---
title: Automating Set-Up with Tilt
slug: tilt
date: 2021-03-04T10:21:48-08:00
chapter: a
order: 4
tags:
    - kubernetes
    - tilt
draft: true
---

A `Tiltfile` is a configuration file written in a dialect of Python called [Starlack](https://github.com/bazelbuild/starlark).

```python
docker_build('localhost:54934/echo', './echo')
k8s_yaml('echo/k8s.yaml')
```

```console
% tilt up
Tilt started on http://localhost:10350/
v0.18.12, built 2021-03-09

(space) to open the browser
(s) to stream logs (--stream=true)
(t) to open legacy terminal mode (--legacy=true)
(ctrl-c) to exit
```

Hit `Space` to open the Tilt front-end interface (the Tilt UI) in your default browser.
