FROM alpine

RUN apk add ca-certificates

ADD dummy-rest-api /usr/local/bin/
EXPOSE 32123

ENTRYPOINT ["/usr/local/bin/dummy-rest-api"]
