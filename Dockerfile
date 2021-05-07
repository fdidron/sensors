FROM golang:1.16.3-alpine3.13

ENV APP_DIR = $GOPATH/src/github.com/fdidron/sensors/

RUN apk update
RUN apk add --no-cache nodejs npm git make gcc musl-dev sqlite

RUN git clone https://github.com/fdidron/sensors $GOPATH/src/github.com/fdidron/sensors

WORKDIR $APP_DIR
RUN go build *.go
RUN cd ui && npm i && npm run build

EXPOSE 8080
CMD cd $APP_DIR && ./main
