FROM golang:1.20.5

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1

CMD ["tail", "-f", "/dev/null"]
