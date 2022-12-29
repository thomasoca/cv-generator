# CV-Generator API

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
