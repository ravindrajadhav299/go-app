FROM golang:1.24.4-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main ./server.go
RUN chmod +x main
EXPOSE 4040
CMD [ "./main" ]