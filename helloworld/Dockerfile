# syntax=docker/dockerfile:1
FROM golang:1.18.3 as build

ENV CGO_ENABLED 0
ENV GO_API_PORT 5000

WORKDIR /helloworld

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o helloworld

FROM alpine:3.16.0 as production
# Copy built binary from builder
COPY --from=build helloworld .
# Expose port
EXPOSE ${GO_API_PORT}
# Exec built binary
CMD ./helloworld