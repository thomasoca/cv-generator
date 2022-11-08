import { Fragment, useState, useEffect } from "react";
import { JsonForms } from "@jsonforms/react";
import Grid from "@mui/material/Grid";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Alert from "@mui/material/Alert";
import AlertTitle from "@mui/material/AlertTitle";
import "./App.css";
import schema from "./schema.json";
import uischema from "./uischema.json";
import initial from "./initial.json";
import {
  materialCells,
  materialRenderers,
} from "@jsonforms/material-renderers";
import Loading from "./Loading";

const renderers = [
  ...materialRenderers,
  //register custom renderers
  //{ tester: ratingControlTester, renderer: RatingControl },
];

const App = () => {
  const savedData = localStorage.getItem("user");
  const initialData = !savedData ? initial : JSON.parse(savedData);
  const [displayDataAsString, setDisplayDataAsString] = useState("");
  const [jsonformsData, setJsonformsData] = useState<any>(initialData);
  const [loading, setLoading] = useState(false);
  const [visible, setVisible] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const onDismiss = () => setVisible(false);

  useEffect(() => {
    localStorage.setItem("user", JSON.stringify(jsonformsData, null, 2));
    setDisplayDataAsString(JSON.stringify(jsonformsData, null, 2));
  }, [jsonformsData]);

  const clearData = () => {
    setJsonformsData({});
  };

  const downloadObject = () => {
    setLoading(true);
    fetch("/api/v1/generate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: displayDataAsString,
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((text) => {
            throw new Error(text.message);
          });
        }
        return response.blob();
      })
      .then((data) => {
        var a = document.createElement("a");
        a.href = window.URL.createObjectURL(data);
        a.download = `${jsonformsData.personal_info.name}` + " Resume" + ".pdf";
        a.click();
        setLoading(false);
      })
      .catch((err): void => {
        setLoading(false);
        setVisible(true);
        setErrorMessage(err.message);
      });
  };

  return (
    <Fragment>
      <Loading loading={loading} />
      {visible ? (
        <Alert severity="error" onClose={onDismiss}>
          <AlertTitle>Error</AlertTitle>
          {errorMessage} —{" "}
          <strong>please check your input or try again later</strong>
        </Alert>
      ) : (
        <></>
      )}
      <div className="App">
        <header className="App-header">
          <h1 className="App-text-header">CV-Generator</h1>
        </header>
      </div>

      <Grid
        container
        justifyContent={"center"}
        spacing={1}
        className="container"
      >
        <Grid item sm={10}>
          <Typography
            variant={"h5"}
            sx={{
              textAlign: "center",
              padding: "0.25em",
            }}
          >
            Fill the form to make your resume
          </Typography>
          <div className="demoform">
            <JsonForms
              schema={schema}
              uischema={uischema}
              data={jsonformsData}
              renderers={renderers}
              cells={materialCells}
              onChange={({ errors, data }) => setJsonformsData(data)}
            />
          </div>
        </Grid>
        <Grid item sm={10}>
          <Button
            className="resetbutton"
            onClick={downloadObject}
            color="primary"
            variant="contained"
          >
            Submit data & get your resume
          </Button>
          &nbsp;
          <Button
            className="resetbutton"
            onClick={clearData}
            color="primary"
            variant="contained"
          >
            Clear data
          </Button>
        </Grid>
      </Grid>
    </Fragment>
  );
};

export default App;
