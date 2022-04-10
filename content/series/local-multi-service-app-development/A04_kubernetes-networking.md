---
title: Exposing Kubernetes using Ingress and Services
slug: kubernetes-networking-ingress-services
date: 2021-03-04T10:21:48-08:00
chapter: a
order: 4
tags:
    - kubernetes
    - networking
draft: true
---

Without any extra configuration, all containers within the same Pod can communicate with each other via a loopback interface. This is similar to how all the services (e.g. database, web server, API server) you run on your local machine can communicate with each other via `localhost`/`127.0.0.1`.

Each Pod is also assigned its own IP address. If we use `kubectl describe` to describe a Pod, we can see that it has at least 1 IP address associated with it.

```console
% kubectl get all
NAME                        READY   STATUS    RESTARTS   AGE
pod/echo-667dbd5c58-vjkps   1/1     Running   0          22h

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   40h

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/echo   1/1     1            1           22h

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/echo-667dbd5c58   1         1         1       22h
d4nyll@tuna sysfig.dev % kubectl describe pod/echo-667dbd5c58-vjkps
Name:         echo-667dbd5c58-vjkps
Namespace:    default
Priority:     0
Node:         acme-control-plane/172.18.0.2
Start Time:   Thu, 04 Mar 2021 10:14:03 -0800
Labels:       app=echo
              pod-template-hash=667dbd5c58
Annotations:  <none>
Status:       Running
IP:           10.244.0.5
IPs:
  IP:           10.244.0.5
Controlled By:  ReplicaSet/echo-667dbd5c58
...
```

Therefore, containers from different Pods can communicate with each other through this IP address. But that's not practical because each container would need to include logic in their code to find out the IP address of other applications that it needs to communicate with. Kubernetes also creates DNS records for each Pod (using the structure `<pod-ip>.<namespace>.pod.<cluster-domain>`, or `<pod-ip>.<deployment>.<namespace>.svc.<cluster-domain>` if it's part of a Deployment or DaemonSet), but the problem of discovering and maintaining a list of IP address/DNS names for these ephemeral Pods remains an issue.

