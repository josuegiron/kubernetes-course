FROM golang:latest AS builder

WORKDIR /build
ADD ./ /build/
RUN env GOOS=linux GOARCH=386 go build -o main .

FROM alpine:latest AS runner

WORKDIR /app
RUN adduser -S -D -H -h /app chamuske && chown -R chamuske /app
USER chamuske
COPY --from=builder /build/main /app/
EXPOSE 8080

CMD ["./main"]
