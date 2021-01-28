---
title: Developing Multi-Service Applications Locally with Kind and Tilt
date: 2021-01-27T17:37:51-08:00
type: series
layout: single
prerequisites:
  required:
    - Some programming knowledge
    - Basic command of the terminal/command-line tools
    - Access to a code editor like VSCode
  recommended:
    - Programming experience in Go
learning-goals:
  - 
out-of-scope:
  - Installing Go (follow https://golang.org/doc/install)
what-you-will-build:
  - name: Echo Server
    description: An echo server is a simple API server that responds with the same string that the client passes in the request
software:
  supported:
    - name: Go
      versionStart: 1.4
      versionEnd: 1.15.6
  tested:
    - name: Go
      version: 1.15.6
tags:
    - web
    - golang
draft: true
---

In this series, we are going to be building, testing, and deploying a multi-service application locally on our machine. We will be using [Go](https://golang.org/), [Docker](https://www.docker.com/), [Kubernetes](https://kubernetes.io/), [kind](https://kind.sigs.k8s.io/), [Helm](https://helm.sh/), and [Tilt](https://tilt.dev/). Although we won't go into each in depth, we will cover the basics from scratch. So if you have no idea how to code in Go or use any of these tools, don't worry, we will explain everything from scratch.

In the first chapter, we are going to be focusing on setting up tools that'll watch for changes in our code and update the running service(s). The quicker this happens, the quicker we see the effect of our changes. A fast feedback loop enable us to iterate quickly. Therefore, to maintain that focus, the application that we'll be building and deploying is going to be a very simple echo server.
