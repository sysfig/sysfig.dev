
You may be tempted to test out the echo server by sending a request to `localhost:8080`; if you did, you'll be met with the error:

```console
$ curl -i http://localhost:8080/\?q=hello
curl: (7) Failed to connect to localhost port 8080: Connection refused
```

This is because the Docker container doesn't bind its port to your local machine. Instead, Docker automatically creates a [bridge](https://en.wikipedia.org/wiki/Bridging_(networking)) network assigns each container its own private IP address from a range (the default is `172.17. 0.0/16`).

You can see a list of networks that Docker creates by running the command `docker network ls`.

```console
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
a06bc031e4de   bridge    bridge    local
0aa11a696cd7   host      host      local
082d716cfc5d   kind      bridge    local
fe7a10b4fc54   none      null      local
```

We can get the private IP address that Docker has assigned by inspecting the data pertinent to the container using the [`docker inspect`](https://docs.docker.com/engine/reference/commandline/inspect/) command.

```console
$ docker inspect echo
[
  {
    "Id": "53316395d7bc399c57d50f048afc224dfce273d3e5e35753ba222f100dc2b200",
    ...
  }
]
```

`docker inspect` prints the data in a large JSON object; but we are only interested in the IP address field, which we can find at `[].NetworkSettings.Networks.bridge.IPAddress`. In my case, it reads `172.17.0.2`, although you may find a different address.

For a more compact output, use the `-f`/`--format` option to select only the fields from the JSON object that you are interested in. The string you pass in is a template format string from Go's [`text/template`](https://golang.org/pkg/text/template/) package.

```console
$ docker inspect --format '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' echo
172.17.0.2
```
