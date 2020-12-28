---
title: {{ replace .Name "-" " " | replaceRE "^(.+_)?(.*)" "$1" | title }}
slug: {{ replaceRE "^(.+_)?(.*)" "$1" .Name | lower }}
date: {{ .Date }}
chapter: {{ replaceRE "^([[:alpha:]]+).*" "$1" .Name | lower }}
order: {{ int (replaceRE "^[[:alpha:]]+(\\d+)_.*" "$1" .Name | strings.TrimLeft "0" | default 0) }}
tags:
    - web
    - devops
    - linux
    - programming
draft: true
---
