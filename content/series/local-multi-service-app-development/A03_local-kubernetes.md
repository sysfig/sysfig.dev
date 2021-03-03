---
title: Creating a Local Kubernetes Cluster with kind
slug: local-kind-kubernetes-cluster
date: 2021-02-014T17:59:09-08:00
chapter: a
order: 3
tags:
    - kubernetes
    - kind
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
          image: "echo"
          ports:
            - name: http
              containerPort: 8080
```

`ctlptl`
