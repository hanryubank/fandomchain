FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /fandomchain
RUN cd /fandomchain && make fandom

FROM alpine:latest

LABEL maintainer="etienne@fandomchain.com"

WORKDIR /fandomchain

COPY --from=builder /fandomchain/build/bin/fandom /usr/local/bin/fandom

RUN chmod +x /usr/local/bin/fandom

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/fandom"]

CMD ["--help"]
