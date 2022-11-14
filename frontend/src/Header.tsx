import { AppBar, Toolbar, Typography } from "@material-ui/core";
import React from "react";
import "./App.css";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export default function Header() {
  const displayDesktop = () => {
    return (
      <Toolbar className="App-toolbar">
        {appLogo}
        {menuBar}
      </Toolbar>
    );
  };

  const appLogo = (
    <Typography variant="h6" component="h1" className="App-text-header">
      CV-Generator
    </Typography>
  );

  const menuBar = (
    <div className="App-menubar">
      <a href="https://github.com/thomasoca/cv-generator">
        <FontAwesomeIcon icon={faGithub} color="#FFFEFE" size="2x" />
      </a>
    </div>
  );

  return (
    <header>
      <AppBar>{displayDesktop()}</AppBar>
    </header>
  );
}
