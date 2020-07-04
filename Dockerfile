# Base build image
FROM golang:1.13.7 AS builder
ENV GO111MODULE=on
WORKDIR /dougaq/go/src/desafioneoway
COPY [ "go.mod", "go.sum", "./" ]
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o main .

# Server image
FROM alpine:latest
LABEL maintainer="Douglas Queiroz"
LABEL description="Docker container for desafioneoway."
# RUN apk update && apk upgrade && apk add bash  ## For debug
RUN apk --no-cache add ca-certificates
EXPOSE 8095
WORKDIR /desafioneoway


COPY --from=builder /dougaq/go/src/desafioneoway/main ./
CMD [ "./main" ]