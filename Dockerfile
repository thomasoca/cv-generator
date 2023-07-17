# Build the Go cv-generator
FROM golang:1.19-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal
COPY main.go ./

RUN CGO_ENABLED=0 go build -mod=readonly -v -o cv-generator

# Build the React application
FROM node:alpine AS node_builder
WORKDIR /app
COPY ./frontend/package*.json /app/
RUN npm install
ENV NODE_OPTIONS --openssl-legacy-provider
COPY ./frontend/ /app/
RUN npm run build

# Production container
FROM debian:buster-slim

# Install TinyTex for Latex compiler
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y && \
    apt-get install -y ca-certificates perl libfontconfig1 wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
RUN wget -q "https://yihui.org/tinytex/install-bin-unix.sh"
COPY --from=builder /app/cv-generator /app/cv-generator
COPY --from=node_builder /app/build/ /app/build/
RUN chmod +x ./cv-generator
COPY examples ./examples/
COPY templates ./templates/

RUN ./cv-generator install
ENV PATH="${PATH}:/root/bin"

ENTRYPOINT ["/app/cv-generator"]
