import { Fragment, useState, useEffect } from "react";
import { JsonForms } from "@jsonforms/react";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import "./App.css";
import schema from "./schema.json";
import uischema from "./uischema.json";
import initial from "./initial.json";
import {
  materialCells,
  materialRenderers,
} from "@jsonforms/material-renderers";
import { makeStyles } from "@material-ui/core/styles";

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

  useEffect(() => {
    setDisplayDataAsString(JSON.stringify(jsonformsData, null, 2));
  }, [jsonformsData]);

  const clearData = () => {
    setJsonformsData({});
  };

  const downloadObject = () => {
    fetch("/api/generate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: displayDataAsString,
    })
      .then((response) => {
        if (response.ok) {
          return response.blob();
        }
        throw new Error(response.statusText);
      })
      .then((data) => {
        var a = document.createElement("a");
        a.href = window.URL.createObjectURL(data);
        a.download = "file.pdf";
        a.click();
      })
      .catch((err): void => {
        console.log(JSON.stringify(err.message, null, 2));
      });
  };

  return (
    <Fragment>
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
        justify={"center"}
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
