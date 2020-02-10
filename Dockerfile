FROM golang:latest

WORKDIR /go/src/cmd/golangcomicstore/

COPY . .

RUN go get -d -t  -v ./...

RUN go install -i -v ./...

EXPOSE 8081
CMD ["cmd/golangcomicstore"]