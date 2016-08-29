FROM docker.io/alpine:3.5

MAINTAINER bobliu0909@gmail.com

RUN mkdir -p /opt/app/gtlservice/etc

COPY gtlappserver /opt/app/gtlservice

COPY etc/config.yaml /opt/app/gtlservice/etc/config.yaml

RUN chmod +x /opt/app/gtlservice/gtlappserver

WORKDIR /opt/app/gtlservice

ENV SERVICE_NAME service1

ENV SERVICE_HOST http://127.0.0.1:5000

ENV API_BIND :8119

CMD ["./gtlappserver"]

EXPOSE 8119