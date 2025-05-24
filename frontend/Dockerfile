FROM golang:latest 

WORKDIR /app

COPY . .

RUN go build -o frontend .

EXPOSE 8080
CMD ["./frontend"]