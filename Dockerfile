FROM alpine:3.4

ADD ./test_travis /test_travis

ENTRYPOINT ["/test_travis"]
