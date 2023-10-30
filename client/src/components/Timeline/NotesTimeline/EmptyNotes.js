import React from "react";
import PropTypes from "prop-types";
import { Button, Grid, Paper } from "@mui/material";

const EmptyNotes = ({ children, clickHandler }) => {
  return (
    <Grid
      container
      spacing={0}
      direction='column'
      alignItems='center'
      justifyContent='center'
      sx={{ minHeight: "70vh" }}
    >
      <Grid item>
        <Button
          onClick={clickHandler}
          sx={{ width: "100%", height: "100%", textAlign: "center" }}
        >
          <Paper
            elevation={3} // Adjust the elevation for the shadow effect
            sx={{
              padding: "30px",
              borderRadius: "10px",
              borderColor: "lightgray",
              borderWidth: "4px",
              borderStyle: "solid",
              backgroundColor: "#f5f5f5", // Lighter gray background
              color: "black",
              transition: "background-color 0.3s",
              "&:hover": {
                backgroundColor: "lightgray", // Lighter gray background on hover
              },
            }}
          >
            {children}
          </Paper>
        </Button>
      </Grid>
    </Grid>
  );
};

EmptyNotes.propTypes = {
  children: PropTypes.node.isRequired,
  clickHandler: PropTypes.func.isRequired,
};

export default EmptyNotes;
