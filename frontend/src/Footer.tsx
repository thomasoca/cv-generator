import React from "react";
import { Typography, Link, Box } from "@mui/material";

export default function Footer() {
  return (
    <Box
      component="footer"
      sx={{ mt: "auto", py: 2, backgroundColor: "#f5f5f5" }}
    >
      <Typography variant="body2" color="textSecondary" align="center">
        Found a bug or problem? Please report it on our{" "}
        <Link
          href="https://github.com/thomasoca/cv-generator/issues"
          color="inherit"
        >
          issue tracker
        </Link>
        .
      </Typography>
    </Box>
  );
}
