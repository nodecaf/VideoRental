## BUILD IMAGE
FROM golang:1.22 AS builder

WORKDIR /app

COPY . /app/
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o bbuster .

## RUN IMAGE
FROM scratch
WORKDIR /app
COPY --from=builder /app/bbuster /app/bbuster
COPY --from=builder /app/pkg/client/templates/* /app/pkg/client/templates/
EXPOSE 8080 50080

ENTRYPOINT ["/app/bbuster"]


