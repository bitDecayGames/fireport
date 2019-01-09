FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/bitDecayGames/fireport/server/
ADD main .
CMD ["./main"]