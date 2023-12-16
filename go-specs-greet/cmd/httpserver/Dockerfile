FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o svr go-specs-greet/cmd/httpserver/*.go
EXPOSE 8080
CMD [ "./svr" ]
