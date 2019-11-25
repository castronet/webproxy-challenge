# webproxy challenge

This repository is created specifically for a challenge.

In this document I will explain how to run the application(s) and it will also contains an explaining of the
decisions taken and a list of requirements.

## How to run

The challenge requires a **one single command** to run the whole application.
To run this application you need to be on the root folder and run:

```
docker-compose up
```

docker-compose will build two docker images and then bring it up. When docker-compose finish their work you will be
able to see logs on the command line.



## ssl certificate
You can request valid certificates to a 3rd party company like Let's Encrypt but there is an option to create self
signed certificates in command line

OpenSSL certificate generated with:
Generate the private key
openssl genrsa -out server.key 2048

Then generate the certificate
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650


## nginx 

### option choosen

### config


## List of requirements
 - docker-compose
 - docker engine
