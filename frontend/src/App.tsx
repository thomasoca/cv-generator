import { Fragment, useState, useEffect } from "react";
import { JsonForms } from "@jsonforms/react";
import Grid from "@mui/material/Grid";
import Button from "@mui/material/Button";
import { makeStyles } from "@material-ui/core/styles";
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
import { Document, Page, pdfjs } from "react-pdf";
import {
  KeyboardArrowLeft,
  KeyboardArrowRight,
  GetApp,
} from "@material-ui/icons";

pdfjs.GlobalWorkerOptions.workerSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.worker.js`;
const renderers = [
  ...materialRenderers,
  //register custom renderers
  //{ tester: ratingControlTester, renderer: RatingControl },
];
const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    marginTop: theme.spacing(2),
  },
  pdfContainer: {
    width: "100%",
    height: "100%",
    position: "relative",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    backgroundColor: "#f5f5f5",
  },
  navigationContainer: {
    width: "100%",
    alignItems: "center",
    display: "flex",
    padding: theme.spacing(2),
    zIndex: 1,
  },
  buttonContainer: {
    display: "flex",
    justifyContent: "center",
    marginBottom: theme.spacing(2),
  },
}));
const App = () => {
  const classes = useStyles();
  const savedData = localStorage.getItem("user");
  const initialData = !savedData ? initial : JSON.parse(savedData);
  const [displayDataAsString, setDisplayDataAsString] = useState("");
  const [jsonformsData, setJsonformsData] = useState<any>(initialData);
  const [loading, setLoading] = useState(false);
  const [visible, setVisible] = useState(false);
  const [errorMessage, setErrorMessage] = useState<string | undefined>("");
  const onDismiss = () => setVisible(false);
  const [error, setError] = useState<string | undefined>("");
  const [pdfBlob, setPdfBlob] = useState<Blob | null>(null);
  const [numPages, setNumPages] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
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
        setPdfBlob(data);
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

  const onDocumentLoadSuccess = ({ numPages }: { numPages: number }) => {
    setNumPages(numPages);
  };

  const goToPrevPage = () => {
    setCurrentPage((prevPage) => Math.max(prevPage - 1, 1));
  };

  const goToNextPage = () => {
    setCurrentPage((prevPage) => Math.min(prevPage + 1, numPages));
  };

  const handleDownload = () => {
    const link = document.createElement("a");
    if (pdfBlob != null) {
      link.href = window.URL.createObjectURL(pdfBlob);
      link.download = `${jsonformsData.personal_info.name} Resume.pdf`;
      link.click();
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
            <div className={classes.root}>
              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <div>
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
                  <Grid item sm={10}>
                    <Button
                      className="resetbutton"
                      onClick={downloadObject}
                      color="primary"
                      variant="contained"
                    >
                      Create your resume
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
                <Grid item xs={6}>
                  <div className={classes.pdfContainer}>
                    <div className={classes.navigationContainer}>
                      <Button
                        variant="text"
                        color="primary"
                        className="resetbutton"
                        onClick={goToPrevPage}
                        disabled={currentPage === 1 || !pdfBlob}
                        startIcon={<KeyboardArrowLeft />}
                      />
                      &nbsp;
                      <Button
                        variant="text"
                        color="primary"
                        className="resetbutton"
                        onClick={goToNextPage}
                        disabled={currentPage === numPages || !pdfBlob}
                        endIcon={<KeyboardArrowRight />}
                      />
                      &nbsp;
                      <Button
                        variant="text"
                        color="primary"
                        className="resetbutton"
                        onClick={handleDownload}
                        disabled={!pdfBlob}
                        startIcon={<GetApp />}
                      >
                        PDF
                      </Button>
                    </div>
                    {!pdfBlob && <div>Your resume will be shown here</div>}
                    {pdfBlob && (
                      <div>
                        {loading && <div>Loading...</div>}
                        {error && <div>Error loading PDF!</div>}
                        <Document
                          file={pdfBlob}
                          onLoadSuccess={onDocumentLoadSuccess}
                          renderMode="svg"
                        >
                          <Page
                            pageNumber={currentPage}
                            renderAnnotationLayer={false}
                            renderTextLayer={false}
                          />
                        </Document>
                        &nbsp;
                      </div>
                    )}
                  </div>
                </Grid>
              </Grid>
            </div>
          </Grid>
        </Grid>
        &nbsp;
        <Footer />
      </div>
    </Fragment>
  );
};

export default App;
