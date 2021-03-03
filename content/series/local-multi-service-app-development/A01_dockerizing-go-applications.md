---
title: Dockerizing Go Application
slug: go-dockerfile
date: 2021-02-03T10:54:25-08:00
chapter: a
order: 1
tags:
    - docker
    - dockerfile
    - golang
draft: true
---

In this post, we will be packaging our Go application into a Docker image. The first step to building a Docker image from source code is creating a [`Dockerfile`](https://docs.docker.com/engine/reference/builder/).

```consol
touch Dockerfile
```

The `Dockerfile` is a file that documents the set of _instructions_ that Docker needs to build the final image. Instead of running `go build *.go` on the command line, we will now include that build instruction inside the `Dockerfile`. When we have written all the instructions, we can build the image by running the [`docker build`](https://docs.docker.com/engine/reference/commandline/build/) command. Whereas `go build` outputs a binary that we can run on the command-line, `docker build` outputs a Docker image which we can run using Docker's [`docker run`](https://docs.docker.com/engine/reference/commandline/run/) command.

## Why Do We Need Docker?

We are using Docker here because we want to run our application inside Kubernetes. If all you need is a binary to run on a particular server, or if you prefer to use Configuration-as-Code (CasC) tools like Puppet or Ansible to manage what runs on your servers, then you don't need to bother with the `Dockerfile` or Docker image - all you need to run is `go build *.go`.

But defining the build instructions inside a `Dockerfile` have further benefits:

- Building from source may require tools (e.g. Git, OpenSSH, Maven, Bazel, Node.js, `go` etc) to be installed on the machine that is running the build. We often take these tools for granted since they are usually the first tools we install, but they may not be available on a fresh machine and the build will fail. However, if you standardize the building and running of your applications using Docker, then the _only_ dependency any developer in your organization needs to build and run your application is Docker. If an application needs a certain tool for the build, that tool is documented inside a `Dockerfile` and will be installed within the Docker image.
- As your application becomes more complex (e.g. when you need to use private Go modules) then the list of instructions to build the binary is going to increase. Instead of listing out a changing set of instructions for developers to run, if we document these instructions inside a `Dockerfile`, then developers would only ever need to run `docker build .`
- It allows the build step to be automated. If we document the build instructions in arbitrarily-structured text, machines can't understand that. For automation to be possible, we need to document the instructions in a script or a format like the `Dockerfile`.

## Adding Instructions to Our `Dockerfile`

The first instruction in any `Dockerfile` is [`FROM`](https://docs.docker.com/engine/reference/builder/#from), which sets the _base image_ for that image. To understand what this means, you must first understand what a Docker image actually is.

### Understanding Base Images

A Docker image is actually made up of one or more layers of filesystem changes, where the changes on a higher layer are made on top of changes made in the lower layers. When we specify instructions inside a `Dockerfile`, such as to copy some files (e.g. `COPY . .`) or run some command (e.g. `RUN go build .`), each of those instructions will turn into a layer in the image. So the base image defines the stack of layers on which the next layer in the image we are building is going to be built on top of.

Since we are building a Go application 'inside' the image, we need to make sure the `go` tool and the Go standard library are set up before we run any `go` commands. We could use a generic base image like [`ubuntu`](https://hub.docker.com/_/ubuntu) and install the tools ourselves. But a more convenient method is to use an image where everything is set up already (e.g. [`golang`](https://hub.docker.com/_/golang)). For this post, we are going to do the latter and use the `golang:1.15.7-alpine3.13` image.

So open the `Dockerfile` and add the following line:

```dockerfile
FROM golang:1.15.7-alpine3.13
```

### Copying Go Source Code

Next, we need to copy the source code into the image. We can do this using the [`COPY`](https://docs.docker.com/engine/reference/builder/#copy) instruction, which have the signature `COPY [--chown=<user>:<group>] <src>... <dest>`.

`<src> ...` is one or more paths on the local machine from which to copy from, and `<dest>` is the location within the image to copy to. For both, you can use an absolute or relative path. For relative paths, the `<src>` paths are interpreted as relative to the build context (which you pass in when you run `docker build`), whereas the `<dest>` path is interpreted as relative to the _working directory_ within the image.

#### Understanding the Working Directory

The working directory of an image is the directory from which commands are ran and can be set inside a `Dockerfile` (or in the `Dockerfile` of its base image) using the [`WORKDIR`](https://docs.docker.com/engine/reference/builder/#workdir) instruction. To see what `WORKDIR` value our base image has set, we can look inside its `Dockerfile`, which can be found [on GitHub](https://github.com/docker-library/golang/blob/45f79a2f9262a34b31ab4de0ac7e0728e4002a6b/1.15/alpine3.13/Dockerfile). At the bottom of the file, you'll find:

```dockerfile
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
```

This means the working directory of the base image is `/go`, which is also set as the `$GOPATH` environment variable.

Since we are using Go modules, `$GOPATH` has significance. Before Go modules became the norm, a Go developer must put the source code of all his Go packages inside `$GOPATH/src`. When you need to download the source code for third-party packages, you would similarly place them at `$GOPATH/src`. With Go modules, you can put the source code in any directory and mark it as a module by adding a `go.mod`. But to not break backwards compatibility, any source code directories inside the `$GOPATH` will be 

Because of this, we cannot simply copy our source code into the working directory of the image, because the `go` tool would not recognize it as a valid Go module. If you had tried to do this, the `go` tool will complain with the following error:

```console
$GOPATH/go.mod exists but should not
```

So we should copy our source code into a different directory, but which directory? Typically, the user's home directory is a good choice. To decipher which user we are using to run our commands, we can look for a [`USER`](https://docs.docker.com/engine/reference/builder/#user) instruction inside our `Dockerfile` (or the `Dockerfile` of our base image). If we look at the [`Dockerfile` for `golang:1.15.7-alpine3.13`](https://github.com/docker-library/golang/blob/45f79a2f9262a34b31ab4de0ac7e0728e4002a6b/1.15/alpine3.13/Dockerfile), there's no `USER` instruction. But from the `FROM alpine:3.13` instruction, we know that the `golang:1.15.7-alpine3.13` image is based off the [`alpine:3.13`](https://hub.docker.com/_/alpine/) image, so let's take a look inside its [source code](https://github.com/alpinelinux/docker-alpine/blob/3ba85e92c35d4a488bc5a565a4e1c5338bc89775/x86_64/Dockerfile).

```dockerfile
FROM scratch
ADD alpine-minirootfs-3.13.1-x86_64.tar.gz /
CMD ["/bin/sh"]
```

So it seems like the user is not set anywhere in any of the `Dockerfile`. In this case, Docker defaults to running commands as the `root` user and group. This is generally considered insecure, because Docker does not, by default, isolate container processes with user namespaces, which means the `root` inside your container is the `root` on the host machine. Instead, we should run containers as a non-privileged user. Let's create a new user with a home directory, which will make our image more secure and also provide a place to copy in our source code.

Below the `FROM` instruction, add a `RUN` instruction that will run the `adduser` command to add a new user.

```dockerfile
RUN adduser --disabled-password echo
```

By default, `adduser` will interactively prompt for a password; but since we don't need a password (since nobody would be signing into the container), we are using the `--disable-password` option to disable the interactivity.

By default, the `adduser` command will create a home directory for the new user at `/home/<username>`. So we can use that as our working directory by adding our own `WORKDIR` instruction to our `Dockerfile`.

```dockerfile
WORKDIR /home/echo/
```

Lastly, don't forget to set the user using the `USER` instruction.

```dockerfile
USER echo
```

We are now ready to copy in the source code. Below the `USER` instruction, add a new `COPY` instruction that copies everything in the project directory into the image's working directory.

```dockerfile
COPY . .
```

### Building the Binary

We have the source code inside our image, now we need to run the `go build` command. We can run arbitrary commands inside a `Dockerfile` by using the [`RUN`](https://docs.docker.com/engine/reference/builder/#run) instruction. Below the `COPY` instruction, add a `RUN` instruction.

```dockerfile
RUN go build .
```

Since our source code file is named `main.go`, the name of the binary should be `main`. But we can be more explicit and specify the output file location and name using the `-o` (output) option. Update the `RUN` instruction with this option; while we are at it, let's give it a more intuitive name.

```dockerfile
RUN go build -o echo .
```

### Specifying an Entrypoint

Up to this point, we have specified how to build a binary of our program, but we haven't specified how Docker should start our program when we run the image using [`docker run`](https://docs.docker.com/engine/reference/run/). We can do this using the [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#entrypoint) instruction, which specifies the executable or command to run when the container starts.

Since the binary we built can be found at `./echo`, we can use that as the entry point. At the bottom of the `Dockerfile`, add this line:

```dockerfile
ENTRYPOINT ./echo
```

### Building the Docker Image

We are now ready to build our image. Making sure you are inside the `acme/echo/` directory, run:

```console
$ docker build -t echo .
```

Every image has an alphanumerical ID (e.g. `sha256:79976bf52154347db5f88c81d0bfdca0d2ef6637e38987c07a3dbd4f6b5846a3`, often shorted to `79976bf52154`); the `-t echo` option tags the resultant image with a more human-friendly name for easier reference. The period (`.`) at the end specifies the build context, which is the directory that contains the `Dockerfile`, the source code, and also the base path for resolving any relative paths referenced in the `Dockerfile`.

The first time you run this, Docker will look for the `golang:1.15.7-alpine3.13` image locally on your machine, and when it cannot find it, download it from Docker Hub.

```console
Sending build context to Docker daemon  10.06MB
Step 1/4 : FROM golang:1.15.7-alpine3.13
1.15.7-alpine3.13: Pulling from library/golang
4c0d98bf9879: Pull complete 
9e181322f1e7: Pull complete 
6422294da7d3: Pull complete 
8b36f00a8e74: Downloading [=================================>                 ]  72.42MB/106.8MB
5e5ebcc3e852: Download complete 
```

After a short while, the image would be built.

```console
...
Removing intermediate container 94bd6a03bb27
 ---> 9244edc3672a
Step 6/6 : ENTRYPOINT ./echo
 ---> Running in aafa2294efb4
Removing intermediate container aafa2294efb4
 ---> 79976bf52154
Successfully built 79976bf52154
Successfully tagged echo:latest
```

And you can find the image when you run `docker images`.

```console
$ docker images echo
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
echo         latest    79976bf52154   6 minutes ago   392MB
```

## Testing the Echo Server Image

We can now test out our Docker image by running using `docker run`.

```console
$ docker run --rm --name echo --publish 8080:8080 echo
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

By default, when we stop a container, the container will still remain, but in a stopped state instead of a running state. This allows you to start the container again later. Here, however, we used the `--rm` option in `docker run` to instruct Docker to automatically remove the container as soon as it has stopped. This is a useful option to specify, especially during testing and development, otherwise you may end up with hundreds of stopped containers which you'd have to clear manually.

The `--name echo` option to `docker run` gives the container a human-friendly name. Without a name, you'd have to refer to the container by its ID (e.g. `6d37ac75ef8b`). The `--publish 8080:8080` option forwards `0.0.0.0:8080` to the container's port `8080`. This allows us to connect with our echo server over an address like `localhost:8080`. The last `echo` is the name of the tag of the image we want to run.

Our `echo` image is now running as a container and you can see the log output from the `gin` package.

On a separate terminal, you can run `docker ps` to see a list of running containers

```console
$ docker ps 
CONTAINER ID  IMAGE  COMMAND              STATUS  PORTS                   NAMES
6d37ac75ef8b  echo   "/bin/sh -c ./echo"  Up      0.0.0.0:8080->8080/tcp  echo
```

We can also send a request to the echo server to confirm the response.

```console
$ curl -i http://localhost:8080/\?q=hello
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Wed, 03 Feb 2021 20:57:42 GMT
Content-Length: 5

hello
```

You'll also see the request logged in the `stdout` of the server.

```console
[GIN] 2021/02/01 - 01:02:09 | 200 |        40.2µs |      172.17.0.1 | GET      "/?q=hello"
```
