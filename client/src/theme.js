// theme.js
import { createTheme } from "@mui/material/styles";

// Define your custom theme
const theme = createTheme({
  palette: {
    primary: {
      main: "#7F56D9",
    },
    background: {
      main: "black",
    },
  },
  typography: {
    fontFamily: ["Arial"],
    fontWeight: 600,
  },
});

export default theme;
