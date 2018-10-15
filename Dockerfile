FROM golang:1.11.1-alpine3.8

COPY ./simple-go-http-rest /app/simple-go-http-rest
RUN chmod +x /app/simple-go-http-rest

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/simple-go-http-rest