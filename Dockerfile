FROM golang:bullseye AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main src/main.go

FROM golang:bullseye
WORKDIR /root/
COPY --from=builder /app/main .

RUN chmod +x ./main

CMD ["./main"]
