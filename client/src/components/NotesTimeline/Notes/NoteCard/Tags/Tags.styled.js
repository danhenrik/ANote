import { styled } from "@mui/material/styles";
import { Link } from "react-router-dom";

const TagList = styled("div")(() => ({
  display: "flex",
  flexWrap: "wrap",
}));

const TagLink = styled(Link)(({ theme }) => ({
  textDecoration: "none",
  backgroundColor: theme.palette.primary.main,
  color: "#fff",
  border: "none",
  borderRadius: "4px",
  margin: "4px",
  padding: "8px 18px",
  cursor: "pointer",

  "&:hover": {
    textDecoration: "underline",
  },
}));

export { TagLink, TagList };
