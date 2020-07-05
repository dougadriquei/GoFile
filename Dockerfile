FROM golang:1.10

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dougadriquei/desafioneoway

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .


# This container exposes port 8080 to the outside world - Docker
EXPOSE 8085

# This container exposes port 8080 to the outside world - Local
# EXPOSE 8080

# Run the executable
CMD ["desafioneoway"]

# Download all the dependencies
RUN go get

# Install the package
RUN go install -v 

# Build
RUN go build

# Test
#TODO - Verificar: purchase_test.go:48: CreatePurchase (dao)
# RUN go get github.com/stretchr/testify/assert
# RUN go get github.com/jinzhu/gorm/dialects/postgres
# RUN go get github.com/dougadriquei/desafioneoway/storage

# RUN CGO_ENABLED=0 go test ./...






