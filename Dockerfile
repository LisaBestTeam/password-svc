FROM golang:1.15

WORKDIR /go/src/github.com/lisabestteam/password-svc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o /usr/local/bin/password .

FROM alpine:3.9

COPY --from=0 /usr/local/bin/password /usr/local/bin/password
RUN apk add --no-cache ca-certificates

ENTRYPOINT [ "password" ]