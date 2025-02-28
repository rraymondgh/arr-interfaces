FROM alpine:latest

LABEL org.opencontainers.image.source = "https://github.com/rraymondgh/arr-interfaces"
RUN ["apk", "--no-cache", "add", "bind-tools", "curl"]
EXPOSE 3335
COPY arr-interfaces /
ENTRYPOINT ["/arr-interfaces"]