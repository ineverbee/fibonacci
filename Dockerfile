#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/server -v ./cmd/server
RUN go build -o /go/bin/client -v ./cmd/client

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/client /client
COPY --from=builder /go/bin/server /server
ENTRYPOINT /server
LABEL Name=fibonacci Version=0.0.1
