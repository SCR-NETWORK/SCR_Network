ARG SCRPAD_IMAGE
ARG SCRPADMINER_IMAGE

FROM ${SCRPAD_IMAGE} as SCR-NETWORK
FROM ${SCRPADMINER_IMAGE} as SCRminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/SCR-NETWORK/SCR-NETWORK

WORKDIR /go/src/github.com/SCR-NETWORK/SCR-NETWORK
ARG SCRPAD_IMAGE
ARG SCRPADMINER_IMAGE

FROM ${SCRPAD_IMAGE} as SCR_Network
FROM ${SCRPADMINER_IMAGE} as SCRminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/SCR-NETWORK/SCR_Network

WORKDIR /go/src/github.com/SCR-NETWORK/SCR_Network

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=SCR-NETWORK /app/ /app/
COPY --from=SCRminer /app/ /app/
COPY --from=SCR_Network /app/ /app/
COPY --from=SCRminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/SCR-NETWORK/SCR_Network/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
