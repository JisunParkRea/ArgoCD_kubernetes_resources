FROM golang:1.9

RUN mkdir /go/src/app
COPY main.go /go/src/app

EXPOSE 8000

CMD ["go", "run", "/go/src/app/main.go"]