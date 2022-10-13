FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && go build .


FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata \
    && mkdir -p /data/linksaver

COPY --from=builder /src/linksaver /app/

ENTRYPOINT ["./linksaver"]