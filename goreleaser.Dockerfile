FROM alpine:latest

LABEL org.opencontainers.image.source = "https://github.com/rraymondgh/arr-interfaces"
RUN ["apk", "--no-cache", "add", "bind-tools"]
EXPOSE 3335
ENTRYPOINT ["/arr-interfaces"]
COPY arr-interfaces /