import { styled } from "@mui/material/styles";
import { Card, Typography } from "@mui/material";
import { Link } from "react-router-dom";

const Title = styled(Typography)(() => ({
  textAlign: "left",
  fontWeight: "bold",
}));

const StyledLink = styled("div")(() => ({
  textDecoration: "none",
  color: "inherit",
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

const FollowButton = styled(Link)({
  marginTop: "7px",
  backgroundColor: "#9370DB",
  color: "white",
  marginRight: "10px",
  padding: "12px 32px",
  borderRadius: "9999px",
  textDecoration: "none",
  display: "inline-block",
  marginLeft: "auto",
  "&:hover": {
    backgroundColor: "purple", // Change to purple on hover
  },
});

const FollowButtonFollowing = styled(Link)({
  marginTop: "7px",
  backgroundColor: "purple",
  color: "white",
  marginRight: "10px",
  padding: "12px 32px",
  borderRadius: "9999px",
  textDecoration: "none",
  display: "inline-block",
  marginLeft: "auto",
  "&:hover": {
    backgroundColor: "#9370DB", // Change to purple on hover
  },
});
export {
  Title,
  StyledLink,
  ContentContainer,
  CommunityCardContainer,
  FollowButton,
  FollowButtonFollowing,
};
