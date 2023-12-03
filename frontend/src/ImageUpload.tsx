import React, { FC } from "react";
import PhotoCamera from "@mui/icons-material/PhotoCamera";
import "./ImageUpload.css";

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
  return (
    <>
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
      <label htmlFor={`icon_${path}`} className="icon-button-photo">
        <PhotoCamera />
      </label>
    </>
  );
};

export default ImageUpload;
