# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS is-builder

ENV PACKAGES curl make git libc-dev bash gcc linux-headers
RUN apk add --no-cache $PACKAGES

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOFLAGS="-buildvcs=false"

# cache go modules
COPY go.mod /go.mod
COPY go.sum /go.sum
RUN go mod download
RUN if [ -d "./cosmos-sdk" ]; then go mod edit -replace github.com/cosmos/cosmos-sdk=./cosmos-sdk; fi

# Copy in the repo under test
ADD . /interchain-security

WORKDIR /interchain-security

# Do not specify version here. It leads to odd replacement behavior 
RUN if [ -d "./cosmos-sdk" ]; then go mod edit -replace github.com/cosmos/cosmos-sdk=./cosmos-sdk; fi

# Install interchain security binary
RUN make install

# Get Hermes build
FROM informalsystems/hermes:1.2.0 AS hermes-builder

FROM --platform=linux/amd64 fedora:36
RUN dnf update -y
RUN dnf install -y which iproute iputils procps-ng vim-minimal tmux net-tools htop jq python3-pip
USER root

# Copy in Tendermock and install requirements
ADD ./tests/integration/tendermock/requirements.txt /requirements.txt
RUN pip install -r /requirements.txt
ADD ./tests/integration/tendermock/ /tendermock


COPY --from=hermes-builder /usr/bin/hermes /usr/local/bin/

COPY --from=is-builder /go/bin/interchain-security-pd /usr/local/bin/interchain-security-pd
COPY --from=is-builder /go/bin/interchain-security-cd /usr/local/bin/interchain-security-cd
COPY --from=is-builder /go/bin/interchain-security-cdd /usr/local/bin/interchain-security-cdd


# Copy in the shell scripts that run the testnet
ADD ./tests/integration/testnet-scripts /testnet-scripts

# Copy in the hermes config
ADD ./tests/integration/testnet-scripts/hermes-config.toml /root/.hermes/config.toml
