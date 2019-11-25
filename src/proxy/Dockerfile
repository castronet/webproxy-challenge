FROM nginx

WORKDIR /proxy
RUN mkdir -p /proxy
COPY ./src/webapp/server.key /proxy/
COPY ./src/webapp/server.crt /proxy/

RUN rm /etc/nginx/conf.d/default.conf
COPY ./src/proxy/webproxy.conf /etc/nginx/conf.d/
