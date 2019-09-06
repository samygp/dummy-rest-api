FROM golang:1.12.9-alpine3.10

LABEL maintainer="Sam <soysamygp@gmail.com>"

RUN set -x \
  && apk add --update --no-cache \
    git \
    make \
    ca-certificates \
  && echo "Installed OS deps."

WORKDIR /go/src/github.com/samygp/
COPY Makefile ./
ENV GO111MODULE on
RUN go mod init

COPY . .
RUN make local-build
EXPOSE 54345

CMD ["./dummy-rest-api"]
