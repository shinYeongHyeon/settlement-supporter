#Dockerfile
FROM golang:1.17-alpine

WORKDIR /app
COPY docker .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex && \
    go get github.com/gofiber/fiber/v2

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]