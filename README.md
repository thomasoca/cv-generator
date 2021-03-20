# CV-Generator
A slightly modified [AltaCV](https://github.com/liantze/AltaCV) generator ReST API written in go. The original altacv latex class was written by LianTze Lim (liantze@gmail.com). The goal of this API is to simplify the creation of beautiful CV with Latex by using simple HTTP request and harnessing the power of [Go template](https://golang.org/pkg/text/template/). Currently only serving the AltaCV class as the CV template.

## Modifications
1. I did not use the `pdfx` package, as it caused error when installed using Docker,
2. The whole `biblatex` package is commented out.

## How to run locally (using Docker)
1. Clone this repo using `git clone https://github.com/thomasoca/cv-generator.git`
2. Install [Docker](https://docs.docker.com/get-docker/) on your local machine
3. Change to the repo directory, and build the image:

    ```docker build -t [IMAGE_NAME] .```
4. Run the container, e.g:

    ```docker run -d -p 8080:8080 [IMAGE_NAME]```
5. The example of the JSON body of the request can be seen on the file `examples/user.json`. Use the JSON as the request body and perform a POST request to the API endpoint `localhost:8080/user`

## Run directly using terminal
You also can run the API directly using the terminal using `go run ./`, but make sure that the whole AltaCV Latex dependencies are installed. You can use the `Dockerfile` as a reference for installing the correct dependencies.

## To-Dos

- [ ] API documentation
- [ ] Directly using AltaCV [github](https://github.com/liantze/AltaCV) to define Latex class in Docker
- [ ] Dynamic CV section
- [ ] User can select color schema
- [ ] Live demo (frontend & backend)
- [ ] Anything else
