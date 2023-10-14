import {
  Divider,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Avatar,
} from "@mui/material";
import NavButtons from "./NavButtons";
import PropTypes from "prop-types";
import { ButtonBox, ListLink } from "./NavList.styled";
import { useAuth } from "../../store/auth-context";

function NavList({ handleLoginModal, handleSignupModal }) {
  const auth = useAuth();

  let listOptions = [
    "Minhas Notas",
    "Notas Públicas",
    "Minhas Comunidades",
    "Comunidades Populares",
    "Amigos",
  ];

  auth.isAuthenticated && listOptions.push("Logout");

  return (
    <>
      <div
        style={{ display: "flex", justifyContent: "center", padding: "10px" }}
      >
        <Avatar
          sx={{ bgcolor: "orange", width: "100px", height: "100px" }}
          alt='Remy Sharp'
          src='/logo.svg'
        >
          B
        </Avatar>
      </div>
      <Divider sx={{ backgroundColor: "white" }} />
      <List>
        {listOptions.map((text) => (
          <ListLink key={text} to={`/${text}`}>
            <ListItem disablePadding>
              <ListItemButton>
                <ListItemIcon></ListItemIcon>
                <ListItemText primary={text} />
              </ListItemButton>
            </ListItem>
          </ListLink>
        ))}
        <ButtonBox textAlign='center'>
          <NavButtons
            sx={{ padding: "2px", display: "flex" }}
            handleLoginModal={handleLoginModal}
            handleSignupModal={handleSignupModal}
          />
        </ButtonBox>
      </List>
    </>
  );
}
NavList.propTypes = {
  handleLoginModal: PropTypes.func.isRequired,
  handleSignupModal: PropTypes.func.isRequired,
};

export default NavList;
