import { withJsonFormsControlProps } from "@jsonforms/react";
import ImageUpload from "./ImageUpload"; // Assuming you have an ImageUpload component
import { rankWith, scopeEndsWith, ControlProps } from "@jsonforms/core";

interface ImageUploadControlTesterProps {
  uischema: any;
  schema: any;
}

interface NewHandleChangeProps {
  handleChange(path: string, newValue: any): void;
  path: string;
  newValue: any;
}

interface ImageUploadControlProps {
  data: any;
  path: string;
  label: string;
  handleChange(path: string, value: any): void;
}

export const ImageUploadControlTester = rankWith(
  3, // increase rank as needed
  scopeEndsWith("picture")
);

const newHandleChange = ({
  handleChange,
  path,
  newValue,
}: NewHandleChangeProps) => {
  handleChange(path, newValue);
};

const ImageUploadControl = ({
  data,
  handleChange,
  path,
  label,
}: ControlProps) => (
  <ImageUpload
    updateValue={(newValue: string) =>
      newHandleChange({ handleChange, path, newValue })
    }
    value={data}
    path={path}
  />
);

export default withJsonFormsControlProps(ImageUploadControl);
