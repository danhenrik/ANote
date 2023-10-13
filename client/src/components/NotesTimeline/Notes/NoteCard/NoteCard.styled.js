import styled from "@emotion/styled";
import { Avatar, Card, Typography } from "@mui/material";
import { Link } from "react-router-dom";

export const AvatarBackground = styled("div")(({ randomColor }) => ({
  backgroundColor: randomColor({
    luminosity: "light",
    format: "rgb",
  }),
  width: "100%",
  display: "flex",
  aligItems: "center",
}));

export const Title = styled(Typography)`
  text-align: left;
  font-weight: bold;
`;

export const AvatarAuthor = styled(Typography)`
  color: black;
  margin-left: 10px;
  font-weight: bold;
`;

export const AvatarContainer = styled("div")`
  height: 60px;
`;

export const AvatarUsernames = styled("div")`
  align-items: center;
`;

export const CustomAvatar = styled(Avatar)`
  background-color: black;
  width: 60px;
  height: 60px;
`;

export const StyledLink = styled(Link)`
  text-decoration: none;
  color: inherit;
  &:hover {
    text-decoration: underline; /* Apply underline on hover */
  }
`;

export const ContentContainer = styled("div")`
  min-height: 9em;
  max-height: 9em;
  overflow-y: hidden;

  word-wrap: break-word;
  line-height: 1.8em;
  text-align: left;
  text-overflow: ellipsis;
  p {
    color: black;
  }
`;

export const NotesCard = styled(Card)`
  &:hover {
    background-color: #f0f0f0;
    transform: scale(1.02);
  }
`;
