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
import { useAuth } from "../../../store/auth-context";

function NavList({ handleLoginModal, handleSignupModal, handleDrawer }) {
  const auth = useAuth();

  const handleLogout = () => {
    handleDrawer();
    auth.logout();
  };
  let listOptions = [
    { text: "Minhas Notas", route: { path: "/timeline" } },
    {
      text: "Notas Públicas",
      route: { path: "/timeline", queryParams: "world=true" },
    },
    { text: "Minhas Comunidades", route: { path: "/communities" } },
    {
      text: "Comunidades Populares",
      route: { path: "/communities", queryParams: "world=true" },
    },
    { text: "Amigos", route: { path: "/friends" } },
    { text: "Logout", action: () => handleLogout() },
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
        {listOptions.map((option, idx) => (
          <ListLink
            key={idx}
            onClick={option.action}
            to={
              option.route && {
                pathname: option.route.path,
                search: option.route.queryParams,
              }
            }
          >
            <ListItem disablePadding>
              <ListItemButton>
                <ListItemIcon></ListItemIcon>
                <ListItemText primary={option.text} />
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
  handleDrawer: PropTypes.func.isRequired,
};

export default NavList;
