# Build stage
FROM golang:latest AS builder

WORKDIR /app
COPY . .

#strip debugging information and reduce binary size
RUN go build -ldflags="-s -w" -o main .


# Final stage
#minimal linux image
FROM alpine:3.18  

WORKDIR /app
#copy only compiled binary from Build stage
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]