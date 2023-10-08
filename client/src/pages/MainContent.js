import Main from "./MainContent.style";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import { PropTypes } from "prop-types";

function MainContent({ open, children }) {
  return (
    <Box sx={{ display: "flex" }}>
      <CssBaseline />
      <Main sx={{ marginTop: "50px" }} open={open}>
        {children}
      </Main>
    </Box>
  );
}
MainContent.propTypes = {
  open: PropTypes.bool.isRequired,
  children: PropTypes.object,
};

export default MainContent;
