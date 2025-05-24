From golang:lates

WORKDIR /app

COPY . .

RUN go build main.go -o auth-service .

EXPOSE 8081
CMD ["./auth-service"]