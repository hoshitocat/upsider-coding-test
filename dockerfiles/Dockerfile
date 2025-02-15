FROM golang:1.24.0

ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY . /app
WORKDIR /app
RUN go mod download
