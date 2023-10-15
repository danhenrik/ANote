import { styled } from "@mui/material/styles";
import { Card, Typography } from "@mui/material";
import { Link } from "react-router-dom";

const Title = styled(Typography)(() => ({
  textAlign: "left",
  fontWeight: "bold",
}));

const StyledLink = styled(Link)(() => ({
  textDecoration: "none",
  color: "inherit",
  "&:hover": {
    textDecoration: "underline",
  },
}));

const ContentContainer = styled("div")({
  minHeight: "9em",
  maxHeight: "9em",
  overflowY: "hidden",
  wordWrap: "break-word",
  lineHeight: "1.8em",
  textAlign: "left",
  textOverflow: "ellipsis",
  p: {
    color: "black",
  },
});

const CommunityCardContainer = styled(Card)({
  "&:hover": {
    backgroundColor: "#f0f0f0",
    transform: "scale(1.02)",
  },
});

export { Title, StyledLink, ContentContainer, CommunityCardContainer };
