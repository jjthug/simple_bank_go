#Build stage
FROM golang:1.19.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
#Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY db/migration ./db/migration
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]