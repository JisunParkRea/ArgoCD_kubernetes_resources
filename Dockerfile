FROM golang:1.9

RUN mkdir /go/src/app
COPY main.go /go/src/app

CMD ["go", "run", "/echo/main.go"]