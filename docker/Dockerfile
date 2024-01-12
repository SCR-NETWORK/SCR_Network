# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/SCR-NETWORK/SCR-NETWORK

WORKDIR /go/src/github.com/SCR-NETWORK/SCR-NETWORK
RUN mkdir -p /go/src/github.com/SCR-NETWORK/SCR_Network

WORKDIR /go/src/github.com/SCR-NETWORK/SCR_Network

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache SCR-NETWORK dependencies
# Cache SCR_Network dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o SCR-NETWORK .
RUN go build $FLAGS -o SCR_Network .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/SCR-NETWORK/SCR_Network/SCR-NETWORK /app/
COPY --from=build /go/src/github.com/SCR-NETWORK/SCR_Network/infrastructure/config/sample-SCR-NETWORK.conf /app/
COPY --from=build /go/src/github.com/SCR-NETWORK/SCR_Network/SCR_Network /app/
COPY --from=build /go/src/github.com/SCR-NETWORK/SCR_Network/infrastructure/config/sample-SCR_Network.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
