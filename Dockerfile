# Build the executable
FROM golang:1.16-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -mod=readonly -v -o api

# Production container
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y && \
    apt-get install -y perl wget libfontconfig1 && \
    wget -qO- "https://yihui.name/gh/tinytex/tools/install-unx.sh" | sh  && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH="${PATH}:/root/bin"
RUN fmtutil-sys --all
ENV PROJECT_DIR /app
WORKDIR /app
# install only the packages you need
# this is the bit which fails for most other methods of installation
RUN tlmgr install pgf fontawesome5 koma-script cmap ragged2e everysel tcolorbox \
    enumitem ifmtarg dashrule changepage multirow environ paracol lato \
    fontaxes

COPY --from=builder app/ /app/

CMD ["/app/api"]