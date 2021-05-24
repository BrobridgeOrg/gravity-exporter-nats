FROM golang:1.15.3-alpine3.12 AS builder
WORKDIR /
COPY . .

RUN apk add --update git
RUN CGO_ENABLED=0 go build -installsuffix cgo -o /gravity-exporter-nats ./cmd/gravity-exporter-nats/gravity-exporter-nats.go

FROM alpine:3.12
COPY --from=builder /gravity-exporter-nats /gravity-exporter-nats
COPY ./configs .

USER 1001
EXPOSE 44144
CMD ["/gravity-exporter-nats"]
