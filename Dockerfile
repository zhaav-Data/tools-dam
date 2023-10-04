FROM golang:1.16-alpine as builder

RUN mkdir /build
ADD *.go /build/
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

# RUN go build -o main
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

EXPOSE 8020

FROM alpine:3.11.3
COPY --from=builder /build/main .

# executable
ENTRYPOINT [ "./main" ]
