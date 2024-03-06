ARG DMGCOIN_IMAGE
ARG DMGCOINMINER_IMAGE

FROM ${DMGCOIN_IMAGE} as dmgcoin
FROM ${DMGCOINMINER_IMAGE} as dmgcoinminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/dmgcoin/dmgcoin

WORKDIR /go/src/github.com/dmgcoin/dmgcoin

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=dmgcoin /app/ /app/
COPY --from=dmgcoinminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/dmgcoin/dmgcoin/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]