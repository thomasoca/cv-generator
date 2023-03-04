# Build the Go API
FROM golang:1.19-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY main.go ./

RUN CGO_ENABLED=0 go build -mod=readonly -v -o api

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
    apt-get install -y perl wget libfontconfig1 && \
    wget -qO- "https://yihui.name/gh/tinytex/tools/install-unx.sh" | sh  && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH="${PATH}:/root/bin"
RUN fmtutil-sys --all

WORKDIR /app
# Download altacv class from the author github
RUN wget -q "https://raw.githubusercontent.com/liantze/AltaCV/main/altacv.cls"
# Download missing class
RUN wget -q "http://tug.ctan.org/tex-archive/macros/latex/contrib/extsizes/extarticle.cls"
# Update tlmgr
RUN tlmgr update --self
# install only the packages you need
RUN tlmgr install pgf fontawesome5 koma-script cmap ragged2e everysel tcolorbox \
    enumitem ifmtarg dashrule changepage multirow environ paracol lato \
    fontaxes accsupp tikzfill

COPY --from=builder /app/api /app/api
COPY --from=node_builder /app/build/ /app/build/
RUN chmod +x ./api
COPY examples ./examples/
COPY templates ./templates/
COPY scripts/run.sh ./

RUN chmod +x run.sh
RUN ./run.sh
CMD ["/app/api", "serve"]

