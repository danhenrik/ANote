import styled from "@emotion/styled";
import { Avatar, Typography } from "@mui/material";
import { Link } from "react-router-dom";

export const Title = styled(Typography)(() => ({
  textAlign: "left",
  fontWeight: "bold",
}));

export const AvatarAuthor = styled(Typography)(() => ({
  color: "black",
  marginLeft: "10px",
  fontWeight: "bold",
}));

export const AvatarBackground = styled("div")(({ randomColor }) => ({
  backgroundColor: randomColor(),
  width: "100%",
  display: "flex",
  aligItems: "center",
}));

export const AvatarContainer = styled("div")(() => ({
  height: "60px",
}));

export const AvatarUsernames = styled("div")`
  alignitems: "center";
`;

export const CustomAvatar = styled(Avatar)(() => ({
  backgroundColor: "black",
  width: "60px",
  height: "60px",
}));

export const StyledLink = styled(Link)`
  text-decoration: none;
  color: inherit;
  &:hover {
    text-decoration: underline; /* Apply underline on hover */
  }
`;
