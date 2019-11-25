# webproxy challenge

This repository is created to solve a challenge.

In this document I will explain how to run all the microservices infrastructure. It will also contains an explanation of the
decisions taken and a list of requirements.

## How to run

The challenge requires a **one single command** to run the whole application.
To run this application you need to go to the challenge root folder and run:

```
$ docker-compose up
```

docker-compose will build two docker images and then bring it up. When docker-compose finish their work you will be
able to see logs on the command line.

Take care that web proxy will be listening on default HTTP and HTTPS ports that means that:
 - You maybe need super user permissions to run this infrastructure.
 - You may need to stop your web server.

## List of requirements
 - [docker engine](https://docs.docker.com/install/) - Engine where to run docker containers.
 - [docker-compose](https://docs.docker.com/compose/install/) - Tool for defining and running multi-container Docker applications.
 - For testing: [curl](https://curl.haxx.se/) - Or another tool or command line browser.


## Decisions taken

In this section I will talk about what decision I take and why.

### Engine

I chose _docker_ because currently microservices are running on docker. And this is because docker is a tool that brings the
opportunity to developers to test their work in a environment pretty similar to production. For infrastructure people work with
docker gives him the opportunity to put the focus on CI/CD pipelines (and lots of other stuff) that deliver the microservice
to a cluster instead of have to create a new customized server for each application.

### Docker orchestrator

My decision was to use _docker-compose_ because of two points:
 - The challenge says that they will appreciate the simplicity.
 - This tool lets us configure and run multiple containers with a single one command.

And on the other hand docker-compose configuration file is a YAML file that could be parsed and transformed to other
configuration files and of course it works with Dockerfile(s).


### Web Proxy

I worked with Apache HTTP Server in the past but 8 years ago I switched to Nginx because it was extremely fast but its
configuration is simple, modular and easy to understand and replicate.
Is for that reason that I chose Nginx.

#### SSL certificate
The first time I read the challenge I guess that HTTPS part has to be placed also on "web application" microservice to
ensure that the communication between process goes protected. After re-read it (I hope not to be wrong) I understood
that HTTPS part is only on proxy microservice. Is for that reason that if you check the commit log you will see code
removed from "web application" and the certificates moved permanently to proxy folder.

You can request valid certificates to a 3rd party company like Let's Encrypt but there is an option to create self
signed certificates in command line

To get a OpenSSL certificate we can follow this steps:

1. Generate the private key ```
openssl genrsa -out server.key 2048
```
2. Get the certificate```
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

The validity of the certificate will be 10 years and it is self-signed it means that on some browsers you will be to
accept the "warning insecurity advice".

#### Configuration

First of all I set a regular redirect from HTTP to HTTPS but because of we are running Nginx inside a docker I decided
to create an [upstream](http://nginx.org/en/docs/http/ngx_http_upstream_module.html) configuration on Nginx. This option
let us thanks to _docker-compose_ to connect directly on container with the other one without knowing beforehand the
IP address of the container.

We can test the challenge (after the deployment) with:
```
$ curl -kL http://127.0.0.1/hello
```
instead of
```
$ LOCAL_IP=`ip a | grep "inet [0-9]" | grep -v 172 | grep -v 127 | awk '{print $2}' | awk -F'/' '{print $1}'`
$ curl -kL http://$LOCAL_IP/hello
```

The line above expects to be executed on a \*nix OS.
