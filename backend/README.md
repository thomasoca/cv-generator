# CV-Generator

A slightly modified [AltaCV](https://github.com/liantze/AltaCV) generator ReST API written in go. The original altacv latex class was written by LianTze Lim (liantze@gmail.com). The goal of this API is to simplify the creation of beautiful CV with Latex by using simple HTTP request and harnessing the power of [Go template](https://golang.org/pkg/text/template/).

## How it works

I used [Go template](https://golang.org/pkg/text/template/) to write Latex file from a JSON input. Then using subprocess and [TinyTex](https://yihui.org/tinytex/) to compile the Latex output and PDF file. All of the latex output are temporarily stored at `/tmp` folder using randomized folder and file name, then it will be deleted at the end of the request, whether the request is success or failed.

## Modifications

**(Update per AltaCV v1.6.3)**

- For the altacv class, I did not use the `pdfx` and `biblatex` package, as it caused error when installed using Docker.
- Package `pdfx` somehow is needed in order to use `withhyper` option on the document. As the bug that causes `pdfx` package error still not resolved, I discard `withhyper` option from the document class for this release version.
- Package `trimclip` somehow is missing after the recent class update (v1.6.3), so it gets ignored for my latest version.
- Package `accsupp` is now needed to generate PDF.

## How to run locally (using Docker)

1. Clone this repo using `git clone https://github.com/thomasoca/cv-generator.git`
2. Install [Docker](https://docs.docker.com/get-docker/) on your local machine
3. Change to the repo directory, and build the image:

   ```
   docker build -t [IMAGE_NAME] .
   ```

4. Run the container, e.g:

   ```
   docker run -d -p 8080:8080 [IMAGE_NAME]
   ```

5. The example of the JSON body of the request can be seen on the file [examples/user.json](/examples/user.json). Use the JSON as the request body and perform a POST request to the API endpoint `localhost:8080/user`
6. Important `ENV` variables:
   - `ENV_MODE` determine the environment that you are using. Set it to `PRD` for production environment (the latex output will be temporarily stored on random /tmp folder and deleted after the request is done)
   - `PROJECT_DIR` determine the working directory

## Run directly using terminal

You also can run the API directly using the terminal using `go run ./` for debugging purpose, but make sure that the whole AltaCV Latex dependencies are installed. You can use the [Dockerfile](./Dockerfile) as a reference for installing the correct dependencies.

## API documentation

The available endpoint for cv-generator is described below

#### Get the example JSON to generate a CV

- **URL**
  `/api/v1/example`
- **Method:**

  `GET`

- **Response**
  - `200` `OK`
    - Will return [example file](examples/user.json).
  - `405` `Method not allowed`
    - `{"message": "Method not allowed"}`
  - `500` `Server Error`
    - `{"message": "Failed processing file"}`

#### Generate PDF file of your CV

- **URL:**

  `/api/v1/user`

- **Method:**

  `POST`

- **Data params:**

  ```json
  {
    "personal_info": {
      "name": "[full name: string; required]",
      "headline": "[CV headline: string; required]",
      "picture": "[public URL or base64: string]",
      "email": "[email address: string]",
      "phone": "[phone number: string]",
      "github": "[github profile URL: string]",
      "linkedin": "[LinkedIn name/URL: string]",
      "twitter": "[Twitter account: string]",
      "location_1": "[Address 1 (street/building): string]",
      "location_2": "[Address 2 (city/state): string]"
    },
    "main_section": {
      "about_me": {
        "label": "About Me",
        "descriptions": "[Short description: string; required]"
      },
      "work_experience": {
        "label": "Experience",
        "lists": [
          {
            "company": "[Name of the company: string]",
            "position": "[Position level: string]",
            "start_period": "[Starting work date: string]",
            "end_period": "[End work date: string]",
            "location": "[company location: string]",
            "descriptions": ["list of short description: string"]
          }
        ]
      },
      "education": {
        "label": "Education",
        "lists": [
          {
            "institution": "[Name of the institution: string; required]",
            "major": "[Major taken: string; required]",
            "level": "[Degree obtained: string; required]",
            "gpa": "[GPA: string]",
            "start_period": "[Starting school date: string]",
            "end_period": "[End school date: string]",
            "location": "[institution location: string]",
            "descriptions": ["list of short description: string"]
          }
        ]
      },
      "extracurricular": {
        "label": "Extra-curricular Activities",
        "lists": [
          {
            "institution": "[Name of the institution: string]",
            "position": "[Position level: string]",
            "start_period": "[Starting work date: string]",
            "end_period": "[End work date: string]",
            "location": "[company location: string]",
            "descriptions": ["list of short description: string"]
          }
        ]
      },
      "skills": {
        "label": "Skills",
        "descriptions": ["list of skills: string"]
      },
      "projects": {
        "label": "Projects",
        "lists": [
          {
            "title": "[Name/title of the project: string]",
            "link": "[Public URL of the project: string]",
            "start_period": "[Starting work date: string]",
            "end_period": "[End work date: string]",
            "descriptions": "[Project description: string]"
          }
        ]
      },
      "languages": {
        "label": "Languages",
        "descriptions": [
          {
            "language": "[language name: string]",
            "fluency": "[fluency level: string]"
          }
        ]
      }
    }
  }
  ```

- **Request data example**

  See [example file](examples/user.json).

- **Response**
  - `200` `OK`
    - Will return a PDF (content type `application/pdf`) of your CV
  - `400` `Bad Request`
    - `{"message": "Bad request"}`
  - `405` `Method not allowed`
    - `{"message": "Method not allowed"}`
  - `500` `Server Error`
    - If the error occurred on Latex/PDF generation process: `{"message": "Failed creating file"}`
    - If the error occurred on the file serving process: `{"message": "Failed processing file"}`
    - If the error occurred on the file uploading process: `{"message": "Failed sending file"}`
