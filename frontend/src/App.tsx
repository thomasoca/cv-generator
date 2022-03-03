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
import { makeStyles } from "@mui/styles";
import Loading from "./Loading";

const useStyles = makeStyles((_theme) => ({
  container: {
    padding: "1em",
    width: "100%",
  },
  title: {
    textAlign: "center",
    padding: "0.25em",
  },
  dataContent: {
    display: "flex",
    justifyContent: "center",
    borderRadius: "0.25em",
    backgroundColor: "#cecece",
    marginBottom: "1rem",
  },
  resetButton: {
    margin: "auto",
    display: "block",
  },
  demoform: {
    margin: "auto",
    padding: "1rem",
  },
}));

const renderers = [
  ...materialRenderers,
  //register custom renderers
  //{ tester: ratingControlTester, renderer: RatingControl },
];

const App = () => {
  const classes = useStyles();
  const [displayDataAsString, setDisplayDataAsString] = useState("");
  const [jsonformsData, setJsonformsData] = useState<any>(initial);
  const [loading, setLoading] = useState(false);
  const [visible, setVisible] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const onDismiss = () => setVisible(false);

  useEffect(() => {
    setDisplayDataAsString(JSON.stringify(jsonformsData, null, 2));
  }, [jsonformsData]);

  const clearData = () => {
    setJsonformsData({});
  };

  const downloadObject = () => {
    setLoading(true);
    fetch("/api/generate", {
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
        a.download = "file.pdf";
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
          <img src="./image-logo.png" alt="logo" className="App-logo" />
          <h1 className="App-title">
            A not so simple way to generate your resume
          </h1>
        </header>
      </div>

      <Grid
        container
        justifyContent={"center"}
        spacing={1}
        className={classes.container}
      >
        <Grid item sm={10}>
          <Typography variant={"h5"} className={classes.title}>
            Fill the form to make your resume
          </Typography>
          <div className={classes.demoform}>
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
        <Grid item sm={6}>
          <Button
            className={classes.resetButton}
            onClick={clearData}
            color="primary"
            variant="contained"
          >
            Clear data
          </Button>
          &nbsp;
          <Button
            className={classes.resetButton}
            onClick={downloadObject}
            color="primary"
            variant="contained"
          >
            Submit data & get your resume
          </Button>
        </Grid>
      </Grid>
    </Fragment>
  );
};

export default App;
