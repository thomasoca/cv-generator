import { Fragment, useState, useEffect } from "react";
import { JsonForms } from "@jsonforms/react";
import Grid from "@mui/material/Grid";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Alert from "@mui/material/Alert";
import AlertTitle from "@mui/material/AlertTitle";
import "./App.css";
import Header from "./Header";
import Footer from "./Footer";
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
  const [errorMessage, setErrorMessage] = useState<string | undefined>("");
  const onDismiss = () => setVisible(false);
  const [error, setError] = useState<string | undefined>("");
  useEffect(() => {
    localStorage.setItem("user", JSON.stringify(jsonformsData, null, 2));
    setDisplayDataAsString(JSON.stringify(jsonformsData, null, 2));
  }, [jsonformsData]);

  const clearData = () => {
    setJsonformsData({});
  };

  const downloadObject = async (): Promise<void> => {
    if (typeof error === "undefined" || error.trim().length === 0) {
      setVisible(false);
      setLoading(true);
      try {
        const response = await fetch("/api/v1/generate", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: displayDataAsString,
        });
        if (!response.ok) {
          const text = await response.json();
          throw new Error(text.message);
        }
        const data = await response.blob();
        const a = document.createElement("a");
        a.href = window.URL.createObjectURL(data);
        a.download = `${jsonformsData.personal_info.name} Resume.pdf`;
        a.click();
        setLoading(false);
      } catch (err: any) {
        setLoading(false);
        setVisible(true);
        setErrorMessage(err.message);
      }
    } else {
      setVisible(true);
      setErrorMessage(error);
    }
  };

  return (
    <Fragment>
      <div className="App">
        <Header />
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
                onChange={({ errors, data }) => {
                  setJsonformsData(data);
                  setError(
                    errors?.map((err) => err.message)[errors.length - 1]
                  );
                }}
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
        <Footer />
      </div>
    </Fragment>
  );
};

export default App;
