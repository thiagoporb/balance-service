FROM golang:1-alpine3.12
LABEL Author="Thiago Tenório"
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]