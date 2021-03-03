---
title: Optimizing Dockerfile
slug: dockerfile-optimization
date: 2021-02-03T13:10:02-08:00
chapter: a
order: 2
tags:
    - docker
    - dockerfile
    - golang
draft: true
---

It's good that our `echo` image runs just like the Go binary we built in the previous post, but if we check the file size of our image and compare it with that of the binary, we'll find that the Docker image is much larger.

```console
$ ls -ahl main 
-rwxr-xr-x  1 d4nyll  staff   9.6M Feb  3 13:01 main

$ docker images
REPOSITORY       TAG                 IMAGE ID       SIZE
echo             latest              79976bf52154   392MB
golang           1.15.7-alpine3.13   54d042506068   299MB
```

Whilst the `main` binary is 9.6 megabytes, the `echo` image is 392 megabytes, more than 40 times larger. Whilst a little bit of overhead is expected, the final image size should be in the same order of magnitude to the binary it is essentially wrapping.

This size increase can be understood when we look at the size of the `golang:1.15.7-alpine3.13` image, which is 299 megabytes. It is large because of all the tools and libraries needed to build the binary are included with the image.

But the build tools and any unused libraries are not used after the binary is built, and won't be used when the image is ran, so it would be ideal to remove them from the image to reduce the image's file size.

But we cannot simply add a `RUN rm -rf /build/tools/dir/` instruction at the end of our image. As we explained earlier, a Docker image is made up of layers of filesystem changes. Each layer builds on top of the previous layers. If we were to delete some files from the image, this will actually _add_ an additional layer on top of the existing layers. Whilst this will only marginally increase the file size of the image, it will most definitely not reduce it.

One solution is to run the `rm` command in the same instruction as the command that added the files. This can be done using shell built-ins like `&&` or `;`. But this is cumbersome as we'd have to install the Go tools, build our binary, and then remove the Go tools, all in the same instruction.

Instead, we can turn to a feature of Docker called _[multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/)_.

## Using Multi-Stage Builds

Multi-stage builds allows you to define multiple images within the same `Dockerfile` and allows you to copy files from one image to another. In the context of a multi-stage build, each image defined in the `Dockerfile` is called a _stage_. We can use multi-stage builds to reduce the size of our final image by building our binary on top of the larger `golang:1.15.7-alpine3.13` image that has all the build tools, but then copy only the binary onto a small base image to use as our final image.

Replace our `Dockerfile` with the following instructions which makes use of multi-stage builds.

```dockerfile
FROM golang:1.15.7-alpine3.13
RUN adduser --disabled-password echo
WORKDIR /home/echo/
USER echo
COPY . .
RUN go build -o echo .

FROM scratch
COPY --from=0 /home/echo/echo /
ENTRYPOINT /echo
```

To define multiple stages, we specify multiple `FROM` instructions. Each stage can have a different base image (e.g. `golang` or `scratch`).

