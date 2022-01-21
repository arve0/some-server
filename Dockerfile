FROM golang as builder

COPY main.go .
RUN go build main.go

FROM debian:11-slim

COPY --from=builder /go/main /opt

ENTRYPOINT ["/opt/main"]