# cv-generator

A web app and CLI tool to generate beautiful LaTeX resume using available open source templates by filling a simple form (or using a JSON file in CLI mode).

## Live website

https://cv-generator-40m5.onrender.com

## Available templates

- A slightly modified [AltaCV](https://github.com/liantze/AltaCV). The original altacv latex class was written by LianTze Lim (liantze@gmail.com).

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

1. Navigate to the project roots directory, and build the Docker image `docker build -t [TAG_NAME] .`
2. Run the image and bind the port, i.e. on port 8080 `docker run -p 8170:8170 [TAG_NAME]`
3. Navigate to `localhost:8170` or any other ports that defined in the previous step

### Local Installation

1. Install go >= 1.16,
2. Run the LaTeX [installation script](./backend/scripts/setup_latex.sh)
3. Run `export PATH=$PATH:/[YOUR_HOME_DIR]/bin` to make sure that `pdflatex` is executable
4. Run the CLI app in webserver mode or local mode
