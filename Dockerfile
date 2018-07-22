FROM golang:1.10 as build

MAINTAINER Masaki Ikeda <ikeda.masaki@gmail.com>

ENV GOPATH /go
WORKDIR /go/src/app
COPY . .

RUN make install


FROM alpine

COPY --from=build /go/bin/app /usr/local/bin/

CMD ["app"]
