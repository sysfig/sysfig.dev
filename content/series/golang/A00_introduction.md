---
title: Introduction
slug: introduction
date: 2020-12-30T00:00:57-08:00
chapter: a
order: 0
tags:
    - golang
---

[Go](https://golang.org/) is an open-source, general-purpose programming language designed by Google. It is typically compiled.

Installing Go gives you access to the `go` and `gofmt` tools on the command line, as well as the [standard library](https://golang.org/pkg/#stdlib) of packages. You'll use the `go` tool to build and test your code.

## Code Organization

Go code is organized into packages. Every source file must belong to a package.

## Types of Packages

There are two types of things you can create with a package - a _library_ that can be imported by another package, or a special `main` package that builds into an executable binary (i.e. a program); when building a binary with the `go` tool, Go will build the `main` package only.

## Creating an Executable

To create an executable binary that prints `Hello, World!` to `stdout` when ran, create a new file inside a new directory, and paste in the following:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

The first line `package main` declares this file to be part of a package named `main`. Every file in the Go source code must have a `package` declaration at the top of the file. `main` is a special name that tells the `go` build tool to start here when building the binary.

The next line (`import "fmt"`) imports the [`fmt`](https://golang.org/pkg/fmt/) package as a dependency into the `main` package. The `fmt` package exports a [`Println`](https://golang.org/pkg/fmt/#Println) function, which is used within the `main()` function to print the `"Hello, World!"` message to the console.

The `func main() { ... }` function is the entry point into the executable; in other words, when you've run the built binary, the execution of the code would start here.

After you've saved the file, you can run `go XXX` which will build the binary. Then you can run the binary and it will print out `Hello, World!` into the console.

## Creating a Library

The `fmt` package that our 'Hello World' example uses is part of Go's standard library. But you can also create your own library package or use a third-party package. Let's see what it takes to create your own library.





A library is a package that exports named entities such as types, functions, variables, and constants. You can make an entity exported simply by naming it with an initial capital letter. So `func multiply() { ... }` within the `math` package would only be available to other source files from the same (`math`) package, and not be available to packages that import the `math` package. `func Multiply() { ... }`, on the other hand, would.

If you are writing a library that you intend to publish for others to use, take care to export as little as necessary - it's much easier to expose more later than to drop support for exported entities that people are already using.

When creating a library package, also make sure to give the directory that houses the package source code the same name as the package (i.e. source code for the `foo` package should live inside a directory called `foo`). Although this is not a hard requirement, it is a well-established convention.

You should only group types, functions, constants, and variables together in the same package if they are related. Packages should be small and fulfill a specific purpose. If a package contains code that does multiple jobs, or if it's conceivable that a user may only need to use a small part of the package, then you may want to consider breaking the package into small packages.

## Importing Packages

The `package` statement defines the _name_ of the package a source file belongs to. However, apart from the packages in the standard library, you cannot import a package by its name. This is because Alice and Bob may both create a package named `foobar`, so if you reference a package by only its name, Go may not be able to deduce which package to use.

Instead, when you import packages, you specify the package's _path_. The path is the location of the package's source code relative to `$GOROOT/src` (for packages in the standard library) or `$GOPATH/src` (for other packages). The `go` build tool uses the path to locate the package to import.

`$GOROOT` is where your Go installation lives on your local machine; by default, it's `/usr/local/go` on Linux and macOS and `C:\Go` on Windows. So if you see the statement `import runtime/pprof`, it's references the [`runtime/pprof`](https://golang.org/pkg/runtime/pprof/) package that's found at `/usr/local/go/src/runtime/pprof`. This means even if there are two packages in the standard library with the same name (e.g. `pprof` - `runtime/pprof` and [`net/http/pprof`](https://golang.org/pkg/net/http/pprof/), or `template` - [`html/template`](https://golang.org/pkg/html/template/) and [`text/template`](https://golang.org/pkg/text/template/)), they can be uniquely-referenced using the path.

When creating your own package locally, or downloading a third-party package, the convention is to use the domain-plus-path of the repository where the source code lives (e.g. `github.com/dgrijalva/jwt-go`) as the path. Because the publisher of the package have control over this path, this helps avoid name conflicts. It also has the side effect of organizing the packages inside `$GOPATH/src` in an sensible way.

If your package is not hosted externally, then you can prefix the path with a unique identifier such as a domain name you control, or your company name.

Once you have imported a package by specifying its path, you can now, in your code, reference its _exported identifiers_ (those that starts with a capital letter) using the syntax `<package-name>.<entity-name>` (e.g. `foo.Bar` references the `Bar` exported name from the package `foo`). Code _within_ a package can refer to any other identifiers defined within the package.

There are some more rules about importing packages that you should know:

- You can import packages with the same name as long as you assign them _local name(s)_ to resolve any name conflict(s)
- You cannot have circular imports (`A` imports `B` but `B` also imports `A`)

## Workspace, `$GOPATH` and `go get`

Go 1.0 introduced `$GOPATH` based workspaces.

`go get` is a tool that fetches packages from a remote source and add it to your workspace. Unlike package managers such as JavaScript's npm, it's designed to be decentralized - without needing a central repository.

Package versioning is not incorporated into the `go get` tool.

## Versioning Go Packages

However, using just the name and path to reference a package is not enough. Packages change over time - new features, bug fixes, some of them incompatible, may get pushed to the package repository at any time. When your code depends on a package, it should depend on a specific version of a package, not the latest version.

When Go was first created, it didn't have an official way of versioning packages, which led to many community-driven solutions popping up to fill this void. [Godep](https://github.com/tools/godep), [gpm](https://github.com/pote/gpm), [gom](https://github.com/mattn/gom), [gb](http://www.getgb.io/), [govendor](https://github.com/kardianos/govendor), [gvt](https://github.com/FiloSottile/gvt), [gopm](https://github.com/gpmgo/gopm), [Glide](https://glide.sh/), and [Dep](https://github.com/golang/dep) were some of the more popular ones.

dep and glide uses vendoring

	"GLOCKFILE":          ParseGLOCKFILE,
	"Godeps/Godeps.json": ParseGodepsJSON,
	"Gopkg.lock":         ParseGopkgLock,
	"dependencies.tsv":   ParseDependenciesTSV,
	"glide.lock":         ParseGlideLock,
	"vendor.conf":        ParseVendorConf,
	"vendor.yml":         ParseVendorYML,
	"vendor/manifest":    ParseVendorManifest,
	"vendor/vendor.json": ParseVendorJSON,

Dep was released in January 2017

## Backwards Compatibility

Go 1.0 also issued another guideline.
Another attempt at resolving issues with package incompatibility is to encourage package publishers to remain backwards compatible.

Since Go 1.0, packages should be backwards compatible. This means new versions of packages must be compatible with older versions of the package using the same import path. This means you should not remove a previously-exported name; the best you can do is deprecate it. If an incompatible change must be introduced, then the new package should use a different import path. This convention is known as _import compatibility_, or the _import uniqueness rule_.

A convention that arose from this was _semantic import versioning_ (SIV), where newer, incompatible versions of a package are given a new major version number according to [Semantic Versioning 2.0.0](https://semver.org/) rules, and the version is included in the import path.

SIV satisfies the import uniqueness rule, and it also allows multiple versions of the same package to be included in the same program.

?? How does that work ??

encourage tagged composite literals
https://golang.org/doc/go1compat

But backwards compatibility is a well-established convention, not a guarantee. There's still a chance that the package author may not abide by it. So you'd still have to rely on vendoring to ensure you have reproducible builds.

### Vendoring

SOME? relied on the existence of a `vendor/` directory.

The ??go1.5+?? vendor folder


With modules, running `go mod vendor` will create/update the `vendor/` directory with information from the `go.mod` and/or `go.sum` files. When building your code, the `go` tool will only use the packages that are downloaded to the `vendor/` directory.


Create a local copy of the package's source code in a directory called `vendor/`. That copy would act as a frozen snapshot of the package. https://go.googlesource.com/proposal/+/master/design/25719-go15vendor.md

This is not ideal for several reasons:

- If multiple projects share the same dependency, then that dependency is duplicated for each project.
- It does not cater for the case where a dependency is unavailable but you do not have a local copy (perhaps you're building from a new machine after your old machine crashed). A solution of this is to also commit the `vendor/` repository into your repository, but that means you are now adding code that is not specific to your application/library into the repository.


## vgo


vgo

vgo introduced the idea of a Go module (https://research.swtch.com/vgo-module)

https://research.swtch.com/vgo

a Go module is a collection of packages versioned as a unit

enabling work outside $GOPATH

https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md
https://pkg.go.dev/golang.org/x/vgo




## Go Modules

A Go module is a collection of Go packages stored in a file tree with a `go.mod` at the root. The `go.mod` file looks like this:

```
module github.com/sysfig/auth

go 1.15

require (
  github.com/gin-gonic/gin v1.6.3
  github.com/rs/zerolog v1.19.0
)
```

The `module` keyword defines the module's _module path_ - the _import path_ of the root directory ??. This means if the module source code is located at `$HOME/projects/auth/`, then a package within that module at `$HOME/projects/auth/twitter` can ??

The `require` keyword defines the module's dependencies. Each dependency is specified using its module path as well as a semantic version string.


In August 2018, [Go 1.11 added preliminary/experimental support for Go modules](https://golang.org/doc/go1.11#modules). `go get` now is aware of modules.

The go command defaults to module mode when run in directory trees outside GOPATH/src and marked by go.mod files in their roots. This setting can be overridden by setting the transitional environment variable $GO111MODULE to on or off; the default behavior is auto mode.

The `go` tool is now aware of modules, and will perform the task of managing dependencies.

Altnernative to having all your Go code inside `$GOPATH`.
Migrate from GOPATH to modules


https://golang.org/doc/go1.11#modules
https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more


Go 1.12 - commands like `go run x.go` or `go get rsc.io/2fa@v1.1.0` can now operate in `GO111MODULE=on` mode without an explicit `go.mod` file.

Go 1.13?? enable Go modules by default

Many tools still assumes to find packages inside `$GOPATH`. Go tool authors should use the https://pkg.go.dev/golang.org/x/tools/go/packages package which is capable of using both of the older `$GOPATH`-based approach and the new Go modules approach.

`go.sum` is _not_ a lock file. It stores the checksums (product of a cryptographic hash function) of the contents of a module's dependencies. The file is used by the `go` tool to verify the integrity of the package source code it downloads from the remote repository.

Go 1.13 - Go notary will act as a global `go.sum` when adding/updating modules.

## Module Proxy

The decentralized nature of the `go get` command means it does not need a centralized location where all packages are stored. However, this also meant that `go get` may have to download from dozens, or even hundreds of different servers. If any of these servers are unavailable (if the repository is deleted, renamed, relocated), or otherwise unreliable (e.g. if GitHub is temporarily down), then `go get` may not be able to retrieve the dependency and the build will fail.

The design of Go modules included support for a _module proxy_. A module proxy is a server that downloads modules from the origin servers, caches them, and serves them in a central location (acting as a _mirror_). The module mirrors would allow modules to be available even when the origin server is down; multiple mirrors can be used in case any one of the mirrors goes down.

Note that this is still not the same as a central repository where module publishers publish to, the module proxy only acts as a mirror server to provide better reliability. Publishers still publishes it wherever they want.

Google runs its own 

https://golang.org/cmd/go/#hdr-Module_proxy_protocol

Self-hosted proxies using open source
- https://github.com/gomods/athens
- https://github.com/goproxy/goproxy
GoCenter

Commercial:

- There's a Go module proxy hosted by Google


Go module proxy (https://proxy.golang.org/) and checksum database are used by the `go` tool by default since 1.13

### Finding Modules



https://pkg.go.dev/ (replaces godoc.org)

Go Module Index ??how does it work?? (https://index.golang.org/)

Has a log that updates with new packages

This log will feed into the notary (https://sum.golang.org/) as well as mirrors. The mirrors can use this log to download and cache new and updated modules.

https://search.gocenter.io/
