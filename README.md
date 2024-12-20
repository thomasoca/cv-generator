# cv-generator

A web app and CLI tool to generate beautiful, ATS-friendly LaTeX resume using available open source templates by filling a simple form (or using a JSON file in CLI mode).

## Live website

https://cv-generator-40m5.onrender.com

## Available templates

- A slightly modified [AltaCV](https://github.com/liantze/AltaCV). The original altacv latex class was written by LianTze Lim (liantze@gmail.com).
- Base [Rover Resume](https://github.com/subidit/rover-resume), originally created by [subidit](https://github.com/subidit/)

### Template Modifications

**(Update per AltaCV v1.6.3)**

- For the altacv class, I did not use the `pdfx` and `biblatex` package, as it caused error when installed using Docker.
- Package `pdfx` somehow is needed in order to use `withhyper` option on the document. As the bug that causes `pdfx` package error still not resolved, I discard `withhyper` option from the document class for this release version.
- Package `trimclip` somehow is missing after the recent class update (v1.6.3), so it gets ignored for my latest version.
- Package `accsupp` is now needed to generate PDF.

## Usage

```sh
cv-generator [command] [flag]
```
### Install Latex dependencies

```sh 
cv-generator install
```
### Run as webserver

```sh
cv-generator serve
```

Optional flags:

- `--port [PORT]`: specify port to run, default at `8170`

### Run as CLI app

```sh
cv-generator generate --input [JSON_INPUT_FILE]
```

Optional flags:

- `--output [OUTPUT_DIRECTORY]`: specify output directory, this command will store all files (pdf and LaTeX output) in the `[OUTPUT_DIRECTORY]/result` directory, default the output directory will be in the current working directory

## Installation

### Docker

Make sure to install Docker in your system

1. Pull the latest image from the registry by running `docker pull ghcr.io/thomasoca/cv-generator:latest`
2. To run the image as web server, bind the port, i.e. on port 8170 and run `docker run -p 8170:8170 ghcr.io/thomasoca/cv-generator serve`
3. To run the image as local file generator, simply run `docker run -v [LOCAL_DIR]:[OUTPUT_DIRECTORY] ghcr.io/thomasoca/cv-generator generate --input [INPUT_FILE] --output [OUTPUT_FILE]`

### Local Installation
#### Installation from source 
1. Install go >= 1.16
2. Clone this repository or download the compressed file in the [release](https://github.com/thomasoca/cv-generator/releases) section
3. [Compile and install](https://go.dev/doc/tutorial/compile-install) the application
4. Run the install command `cv-generator install`
5. Run `export PATH=$PATH:/[YOUR_HOME_DIR]/bin` to make sure that `pdflatex` and `tlmgr` is executable
6. Run the CLI app in webserver mode or local mode
