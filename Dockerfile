FROM golang:1.17 as gobuilder

WORKDIR /root

ENV GOOS=linux\
    GOARCH=amd64

COPY * /root/

RUN go mod edit -replace github.com/fluent/fluent-bit-go=github.com/fluent/fluent-bit-go@master && go mod tidy && make all

FROM fluent/fluent-bit:1.8.11

COPY --from=gobuilder /root/out_prettyslack.so /fluent-bit/bin/
COPY --from=gobuilder /root/fluentbit.conf /fluent-bit/etc/
COPY --from=gobuilder /root/plugins.conf /fluent-bit/etc/

EXPOSE 2020

# CMD ["/fluent-bit/bin/fluent-bit", "--plugin", "/fluent-bit/bin/out_multiinstance.so", "--config", "/fluent-bit/etc/fluent-bit.conf"]
CMD ["/fluent-bit/bin/fluent-bit", "--config", "/fluent-bit/etc/fluentbit.conf"]