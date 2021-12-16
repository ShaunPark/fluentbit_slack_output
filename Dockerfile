FROM golang:1.17 as gobuilder

WORKDIR /go/src/github.com/ShaunPark/fluentbit_slack_output

ENV GOOS=linux\
    GOARCH=amd64

COPY . .

RUN go mod edit -replace github.com/fluent/fluent-bit-go=github.com/fluent/fluent-bit-go@master && go mod tidy && go build -buildmode=c-shared -o out_prettyslack.so .

FROM fluent/fluent-bit:1.8.11

COPY --from=gobuilder /go/src/github.com/ShaunPark/fluentbit_slack_output/out_prettyslack.so /fluent-bit/bin/
COPY --from=gobuilder /go/src/github.com/ShaunPark/fluentbit_slack_output/fluent-bit.conf /fluent-bit/etc/
COPY --from=gobuilder /go/src/github.com/ShaunPark/fluentbit_slack_output/plugins.conf /fluent-bit/etc/

EXPOSE 2020

# CMD ["/fluent-bit/bin/fluent-bit", "--plugin", "/fluent-bit/bin/out_multiinstance.so", "--config", "/fluent-bit/etc/fluent-bit.conf"]
CMD ["/fluent-bit/bin/fluent-bit", "--config", "/fluent-bit/etc/fluent-bit.conf"]