FROM golang:1.25.3 AS builder

WORKDIR /app

COPY . .

RUN go build -o go1f


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/go1f .

COPY web ./web/

EXPOSE 7540

ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db
ENV TODO_PASSWORD=""
CMD ["./go1f"]