Here, we added a new stage that uses [`scratch`](https://hub.docker.com/_/scratch/) as the base image. Technically, `scratch` is not an image _per se_, but rather an empty placeholder for an empty base image, used when you need to construct images from scratch. It does not add a layer to the image.

We can then copy the `echo` binary located at `/home/echo/echo` from the first stage to the second stage using `COPY --from=0` (`0` refers to the first stage since it's zero-indexed). We've also moved the `ENTRYPOINT` instruction from the first stage to the second stage, since that's the final image we want to keep.

We can build the two images specified in this `Dockerfile` using the same command as before.

```console
$ docker build -t echo .
Sending build context to Docker daemon  10.06MB
Step 1/9 : FROM golang:1.15.7-alpine3.13 as build
 ---> 54d042506068
...
Removing intermediate container d22df3e0faf4
 ---> 2878763d7c9d
Step 7/9 : FROM scratch
 ---> 
...
Removing intermediate container 87dd370daff8
 ---> f5073b67aa37
Successfully built f5073b67aa37
Successfully tagged echo:latest
```

After the images are built successfully, we can run `docker images` to list them out.

```console
$ docker images
REPOSITORY  TAG     IMAGE ID      SIZE
echo        latest  f5073b67aa37  10.1MB
<none>      <none>  2878763d7c9d  392MB
<none>      <none>  79976bf52154  392MB
```

The first thing to point out is that the new `echo:latest` image (with ID `f5073b67aa37`) has drastically decreased in size to 10.1 megabytes, which is just half a megabyte of overhead over the raw binary. You may also notice that the previous `echo` image (with ID `79976bf52154`) has now been untagged. This is because the new `echo:latest` image (with ID `f5073b67aa37`) has taken over the use of the tag.

There's another untagged image (with ID `2878763d7c9d`) which corresponds to the first stage. When we used `docker build`'s `-t` flag to tag an image, the image that corresponds with the last stage in the `Dockerfile` will be the one that is tagged, the other images will not have a tag and are called _intermediate images_. So the image with ID `2878763d7c9d` is an intermediate image.

## Using 

But when we try to run our new image, we get the following error:

```
$ docker run --rm --name echo --publish 8080:8080 echo
docker: Error response from daemon: OCI runtime create failed: container_linux.go:370: starting container process caused: exec: "/bin/sh": stat /bin/sh: no such file or directory: unknown.
```

To understand this error, we must first understand that there are two syntax for instructing Docker to execute a command: the _shell form_ and the _exec form_. The shell form, which is the form we have been using so far, executes each `RUN`, `ENTRYPOINT`, `CMD` command by first invoking a shell and then running the command within the shell.

This error arises because we are using the shell form but the default shell (`/bin/sh -c` on Linux) is not present in our `scratch` image. To avoid this error, we should use the exec form, which will execute the command directly, without first invoking a shell.

To use the exec form, convert the command you want to run into an array of strings. For instance, we can write:

```Dockerfile
...
FROM scratch
COPY --from=0 /home/echo/echo /
ENTRYPOINT ["/echo"]
```

##

This time, 

```
$ docker run --rm --name echo --publish 8080:8080 echo
standard_init_linux.go:219: exec user process caused: no such file or directory
```

This is because the `echo` binary we built in the first step uses C libraries provided by Alpine Linux (Alpine uses the musl libc library) and is _dynamically linked_ to some of the libraries. This means the C libraries it uses are not statically compiled into the binary; rather, the binary expects to find the same libraries on the host when it is run.

This is made possible by _cgo_, which is a tool that allows Go programs to interoperate with C libraries. You can use cgo, for example, to wrap some C code instead of reimplementing it in Go.

When we copied the binary into the `scratch` 'image', our new base image does not have those libraries and this is why the error is complaining that it cannot find those libraries.

Instead, we can disable cgo when building our binary, which ensures that the final binary is not linked to any libraries, and everything it needs to run is included into the binary itself. We can disable cgo using the environment variable `CGO_ENABLED`.

```Dockerfile
FROM golang:1.15.7-alpine3.13
RUN adduser --disabled-password echo
WORKDIR /home/echo/
USER echo
COPY . .
ENV CGO_ENABLED=0
RUN go build -o echo .

FROM scratch
COPY --from=0 /home/echo/echo /
ENTRYPOINT ["/echo"]
```

```console
$ docker build -t echo .
$ docker run --rm --name echo --publish 8080:8080 echo                                    
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

## Naming Stages

We have shown you a basic use of multi-stage builds that uses only two stages. For more complicated build processes, a `Dockerfile` may contain half a dozen stages. Referring to each stage by its index is error-prone and it's not immediately obvious what stage the `--from` option is referring to.

Instead, we can give the first stage a name (e.g. `build`) using the `AS` directive. We can then refer to the stage later on using the `--from` option.

```Dockerfile
FROM golang:1.15.7-alpine3.13 AS build
RUN adduser --disabled-password echo
WORKDIR /home/echo/
USER echo
COPY . .
ENV CGO_ENABLED=0
RUN go build -o echo .

FROM scratch
COPY --from=build /home/echo/echo /
ENTRYPOINT ["/echo"]
```

## Securing Our Scratch Image

In the previous post, we mentioned how it's bad security practice to run a container process as the `root` user, even if it's inside the container. This is why we added the `RUN adduser`, `WORKDIR` and `USER echo` instructions. But these instructions are now in the intermediate image, not in our final image.

We may be tempted to simply copy and paste the three lines under the `FROM scratch` instruction, but that won't work. This is because the `adduser` program is provided by Alpine Linux. Since our final image is based on the `scratch` 'image', the `adduser` program is not available inside that image.

So what can we do to add a non-root user to the `scratch` image? Well, the `adduser` command actually only does 2 things - create a new entry in the `/etc/passwd`, `/etc/group` and `/etc/shadow` files, and create the user's home directory (at `/home/<user_name>`). So we don't need the `adduser` program, we can just do it ourselves.

The `/etc/passwd` file lists out all the named users on the system. Seven pieces of information, on the same line and delimited by colons (`:`), are defined for each user.

```
username:password:user_id:group_id:comment:home_directory:default_shell
```

For instance, you may find this entry in the `/etc/passwd` file of your development machine

```
root:*:0:0:System Administrator:/var/root:/bin/sh
```

The `*` in the password field denotes that direct log in is disabled for this user.

```
echo:*:12345:12345:Echo Server:/:
```


```Dockerfile
FROM scratch
COPY --chown=12345:12345 ./passwd /etc/passwd
USER 12345
COPY --from=build /home/echo/echo /
ENTRYPOINT /echo
```


```Dockerfile
FROM scratch
COPY ./passwd /etc/passwd
COPY ./group /etc/group
USER echo:echo
COPY --from=build /home/echo/echo /
ENTRYPOINT ["/echo"]
```
