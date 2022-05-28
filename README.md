# cv-generator

A web app to generate beautiful LaTeX resume using available open source templates (currently only serve AltaCV) by filling a simple form.

### Tech stack

- go backend to serve REST API to run templated latex file, the whole backend codebase can be seen on the [backend](./backend) directory
- Dockerized tinytex for latex compiler
- React SPA generated using create-react-app, the whole frontend codebase can be seen on the [frontend](./frontend) directory

## Live website

https://cv-generator-tex.herokuapp.com/

## Run it locally

Make sure to install Docker in your system

1. Navigate to the project roots directory, and build the Docker image `docker build -t [TAG_NAME] .`
2. Run the image and bind the port, i.e. on port 8080 `docker run -p 8080:8080 [TAG_NAME]`
3. Navigate to `localhost:8080` or any other ports that defined in the previous step
