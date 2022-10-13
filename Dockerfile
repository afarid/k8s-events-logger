FROM golang:1.19 as builder
WORKDIR /build
COPY go.mod go.mod
COPY go.sum go.sum
RUN  go mod download
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app main.go

FROM scratch as prod
COPY --from=builder /build/app .
ENTRYPOINT ["./app"]