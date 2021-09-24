FROM alpine:latest

RUN mkdir -p /file
COPY build/main /main
COPY file/input.yml /file
COPY file/output.yml /file

RUN chmod +x /main

ENTRYPOINT ["/main"]
