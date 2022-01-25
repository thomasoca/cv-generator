import { Fragment, useState, useEffect } from "react";
import { JsonForms } from "@jsonforms/react";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import logo from "./logo.svg";
import "./App.css";
import schema from "./schema.json";
import uischema from "./uischema.json";
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

const initialData = {
  personal_info: {
    name: "Jane Doe",
    headline: "Software Engineer",
    picture:
      "https://trimelive.com/wp-content/uploads/2020/12/Gambar-wa-12.jpg",
    email: "jane.doe@email.com",
    github: "https://github.com/",
    linkedin: "Jane Doe",
    twitter: "",
    location_1: "Mansion",
    location_2: "City state",
  },
};

const renderers = [
  ...materialRenderers,
  //register custom renderers
  //{ tester: ratingControlTester, renderer: RatingControl },
];

const App = () => {
  const classes = useStyles();
  const [displayDataAsString, setDisplayDataAsString] = useState("");
  const [jsonformsData, setJsonformsData] = useState<any>(initialData);

  useEffect(() => {
    setDisplayDataAsString(JSON.stringify(jsonformsData, null, 2));
  }, [jsonformsData]);

  const clearData = () => {
    setJsonformsData({});
  };

  return (
    <Fragment>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to JSON Forms with React</h1>
          <p className="App-intro">More Forms. Less Code.</p>
        </header>
      </div>

      <Grid
        container
        justify={"center"}
        spacing={1}
        className={classes.container}
      >
        <Grid item sm={10}>
          <Typography variant={"h3"} className={classes.title}>
            Rendered form
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
          <Typography variant={"h3"} className={classes.title}>
            Bound data
          </Typography>

          <div className={classes.dataContent}>
            <pre id="boundData">{displayDataAsString}</pre>
          </div>
          <Button
            className={classes.resetButton}
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
