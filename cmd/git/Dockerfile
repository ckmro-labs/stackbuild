FROM alpine:3.7
RUN apk add --no-cache ca-certificates git git-lfs openssh curl perl

ADD scripts/* /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/clone"]