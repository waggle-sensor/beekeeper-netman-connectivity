FROM golang:1.17 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o netman-connectivity

FROM scratch
COPY --from=builder /build/netman-connectivity /netman-connectivity
ENTRYPOINT [ "/netman-connectivity" ]
