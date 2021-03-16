FROM golang:rc-alpine3.12 AS build

RUN apk add --no-cache git && \
    git clone https://github.com/WhaleMountain/filegea.git /tmp/filegea && \
    cd /tmp/filegea && \
    go build

FROM alpine:3.12

RUN mkdir -p /opt/filegea/Data
COPY --from=build /tmp/filegea/filegea /opt/filegea/filegea
COPY --from=build /tmp/filegea/config.toml /opt/filegea/config.toml

RUN addgroup -S -g 1001 filegea && \
    adduser -S filegea -g filegea --uid 1001 && \
    chown filegea:filegea -R /opt/filegea

USER filegea

EXPOSE 1270
CMD ["/opt/filegea/filegea"]
