FROM node:lts-alpine as node-builder

WORKDIR /build

COPY . /build

RUN cd /build/assets/frontend \
    && yarn && yarn build

FROM golang:alpine as go-builder

WORKDIR /build

COPY --from=node-builder /build /build

RUN go build -o /build/status

FROM alpine:latest

COPY --from=go-builder /build/status /app/status

VOLUME [ "/etc/config.yml" ]

ENTRYPOINT [ "/app/status","-f","/etc/config.yml" ]