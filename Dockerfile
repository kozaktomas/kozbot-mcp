FROM golang:1.25 AS backend
ENV CGO_ENABLED=0
ADD . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -v -o kozbot .

FROM ghcr.io/open-webui/mcpo:main
RUN apt-get update && \
    apt-get install -y openssl tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ADD Dockerfile /Dockerfile
ADD config.json /root/.mcpo/config.json
COPY --from=backend /app/kozbot /bin/kozbot

ENTRYPOINT ["mcpo", "--config", "/root/.mcpo/config.json"]