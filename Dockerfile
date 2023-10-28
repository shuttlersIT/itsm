FROM golang:1.19.0-alpine AS builder

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /intel

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

#RUN go get -d -v golang.org/x/net/html
#RUN go install golang.org/x/oauth2@latest
#RUN go install golang.org/x/oauth2/google@latest
#RUN go install github.com/google/tink/go/core/registry@v1.3.0-rc1

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
RUN go mod vendor
RUN go mod verify

COPY . ./
RUN ls -la ./*

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /intel

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /intel

COPY --from=builder /intel .


# create and set non-root USER
RUN addgroup -g 1001 docker && \
    adduser -S -u 1001 -G docker docker

RUN chown -R docker:docker /intel && \
    chmod 755 /intel

USER docker:docker


EXPOSE 9193

#RUN ["chmod", "+x", "./intel"]

ENTRYPOINT ["./intel"]
