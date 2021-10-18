FROM golang:1.17-alpine3.14 as build-env
RUN apk add git gcc
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o te
FROM alpine:3.14
RUN addgroup -S app && adduser -S app -G app
RUN mkdir /app && chown -R app /app && chgrp -R app /app
USER app
WORKDIR /app
COPY --from=build-env /app/te /app
ENTRYPOINT ["./te"]
