# webproxy challenge

## ssl certificate
You can request valid certificates to a 3rd party company like Let's Encrypt but there is an option to create self signed certificates in command line

OpenSSL certificate generated with:
Generate the private key
openssl genrsa -out server.key 2048

Then generate the certificate
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650


## nginx 

### option choosen

### config
