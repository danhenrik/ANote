import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
import { Link } from "react-router-dom";

const ButtonBox = styled(Box)(({ theme }) => ({
  [theme.breakpoints.up("sm")]: {
    display: "none",
  },
}));

const ListLink = styled(Link)(() => ({
  textDecoration: "none",
  color: "orange",
  "&:hover": {
    backgroundColor: "orange",
  },
}));

export { ButtonBox, ListLink };
