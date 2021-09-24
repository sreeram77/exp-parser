FROM alpine:latest

RUN mkdir -p /file

COPY build/main /main
COPY file /file

RUN chmod +x /main

ENTRYPOINT ["/main"]
