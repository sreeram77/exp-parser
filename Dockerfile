FROM alpine:latest

COPY build/main /main

RUN chmod +x /main

ENTRYPOINT ["/main"]
