FROM golang:1.19.3-alpine3.16
RUN mkdir customer
COPY . /customer
WORKDIR /customer
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9000