This process of getting the IP address/DNS names of other applications is known as _service discovery_ and there are tools out there that facilitates this. But instead of introducing another tool into the pile of tools, Kubernetes has a concept of _[Services](https://kubernetes.io/docs/concepts/services-networking/service/)_ that allows for DNS-base service discovery - allowing communication between different applications (running on 1 or more Pods) using structured DNS names, without external tooling. In short, Services allows you to refer to some service (here, it's a selection of Pods) using a single, predictable DNS name.

> We will use the capital-case 'Service' to refer to the Kubernetes resource, and lower-case 'service' to refer to a generic networked service backed by an API server.

In our case, we can layer a Service resource on top of all the Pods that are part of our Deployment, to give it a single DNS name.

One benefit of using a Service is that even if Pods are terminated due to downscaling, or restarted with a different IP address due to an update, the DNS name will remain the same. Having a single DNS names saves other services the need to update their records to use a new IP address - they can use the same DNS name as before.

The set of Pods included under a Service is determined by the Service's selector, similar to the selector in a Deployment object.


```yaml
...
            - name: http
              containerPort: 8080

---

apiVersion: v1
kind: Service
metadata:
  name: echo
spec:
  selector:
    app: echo
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080

```

```console
$ kubectl apply -f echo.yaml                
deployment.apps/echo unchanged
service/echo created
```

```console
$ kubectl get services
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
echo         ClusterIP   10.96.106.219   <none>        80/TCP    88s
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP   42h
```

Kubernetes will assign the service a _cluster IP_ (a private IP address which can only be reached from within the cluster).

More importantly, Kubernetes also adds an A or AAAA DNS record for the Service to the internal DNS resolver running within the cluster. As a result, all Services have a fixed DNS name. The DNS name itself follows the structure `<service>.<namespace>.svc.<cluster-domain>`, where `<namespace>` defaults to `default`, and `<cluster-domain>` defaults to `cluster.local`.

The internal DNS resolver (or simply 'DNS server') is automatically deployed by Kubernetes and exists as one or more Pod(s) and a Service object.

Previously, a DNS server called `kube-dns` was used; but that has since been replaced by _[CoreDNS](https://coredns.io/)_. Since this DNS has a Service attached, it, too, will have an IP address and DNS name. The Service `metadata.name` field for both the legacy `kube-dns` server and the current CoreDNS server are both `kube-dns`, a name that is kept for interoperability. So if you run `kubectl get services --namespace=kube-system` and see an entry for `kube-dns`, it doesn't necessarily mean it is running the legacy DNS server.

```console
$ kubectl get svc --namespace=kube-system
NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)               
kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP
```

The `kubelet`s running on each node would then configure the `/etc/resolv.conf` file in every container running on the node to use this internal DNS resolver when resolving DNS names. In the `/etc/resolv.conf`, the `kubelet` would also add the search domains `<namespace>.svc.<cluster-domain>`, `svc.<cluster-domain>`, and `<cluster-domain>`, which means these containers can refer to the Service simply using its name (i.e. you can simply `curl echo` the same way you run `curl localhost:5000`).

So for our `echo` Service, we should be able to communicate with any Pods with the label `app=echo` using the DNS name `echo.default.svc.cluster.local` or `echo`.

Under the hood, Kubernetes also creates an Endpoint object _with the same name_.

```console
$ kubectl get endpoints
NAME         ENDPOINTS          AGE
echo         10.244.0.10:8080   2m50s
kubernetes   172.18.0.2:6443    42h
```

An Endpoint object lists the set of endpoints (IP address and port) which your Service sends traffic to. In our case, the set of endpoints corresponds to the IP addresses and port of the echo service's Pods. Whenever a Pod is started, terminated, or restarted, the list of endpoints in the Endpoints object is updated. When the internal DNS resolver resolves a name, it resolves the IP addresses listed in the Endpoint object.

If that's the only uses of Services and Endpoints, the creators or Kubernetes could have very well combined the Endpoints object into the Service object. But the reason the Service and Endpoint objects are decoupled from each other is because you can define a Service with no selectors, and manually define the corresponding Endpoint pointing to services outside of the cluster and/or outside of the namespace. You may want to do this because you want to use different databases (a service) for local testing, development, and production. Decoupling Services and Endpoints allows you to keep the same Services manifest and change only the Endpoints object for each environment.

We've provided a lot of details here, but the take home message is this - by creating a Service object for our `echo` service, we are giving the `echo` service a single DNS name.

## Exposing Our Service with Ingress

Our `echo` service now has a DNS name that can be resolved using the _internal_ DNS server, but how do we reach the Service from outside of the cluster? For example, from our local machine?

We have 4 options:

1. Deploy a (temporary) Pod inside the cluster and SSH into the Pod (using `kubectl exec --stdin --tty <pod> -- /bin/bash`). From the Pod (which is within the cluster), you can communicate with the desired Service.
2. There are actually many 4 different types of Services - `ClusterIP`, [`NodePort`](https://kubernetes.io/docs/concepts/services-networking/service/#nodeport), `LoadBalancer`, and `ExternalName`. The default Service type (the one we've been using) is `ClusterIP`, which assigns each Service a cluster-internal IP address that is only reachable from within the cluster. But we can actually change the Service type to `NodePort`, which exposes the Service on a random port (range defaults to 30000-32767) on each node. With Service type of `NodePort`, a cluster IP is still assigned to the Service, but now Kubernetes will instruct the `kubelet` running on each node to add routing rules so that when a client sends a request to the specified port on _any_ of the nodes, that request will be proxied to the cluster IP of the Service.
3. When using a cloud provider such as AWS, you can also use the Service type `LoadBalancer`, which causes Kubernetes to use an Ingress Controller (more about them below) to provision a network-layer load balancer that routes traffic to from the load balancer to the Service's internal cluster IP.
4. Configure an _[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)_ and _[Ingress Controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/)_, which, together, acts as a _gateway_ between the cluster's network and an external network (e.g. the Internet in the context of cloud providers, and our host's local network in the context of our local cluster)

You'd choose the first option when you want all your Services to _not_ be exposed outside the network (e.g. for security reasons) but you want to perform a one-off test or debug session. But since we want our services to be exposed externally (in development to our local network, and in production to the Internet), this is not an option we'll consider here.

The second is acceptable for local use, as its reasonable to assume that you have access and information about the Kubernetes nodes. But this is not ideal in a production environment because you'd have to make your Kubernetes nodes public, or have to manually provision a load balancer that sits in front of the nodes. Furthermore, HTTP and HTTPS typically uses the ports `80` and `443`; with this set up, you're forcing HTTP(S) traffic to use an unconventional port.

The third option is good but it's only available on cloud providers - we cannot use it locally.

An Ingress is a set of routing rules that routes HTTP and HTTPS requests from external clients to Services running within the cluster. Ingresses are unique amongst Kubernetes resources in that they exist within the cluster but have external IP addresses. In the case of cloud providers (e.g. AWS, Google Cloud), these external IP addresses must first be provisioned (usually at a cost) and assigned; in the case of a self-managed cluster, the Ingress is given an available address from the range of private IP addresses.

> Note that using Service of type `NodePort` routes traffic at the network (IP) layer. Routing is done based on the IP address only - it cannot read anything inside the IP packet. An Ingress works at the application layer (typically HTTP(S)) and can route traffic based on IP address, HTTP request headers (e.g. [`Host`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host)), HTTP method, path, and anything else within the HTTP message.
> Also because Ingresses work with HTTP load balancers at the application layer, you can also perform TLS/SSL termination for multiple hostnames.

## Ingress Controller

An Ingress is just a set of rules; an Ingress Controller is the software that fulfills those rules. Specifically, an Ingress Controller runs within the cluster and configures Services and load balancer(s) based on the Ingress rules. The load balancer it configures can be a software load balancer running within the cluster (as is the case with the [NGINX Ingress Controller](https://github.com/nginxinc/kubernetes-ingress/)), a hardware load balancer, or a cloud load balancer (e.g. [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/)).

For example, to expose a Service outside the cluster on AWS's Elastic Kubernetes Service (EKS), you need to [install](https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html) the [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/). The AWS Load Balancer Controller interacts with AWS API to create an Application Load Balancers (ALBs), [Listeners](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html), and [Listener rules](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-update-rules.html) for any Ingress objects created. Once set up, the ALB accepts external traffic and direct them to the correct Service within the cluster based on the rules. You can find out more about how this controller works at [How AWS Load Balancer controller works](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/how-it-works/).

With the [NGINX Ingress Controller](https://github.com/nginxinc/kubernetes-ingress/)), how does it work ?? https://kubernetes.github.io/ingress-nginx/deploy/baremetal/

?? External load balancer sends traffic to an Ingress Controller Service which does the actual load balancing ??

## Configuring Ingress in `kind`

To make Ingress work in `kind`, you'd first deploy an Ingress Controller as a `NodePort` Service within the cluster, and add an `extraPortMappings` field in the `kind` configuration file that will forward traffic sent to a specific host port to the node port on a specific node, which will, in turn, reach the Ingress Controller Service, which, in turn, will route the request to the appropriate Pod.

Update `cluster.yaml` to:

```yaml
apiVersion: ctlptl.dev/v1alpha1
kind: Cluster
product: kind
registry: registry
kindV1Alpha4Cluster:
  name: acme
  nodes:
  - role: control-plane
    kubeadmConfigPatches:
    - |
      kind: InitConfiguration
      nodeRegistration:
        kubeletExtraArgs:
          node-labels: "ingress-ready=true"
    extraPortMappings:
    - containerPort: 80
      hostPort: 80
      protocol: TCP
    - containerPort: 443
      hostPort: 443
      protocol: TCP
```

```console
% ctlptl apply -f cluster.yaml
Deleting cluster kind-acme because desired Kind config does not match current.
Cluster config diff:   &v1alpha4.Cluster{
        TypeMeta: {},
        Name:     "acme",
-       Nodes:    nil,
+       Nodes: []v1alpha4.Node{
+               {
+                       Role:                 "control-plane",
+                       ExtraPortMappings:    []v1alpha4.PortMapping{{...}, {...}},
+                       KubeadmConfigPatches: []string{"kind: InitConfiguration\nnodeRegi"...},
+               },
+       },
        Networking:   {},
        FeatureGates: nil,
        ... // 5 identical fields
  }

Deleting cluster "acme" ...
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

Thanks for using kind! 😊
Switched to context "kind-acme".
 🔌 Connected cluster kind-acme to registry registry at localhost:54934
 👐 Push images to the cluster like 'docker push localhost:54934/alpine'
cluster.ctlptl.dev/kind-acme created
```

```console
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
configmap/ingress-nginx-controller created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
service/ingress-nginx-controller-admission created
service/ingress-nginx-controller created
deployment.apps/ingress-nginx-controller created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
serviceaccount/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created
```

```yaml
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: echo
spec:
  rules:
  - host: auth.goji.local
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: echo
            port:
              number: 80
```


This means to switch from local testing to production, all we need to change is to change the Ingress class.
