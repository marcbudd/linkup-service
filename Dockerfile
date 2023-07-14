FROM golang:1.18.0

WORKDIR /usr/src/app

RUN apt-get -y update

# Install Swagger UI dependencies
RUN apt-get install -y npm
RUN npm install swagger-ui-dist

# Build the Go application
COPY . .
RUN go mod tidy
RUN go build -o main .

EXPOSE 3000

CMD ["./main"]