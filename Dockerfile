FROM golang:latest
WORKDIR /app
RUN go get -d -v github.com/gin-gonic/gin
COPY cmd/sonar-badges/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

ENV GIN_MODE=release

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./main"]