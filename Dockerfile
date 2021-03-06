FROM golang:1.18 AS builder
WORKDIR /app
COPY *.go ./
RUN go mod init testmod && go mod tidy && go get -u && CGO_ENABLED=0 go build -v main.go

FROM alpine
LABEL maintainer="sagit"
ENV WORKDIR /
COPY --from=builder /app/main /
COPY *.sh /
RUN chmod +x main
RUN chmod +x *.sh
VOLUME [ "/backup" ]
ENTRYPOINT ["./main"]