FROM golang:1.18.0

WORKDIR /usr/src/app

RUN apt-get -y update

# Install Node.js for Swagger UI setup
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -

# Generate Swagger docs
RUN go install -mod=mod github.com/swaggo/swag/cmd/swag
COPY . .
RUN swag init

# Install Swagger UI dependencies
RUN npm install swagger-ui-dist

# Build the Go application
RUN go mod tidy
RUN go build -o main .

EXPOSE 3000

CMD ["./main"]