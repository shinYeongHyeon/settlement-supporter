#Dockerfile
FROM golang:1.18-alpine

WORKDIR /app
COPY docker .

RUN apk update && \
    apk add git && \
    go install github.com/cespare/reflex@latest

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]