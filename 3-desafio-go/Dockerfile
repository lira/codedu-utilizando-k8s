FROM golang:alpine as builder
RUN apk update
COPY src/webserver .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver .

FROM scratch
COPY --from=builder /go/webserver /usr/bin/webserver
EXPOSE 8000
CMD ["webserver"]