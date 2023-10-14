import { useTheme } from "@mui/material/styles";
import Drawer from "@mui/material/Drawer";
import Toolbar from "@mui/material/Toolbar";
import Groups2RoundedIcon from "@mui/icons-material/Groups2Rounded";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import FilterAltRoundedIcon from "@mui/icons-material/FilterAltRounded";
import SearchBar from "./Search/SearchBar";
import {
  DrawerHeader,
  AppBar,
  StyledIconButton,
  ButtonStack,
} from "./NavBar.styled";
import NavList from "./NavList";
import AuthModalComponent from "../AccessControl/AuthModalComponent";
import { PropTypes } from "prop-types";
import { useState } from "react";
import SearchModalComponent from "./Search/SearchModalComponent";
import { Link } from "react-router-dom";
import NavButtons from "./NavButtons";
import AnoteLogo from "./Logo/AnoteLogo";
import { useAuth } from "../../store/auth-context";

function NavBar({ open, setOpen }) {
  const drawerWidth = 240;
  const theme = useTheme();
  const auth = useAuth();
  const [openAuth, setOpenAuth] = useState(false);
  const [openSearch, setOpenSearch] = useState(false);
  const [authType, setAuthType] = useState("");

  const handleLoginModal = () => {
    setOpenAuth((open) => !open);
    setAuthType("Login");
  };

  const handleSignupModal = () => {
    setOpenAuth((open) => !open);
    setAuthType("Signup");
  };

  const handleDrawer = () => {
    setOpen((open) => !open);
  };

  const handleOpenSearch = () => setOpenSearch(true);

  const handleCloseSearch = () => setOpenSearch(false);

  const handleCloseAuth = () => setOpenAuth(false);

  return (
    <>
      <AuthModalComponent
        open={openAuth}
        handleClose={handleCloseAuth}
        authType={authType}
      />
      <SearchModalComponent open={openSearch} handleClose={handleCloseSearch} />

      <AppBar position='fixed' open={open}>
        <Toolbar>
          {auth.isAuthenticated && (
            <IconButton
              color='inherit'
              aria-label='open drawer'
              onClick={handleDrawer}
              edge='start'
              sx={{ mr: 2, ...(open && { display: "none" }) }}
            >
              <MenuIcon />
            </IconButton>
          )}
          <Link to='/' style={{ textDecoration: "none" }}>
            <AnoteLogo />
          </Link>
          <SearchBar />
          <IconButton onClick={handleOpenSearch} sx={{ color: "white" }}>
            <FilterAltRoundedIcon />
          </IconButton>
          <StyledIconButton>
            <Groups2RoundedIcon />
          </StyledIconButton>
          {!auth.isAuthenticated && (
            <ButtonStack direction='row' spacing={3}>
              <NavButtons
                handleLoginModal={handleLoginModal}
                handleSignupModal={handleSignupModal}
              />
            </ButtonStack>
          )}
        </Toolbar>
      </AppBar>
      {auth.isAuthenticated && (
        <Drawer
          sx={{
            width: drawerWidth,
            flexShrink: 0,
            "& .MuiDrawer-paper": {
              boxSizing: "border-box",
              backgroundColor: "black",
              color: "orange",
            },
          }}
          variant='persistent'
          anchor='left'
          open={open}
        >
          <DrawerHeader>
            <IconButton onClick={handleDrawer}>
              {theme.direction === "ltr" ? (
                <ChevronLeftIcon sx={{ color: "white" }} />
              ) : (
                <ChevronRightIcon />
              )}
            </IconButton>
          </DrawerHeader>
          <NavList
            handleLoginModal={handleLoginModal}
            handleSignupModal={handleSignupModal}
          />
        </Drawer>
      )}
    </>
  );
}

NavBar.propTypes = {
  open: PropTypes.bool.isRequired,
  setOpen: PropTypes.func.isRequired,
};

export default NavBar;
