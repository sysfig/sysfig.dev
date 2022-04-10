---
title: Creating a Local Kubernetes Cluster with kind
slug: local-kind-kubernetes-cluster
date: 2021-02-14T17:59:09-08:00
chapter: a
order: 3
tags:
    - kubernetes
    - kind
    - docker
    - registry
    - ctlptl
    - kubectl
draft: true
---

We now have an lean image which we can run. However, instead of simply running a standalone instance with `docker run`, we'd like to deploy it in a Kubernetes cluster.

Typically, this is the part where you'd go to AWS and spin up a test or development Kubernetes cluster, but services like AWS EKS can be costly ($73/month just for the master control plane, which does not include worker nodes and storage). There are cheaper alternatives - DigitalOcean's managed Kubernetes offering provides the master control plane for free, meaning you only need to pay for the worker nodes, load balancers, and storage. You can get a minimal cluster going for as little as $10/month.

However, if you want to pay nothing for your development cluster, or you want to work on your application offline, then you should consider running a local Kubernetes cluster on your development machine.

## Survey of the Local Kubernetes Cluster Landscape

There are many tools that allows you to run a local development Kubernetes cluster:

- [minikube](https://minikube.sigs.k8s.io/docs/) was the first tool that allowed you to run Kubernetes locally. As such, there's a mature community around it and support tends to be more available. At its conception, minikube was limited to running one single-node cluster, where your local machine would act as the single node. However, [multi-node clusters](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/) have been supported since v1.10.1. Nowadays, you can use virtual machines (VMs) and containers to act as nodes and deploy a Kubernetes cluster on top of them.

  The double-edge sword with minikube is that it is very configurable, and thus there's a lot of options to configure and documentation to read.

  ![](https://raw.githubusercontent.com/kubernetes/minikube/master/images/logo/logo.png)
  
- [kind](https://kind.sigs.k8s.io/) (short for **k**ubernetes **in** **d**ocker) is a simpler alternative to minikube that runs Kubernetes atop of Docker container nodes. kind is used internally by the Kubernetes team to test Kubernetes.

  Its biggest benefit is that it's fast. You can spin up a new cluster in less than 30 seconds, and remove it in a few seconds.

There are other tools such as [Microk8s](https://microk8s.io/) (created by Canonical, the owner of Ubuntu), [k3d](https://github.com/rancher/k3d) (supported by Rancher, a web-based GUI interface for Kubernetes); but we will not discuss them here.

For our project, we are going to use kind to set up our local development cluster.

## kind

`kind` comes as a command-line tool which you can install by running the following command outside of any Go module directories (a good place would be in your home directory):

```console
$ GO111MODULE="on" go get sigs.k8s.io/kind@v0.10.0
```

This will download the Go source code and place them at `$GOPATH/bin`. If you had run this inside a Go module, it will be treated as a dependency local to that module, instead of one that is local to your user.

Alternatively, you can use package managers. For example, on macOS, you can use Homebrew.

```console
$ brew install kind
```

Then, we can can create a new Kubernetes cluster using `kind` by running `kind create cluster`. This will spin up new Docker containers that acts as Kubernetes nodes. Each container is started using a _[node image](https://kind.sigs.k8s.io/docs/design/node-image)_, which includes programs such as Docker, `systemd`, and Kubernetes components (e.g. `kubeadm`, `kubectl`, `kubelet`) - needed to run nested containers within our node containers. The node image is based on the _[base image](https://kind.sigs.k8s.io/docs/design/base-image)_, which, in turn, is based on the [`ubuntu`](https://hub.docker.com/_/ubuntu) image.

Each node image includes a different version of Kubernetes. So, to use a specific version of Kubernetes in your local cluster, find the relevant tag in the [`kindest/node`](https://hub.docker.com/r/kindest/node/tags) registry (the tag name for the image corresponds to the Kubernetes version).

So let's run the `kind create cluster` command to spin up a new cluster.

```
$ kind create cluster --name acme
Creating cluster "acme" ...
 ✓ Ensuring node image (kindest/node:v1.20.2) 🖼 
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-acme"
You can now use your cluster with:

kubectl cluster-info --context kind-acme
```

By default, the cluster will use the default name of `kind`; we use the `--name` option here to explicitly set a name, making it clearer if we were to run multiple local clusters using `kind`. Setting the `--name acme` will create a cluster named `kind-acme`.

This will create a Kubernetes configuration file at `$KUBECONFIG` (which defaults to `$HOME/.kube/config`).

```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSU...
  name: kind-acme
contexts:
- context:
    cluster: kind-acme
    user: kind-acme
  name: kind-acme
current-context: kind-acme
kind: Config
preferences: {}
users:
- name: kind-acme
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FUR...
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJ...
```

This configuration file defines all the Kubernetes cluster, both local and remote, that your user has access to. You can also get a list of all local `kind` clusters by running `kind get clusters`.

```
$ kind get clusters
acme
```

You can also see the node containers that `kind` have spun up.

```
% docker ps -a
CONTAINER ID   IMAGE                  COMMAND                  PORTS                       NAMES
e66145d0c32e   kindest/node:v1.20.2   "/usr/local/bin/entr…"   127.0.0.1:50386->6443/tcp   acme-control-plane
```

By default, `kind` spins up a single-node cluster. Later on, we will show you how to create multi-node clusters by passing in a configuration file to `kind create cluster`. Using configuration file(s) is preferred to passing in options on the command-line as you can commit the file to source control.

Now we have a local Kubernetes cluster, we can interact with it using [`kubectl`](https://kubernetes.io/docs/reference/kubectl/overview/). `kubectl` is the command-line tool that allows you to manage your Kubernetes cluster(s) and the workloads that run on them. For example, you can use `kubectl` to create a new [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/). (N.B. `kubectl` works with an existing Kubernetes cluster, it cannot create a new cluster. For that, you'll have to use tools like [`kubeadm`](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/))

Before continuing, install `kubectl` following the linked [instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl/). For macOS, it's recommended to install `kubectl` with Homebrew:

```
$ brew install kubectl
```

To test if you have installed `kubectl` successfully, run `kubectl version`. You should see something like this:

```
$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.2", GitCommit:"faecb196815e248d3ecfb03c680a4507229c2a56", GitTreeState:"clean", BuildDate:"2021-01-14T05:15:04Z", GoVersion:"go1.15.6", Compiler:"gc", Platform:"darwin/amd64"}
```

We can now use `kubectl` to interact with our local `kind-acme` cluster. We will run the `kubectl cluster-info` command to get some information about our cluster.

```console
$ kubectl cluster-info --context kind-acme
Kubernetes control plane is running at https://127.0.0.1:50386
KubeDNS is running at https://127.0.0.1:50386/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

We use the `--context` option to specify the name of the cluster we want to interact with. This allows us to use `kubectl` to manage multiple clusters.

If the `--context` option is not specified, the default context (the one specified with the `current-context` key in `$KUBECONFIG`) is used. Since we have only a single cluster right now, and that is already set to the default context, we can simply omit the `--context` option.

But if you have multiple clusters configured in your `$KUBECONFIG`, you can change the context for any subsequent commands by running `kubectl config use-context`. This will save you from setting `--context` for every command.

```console
$ kubectl config use-context kind-acme
Switched to context "kind-acme".
```

In the output of `kubectl cluster-info`, the `Kubernetes control plane is running at` line tells you that our Kubernetes cluster is running and we can authenticate (as an administrative user) with it.

We can also see all the resources running on the cluster by running `kubectl get all`

```console
$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   88m
```

## Writing Kubernetes Manifests

Now that we have a working Kubernetes cluster, it's time to deploy our echo server on it. We already have a Docker image, but Kubernetes doesn't deploy container images; instead, it deploys _workload resources_ such as [Pods](https://kubernetes.io/docs/concepts/workloads/pods/), [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/), [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/), [StatefulSets](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/), etc.

The typical way to deploy a service (lowercase 's', to distinguish it from a Kubernetes [Service](https://kubernetes.io/docs/concepts/services-networking/service/)) is to declare a Deployment resource in Kubernetes.

So let's write a Kubernetes manifest for a Deployment for the echo service.

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echo
spec:
  selector:
    matchLabels:
      app: echo
  replicas: 1
  template:
    metadata:
      labels:
        app: echo
    spec:
      containers:
        - name: echo
          image: echo
          ports:
            - name: http
              containerPort: 8080
```

Save this in a file named `echo_deployment.yaml` in the `acme/` directory. Then, apply the manifest into our cluster by running 

```console
$ kubectl apply -f echo_deployment.yaml
deployment.apps/echo created
```

We can run `kubectl get all` to see the resources that are being created from our `kubectl apply` call.

```console
$ kubectl get all
NAME                        READY   STATUS             RESTARTS   AGE
pod/echo-86bb48bc78-8gxxc   0/1     ImagePullBackOff   0          14s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   2d23h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   0/1     1            0           14s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-86bb48bc78   1         1         0       14s
```

As expected, a new `deployment.apps/echo` Deployment resource is created that controls the `replicaset.apps/echo-86bb48bc78` ReplicaSet, which, in turn, controls a single `pod/echo-86bb48bc78-8gxxc` Pod resource.

However, if we run `kubectl get all` again after a few more moments, it seems as though the Pod is failing to transition into a ready state. We can see from the output that its status is now `ErrImagePull`.

```console
% kubectl get all
NAME                        READY   STATUS         RESTARTS   AGE
pod/echo-86bb48bc78-8gxxc   0/1     ErrImagePull   0          114s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   2d23h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   0/1     1            0           114s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-86bb48bc78   1         1         0       114s
```

This `ErrImagePull` status means Kubernetes cannot find the image specified in the manifest (`echo` in our case). This is because although the image is available on our local machine, it is not available inside the Kubernetes nodes. When a Pod needs to run on a Kubernetes node but it does not have the image locally, Docker will try to resolve the name from all known registries and download it from a registry. Since we have not configured containerd to any registries, it will default to `docker.io` (i.e. Docker Hub). But since that image does not exist on Docker Hub, Kubernetes is giving us the feedback that an image cannot be resolved and thus the Pod cannot be started.

If we had used an image that _is_ on Docker Hub (e.g. [`ghost`](https://hub.docker.com/_/ghost)), then our `kubectl apply` would have worked. Just to demonstrate, update the `spec` portion of your `echo_deployment.yaml` file to:

```yaml
spec:
  containers:
    - name: echo
      image: ghost
      ports:
        - name: http
          containerPort: 8080
```

And run `kubectl apply` again.

```console
$ kubectl apply -f echo_deployment.yaml
deployment.apps/echo configured
```

Give Docker a few seconds to download the `ghost` image on the Kubernetes node (_not_ your local machine) and for Kubernetes to start the Pod, ReplicaSet, and Deployment. The end result should be that all three resources are in the ready state.

```console
% kubectl get all
NAME                        READY   STATUS    RESTARTS   AGE
pod/echo-667dbd5c58-2sps9   1/1     Running   0          4m1s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   2d23h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   1/1     1            1           10m

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-667dbd5c58   1         1         1       4m1s
replicaset.apps/echo-86bb48bc78   0         0         0       10m
```

This demonstrates that our Kubernetes manifest is valid, we just need to configure [containerd](https://containerd.io/) running inside our Kubernetes node(s) to know where to download our `echo` image from. There are several ways to do this, listed from worst to best:

- Build our own custom node image by updating the `Dockerfile` to include our `echo` image - this is not sustainable especially if we want to run many services and new versions are released frequently
- Upload our `echo` image to Docker Hub with a name like `sysfig/echo`, and update our manifest to use this name - this is fine as long as you don't mind having your images public, or paying $5/month for Docker Hub's Pro plan.
- Load images into the node with `kind load` - this works but you'd have to manually load images one by one into each cluster. If you recreate your cluster (common for testing/development), you'd have to repeat the same process again.
- [Run a local Docker registry](https://docs.docker.com/registry/deploying/), deploy a `kind` cluster with a special [containerd configuration file](https://github.com/containerd/containerd/blob/master/docs/cri/registry.md) that registers this local registry, and then link the registry container with the `kind` cluster's network. The `kind` documentation provides a [shell script](https://kind.sigs.k8s.io/docs/user/local-registry/) that performs these tasks.

  In more details, the steps involve:
  1. Running the [`registry`](https://hub.docker.com/_/registry) container on your local environment
  2. Create a new `kind` cluster using a special containerd configuration that adds the local registry to the list of registries to try
  3. Running [`docker network connect`](https://docs.docker.com/engine/reference/commandline/network_connect/) to connect the `registry` container to the `kind` cluster's network
  4. Create a ConfigMap resource within the cluster to [document the local registry](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry)

- Use a command-line tool called [`ctlptl`](https://github.com/tilt-dev/ctlptl) (pronounced "cattle paddle") to abstract everything mentioned above into a single command.

`ctlptl` is written in Go and built into a binary. You can download the binary from the [Releases](https://github.com/tilt-dev/ctlptl/releases) page and move the binary into your `$PATH`.

```
$ curl -L https://github.com/tilt-dev/ctlptl/releases/download/v0.4.2/ctlptl.0.4.2.linux.x86_64.tar.gz | tar -xz
$ sudo mkdir /opt/bin/
$ sudo mv ctlptl /opt/bin/
```

Make sure the `/opt/bin/` directory is in your `$PATH`. You can do this by adding the following line to your profile (e.g. `$HOME/.profile`)

```sh
export PATH=$PATH:/opt/bin
```

The next time your profile is loaded (either by logging out and in again, or by sourcing the file by running `. ~/.profile`), the `ctlptl` command should be available on your shell.

For convenience, on macOS, you can also install `ctlptl` using [Homebrew](https://brew.sh/):

```console
$ brew install tilt-dev/tap/ctlptl
```

On Windows, you can use [Scoop](https://scoop.sh/):

```console
$ scoop bucket add tilt-dev https://github.com/tilt-dev/scoop-bucket
$ scoop install ctlptl
```

Whichever method you choose to install `ctlptl`, you can confirm it is installed properly by running `ctlptl version`.

```console
$ ctlptl version
v0.4.2, built 2021-01-25
```

With `ctlptl` installed, we can now use it to re-create our existing cluster that makes use of a local registry.

## Creating a New Cluster with Local Registry with `ctlptl`

First, create a `ctlptl` cluster configuration file that defines what cluster we are making.

```yaml
apiVersion: ctlptl.dev/v1alpha1
kind: Cluster
product: kind
registry: registry
kindV1Alpha4Cluster:
  name: acme
```

`ctlptl` supports creating local Kubernetes clusters using different tools. It supports kind, which we are using, but also Minikube, Docker for Mac, and Docker for Windows. So the `product` key here specifies which tool we are using.

The `registry` key declares that we want to create a local registry that is connected to this cluster. The value of the key is the name of the Docker container that will run this registry; here, we are going to name it `registry`. The `kindV1Alpha4Cluster` key takes a kind [configuration file](https://kind.sigs.k8s.io/docs/user/configuration/) where certain keys (`kind: Cluster` and `apiVersion: kind.x-k8s.io/v1alpha4`) are implied and are omitted. Here, we are using the `kindV1Alpha4Cluster.name` key to specify the name of the cluster.

Just like `kubectl`, we can apply a `ctlptl` cluster configuration by running `ctlptl apply`.

```console
$ ctlptl apply -f cluster.yaml
Deleting cluster kind-acme to initialize with registry registry
Deleting cluster "acme" ...
Creating registry "registry"...
Creating cluster "acme" ...
 ✓ Ensuring node image (kindest/node:v1.20.2) 🖼 
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-acme"
You can now use your cluster with:

kubectl cluster-info --context kind-acme

Have a nice day! 👋
   Connecting kind to registry registry
Switched to context "kind-acme".
 🔌 Connected cluster kind-acme to registry registry at localhost:54934
 👐 Push images to the cluster like 'docker push localhost:54934/alpine'
cluster.ctlptl.dev/kind-acme created
```

This will delete our existing cluster, create a Docker container running the registry, create a new cluster named `acme`, and connecting the registry with the cluster.

There's actually a quicker way to create a cluster with `ctlptl`, and that's to run `ctlptl create cluster kind --name=kind-acme --registry=registry`. But the benefit of defining our cluster in a configuration file and using `ctlptl apply` is that no matter how much changes you make to the cluster configuration, the end-user would only ever have to run the same command (i.e. `ctlptl apply -f cluster.yaml`). The `ctlptl apply` command is also idempotent, which means you can apply them many times and get the same result. For example, if we already have the `acme` cluster running, running `ctlptl apply` again would simply do nothing.

```console
$ ctlptl apply -f cluster.yaml
Switched to context "kind-acme".
cluster.ctlptl.dev/kind-acme created
$ echo $?
0
```

If we try to run `ctlptl create cluster` on a cluster that already exists, the command returns with an error exit code.

```console
$ ctlptl create cluster kind --name=kind-acme --registry=registry
Cannot create cluster: already exists
$ echo $?
1
```

Now `ctlptl` has finished setting up our registry and cluster, we can run `docker ps` to see what containers are running.

```
% docker ps -a
CONTAINER ID   IMAGE                  COMMAND                 PORTS                                        NAMES
684ab12a9cbd   kindest/node:v1.20.2   "/usr/local/bin/entr…"  127.0.0.1:54987->6443/tcp                    acme-control-plane
7194df851661   registry:2             "/entrypoint.sh /etc…"  0.0.0.0:54934->5000/tcp, :::54934->5000/tcp  registry
```

You can see a new container named `registry` is running the image `registry:2` and our host's port `54934` is being forwarded to the container's port `5000`. This means we can now push images to our local registry using `localhost:54934` as the registry address (i.e. `docker push localhost:54934/echo`). Let's try that.

```console
$ docker images
REPOSITORY  TAG     IMAGE ID      SIZE
registry    2       5c4008a25e05  26.2MB
echo        latest  785abbf65c45  10.1MB
```

First, we run `docker images` to confirm that the `echo` image is still available on our local machine. We then need to tag the image again, prefixing the name with the registry's address.

```console
$ docker tag 785abbf65c45 localhost:54934/echo:$(date '+%s')
$ docker tag 785abbf65c45 localhost:54934/echo:latest
```

We are also giving the image a more informative tag using the UNIX timestamp (alongside the `latest` tag). After we have tagged it, we can push it to our local registry.

```console
$ docker push localhost:54934/echo:1614881428
The push refers to repository [localhost:54934/echo]
b17eefab8db4: Pushed 
2da09b410833: Pushed 
2bcf9052bd65: Pushed 
1614881428: digest: sha256:b22fddbe74af74e36f8d17b4a8772baee64353d8b68f11548bf8d2b1a425dc66 size: 942
$ docker push localhost:54934/echo:latest
The push refers to repository [localhost:54934/echo]
b17eefab8db4: Layer already exists 
2da09b410833: Layer already exists 
2bcf9052bd65: Layer already exists 
latest: digest: sha256:b22fddbe74af74e36f8d17b4a8772baee64353d8b68f11548bf8d2b1a425dc66 size: 942
```

We can confirm that the `containerd` running inside our node containers are indeed configured to check our local repository by getting into the node container and checking the `/etc/containerd/config.toml`.

```
$ docker exec -it 684ab12a9cbd /bin/bash
root@acme-control-plane:/# cat /etc/containerd/config.toml 
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "k8s.gcr.io/pause:3.3"
    snapshotter = "overlayfs"
    tolerate_missing_hugepages_controller = true
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.test-handler]
          runtime_type = "io.containerd.runc.v2"
    [plugins."io.containerd.grpc.v1.cri".registry]
      [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
        [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:54934"]
          endpoint = ["http://registry:5000"]
        [plugins."io.containerd.grpc.v1.cri".registry.mirrors."registry:5000"]
          endpoint = ["http://registry:5000"]
```

Now, update the Kubernetes deployment manifest with the new name for the image (`localhost:54934/echo`).

```yaml
spec:
  containers:
    - name: echo
      image: localhost:54934/echo
      ports:
        - name: http
          containerPort: 8080
```

And try to apply the manifest again, and we should not get the same `ImagePullBackOff` error.

```console
$ kubectl apply -f echo_deployment.yaml
deployment.apps/echo created
```

Now let's run `kubectl get all` to get the status of all resources deployed in our cluster.

```console
% kubectl get all
NAME                        READY   STATUS              RESTARTS   AGE
pod/echo-667dbd5c58-vjkps   0/1     ContainerCreating   0          15s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   18h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   0/1     1            0           15s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-667dbd5c58   1         1         0       15s
```

Seeing the Pod going into the `ContainerCreating` status is a good sign, because it means the image was successfully found and downloaded. This must occur before you start the container.

After a few more seconds, we can run` kubectl get all` again and see that our echo server is now running in the cluster without issues.

```console
% kubectl get all
NAME                        READY   STATUS    RESTARTS   AGE
pod/echo-667dbd5c58-vjkps   1/1     Running   0          32s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   18h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   1/1     1            1           32s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-667dbd5c58   1         1         1       32s
```

Fantastic! We have now successfully created a local Kubernetes cluster and deployed our first workload on it.

Now, you might be thinking - Great! `kubectl` is telling me that my echo server is running, but how do it reach it? How do I know it's _really_ running?

As we have it now, only workloads within the cluster can communicate with each other. So right now, there's actually no way to hit the echo server. But fret not, in the next post, we will cover some basics of Kubernetes networking, and expose our echo server using a _[Service](https://kubernetes.io/docs/concepts/services-networking/service/)_ resource.
