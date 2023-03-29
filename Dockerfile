FROM golang:1.20.2-alpine AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY go.mod  .
COPY go.sum .
RUN apk --no-cache add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
# RUN mv migrate.linux-amd64 $GOPATH/bin/migrate

RUN go mod download 

COPY . .

RUN go mod tidy

RUN GOOS=linux go build -o main -ldflags=-X=main.version=${VERSION} ./cmd/talent/.

FROM alpine
WORKDIR /go/bin
COPY --from=builder /go/src/app/main /go/src/app/.env  /go/bin/

ENV PATH="/go/bin:${PATH}"
ENTRYPOINT ["/go/bin/main", ".env"]
