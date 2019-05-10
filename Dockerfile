FROM golang:1.12.5 as  builder

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/0.5.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/yorick1101/currency_ex_board
WORKDIR /go/src/github.com/yorick1101/currency_ex_board

COPY  . .
RUN dep ensure -vendor-only 
RUN env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o /bin/rate-server cmd/rate-server/rate-server.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/rate-server .
CMD ["/bin/sh","-c","/app/rate-server"]