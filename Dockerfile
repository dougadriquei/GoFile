FROM golang:1.12-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dougadriquei/desafioneoway

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -d -v ./...


RUN go install -v ./...

# Build the Go app
# RUN go build -o ./out/desafioneoway .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



# This container exposes port 8080 to the outside world
EXPOSE 8191
CMD ["./main"]

FROM postgres:11.3-alpine as teste
ADD . /docker-entrypoint-initdb.d
COPY init.sql /docker-entrypoint-initdb.d/
# Run the binary program produced by `go install`

# FROM postgres
# ENV POSTGRES_USER postgress
# ENV POSTGRES_PASSWORD floripa@123
# ENV POSTGRES_DB postgres