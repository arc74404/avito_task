FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN ["go", "mod", "download"]

COPY server/ ./server/
COPY internal/ ./internal/
COPY server/static/ ./static/  

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN ["go" ,"build" , "-o", "server-app" ,"./server/main.go"]

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server-app .
COPY --from=builder /app/static ./static/ 
CMD ["./server-app"]