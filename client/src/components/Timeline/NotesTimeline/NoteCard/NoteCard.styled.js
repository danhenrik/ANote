import { styled } from "@mui/material/styles";
import { Avatar, Card, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import CommentIcon from "@mui/icons-material/Comment";

const AvatarBackground = styled("div")(({ randomColor }) => ({
  backgroundColor: randomColor,
  width: "100%",
  display: "flex",
  alignItems: "center",
}));

const Title = styled(Typography)(() => ({
  textAlign: "left",
  fontWeight: "bold",
}));

const AvatarAuthor = styled(Typography)(() => ({
  color: "black",
  marginLeft: "10px",
  fontWeight: "bold",
}));

const AvatarContainer = styled("div")({
  height: "60px",
});

const AvatarUsernames = styled("div")({
  alignItems: "center",
});

const CustomAvatar = styled(Avatar)(() => ({
  backgroundColor: "black",
  width: "60px",
  height: "60px",
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

const CommentContainer = styled("div")({
  maxHeight: "200px",
  overflowY: "auto",
  marginTop: "10px",
  marginBottom: "10px",
  padding: "3px",

  "&::-webkit-scrollbar": {
    width: "6px", // Width of the entire scrollbar
  },
  "&::-webkit-scrollbar-thumb": {
    background: "#888", // Color of the scrollbar handle
  },

  "&::-webkit-scrollbar-track": {
    background: "#f1f1f1", // Color of the scrollbar track
  },
});

const NotesCardContainer = styled(Card)({
  "&:hover": {
    backgroundColor: "#f0f0f0",
    transform: "scale(1.02)",
  },
});

const ModalStyling = {
  display: "flex",
  alignItems: "flex-start",
  justifyContent: "center",
  overflow: "auto",
  maxHeight: "100vh",
  "@media (min-height: 400px)": {
    alignItems: "center",
  },
};

const CommentButton = styled(CommentIcon)(() => ({
  marginLeft: "10px",
  color: "blue",
  "&:hover": {
    transform: "scale(1.2)",
  },
}));

export {
  AvatarBackground,
  Title,
  AvatarAuthor,
  AvatarContainer,
  AvatarUsernames,
  CustomAvatar,
  StyledLink,
  ContentContainer,
  CommentContainer,
  NotesCardContainer,
  ModalStyling,
  CommentButton,
};
