import React, { FC } from "react";
import { Grid, Button } from "@material-ui/core";
import CloudUploadIcon from "@material-ui/icons/CloudUpload";
import DeleteIcon from "@material-ui/icons/Delete";

interface ImageUploadProps {
  path: string;
  updateValue: (newValue: string) => void;
  value?: string;
}

const getBase64 = (file: File) => {
  return new Promise<string>((resolve) => {
    let baseURL: string = "";
    // Make new FileReader
    let reader = new FileReader();

    // Convert the file to base64 text
    reader.readAsDataURL(file);

    // on reader load somthing...
    reader.onload = () => {
      // Make a fileInfo Object
      baseURL = reader.result as string;
      resolve(baseURL);
    };
  });
};

const ImageUpload: FC<ImageUploadProps> = ({ path, updateValue, value }) => {
  const handleClear = () => {
    updateValue("");
  };
  return (
    <div style={{ marginBottom: "20px" }}>
      <Grid container spacing={2} alignItems="center">
        <Grid item>
          <input
            accept="image/*"
            style={{ display: "none" }}
            id={`icon_${path}`}
            onChange={(e) => {
              let file = e?.target.files?.item(0)!;
              getBase64(file)
                .then((result) => {
                  updateValue(result);
                })
                .catch((err) => {
                  console.log(err);
                });
            }}
            type="file"
            multiple={false} // Change to single image upload
          />
          <label htmlFor={`icon_${path}`}>
            <Button
              variant="contained"
              color="primary"
              component="span"
              startIcon={<CloudUploadIcon />}
            >
              Upload Photo
            </Button>
          </label>
        </Grid>
        &nbsp;
        <Grid item>
          <Button
            onClick={handleClear}
            className="materialBtn"
            startIcon={<DeleteIcon />}
          >
            Clear
          </Button>
        </Grid>
        &nbsp;
      </Grid>
    </div>
  );
};

export default ImageUpload;
