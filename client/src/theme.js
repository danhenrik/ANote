import { createTheme } from "@mui/material/styles";

const defaultTheme = createTheme({
  palette: {
    primary: {
      main: "#7F56D9",
      hover: "#593f94",
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

export default defaultTheme;
