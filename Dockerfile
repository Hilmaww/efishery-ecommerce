FROM golang:alpine AS builder

#BUILD
COPY . /app
WORKDIR /app

#    && apk upgrade \
#    && apk add --no-cache \
#    ca-certificates \
#    && update-ca-certificates 2>/dev/null || true
ENV DB_URL="host=172.19.0.1 port=5432 user=HilmiFaww dbname=template1 sslmode=disable"

RUN apk  update
RUN apk add --no-cache git

RUN mkdir app && \
    mkdir app/resources && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /app/echo
RUN ls app

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]

# SERVE
#FROM busybox
#
#COPY --from=builder /app/ /app
#WORKDIR /app
#EXPOSE 1323
#CMD ["./echo"]