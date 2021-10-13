# Build the executable
FROM golang:1.16-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -mod=readonly -v -o api

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
ENV PROJECT_DIR /app
ENV ENV_MODE PRD
WORKDIR /app
# Download altacv class from the author github
RUN wget -q "https://raw.githubusercontent.com/liantze/AltaCV/main/altacv.cls"
# install only the packages you need
RUN tlmgr install pgf fontawesome5 koma-script cmap ragged2e everysel tcolorbox \
    enumitem ifmtarg dashrule changepage multirow environ paracol lato \
    fontaxes accsupp

COPY --from=builder /app/api /app/api
ADD templates ./templates/
COPY run.sh ./

RUN chmod +x run.sh
RUN ./run.sh
CMD ["/app/api"]
