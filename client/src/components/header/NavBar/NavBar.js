import { useTheme } from "@mui/material/styles";
import Drawer from "@mui/material/Drawer";
import Toolbar from "@mui/material/Toolbar";
import Groups2RoundedIcon from "@mui/icons-material/Groups2Rounded";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import FilterAltRoundedIcon from "@mui/icons-material/FilterAltRounded";
import SearchBar from "../Search/SearchBar";
import {
  DrawerHeader,
  AppBar,
  StyledIconButton,
  ButtonStack,
} from "./NavBar.styled";
import NavList from "./NavList";
import { PropTypes } from "prop-types";
import { useEffect, useState } from "react";
import SearchModalComponent from "../Search/SearchModalComponent";
import { Link } from "react-router-dom";
import NavButtons from "./NavButtons";
import AnoteLogo from "../Logo/AnoteLogo";
import { useAuth } from "../../../store/auth-context";
import { useModal } from "../../../store/modal-context";
import LoginForm from "../../AccessControl/Login/LoginForm";
import SignupForm from "../../AccessControl/Signup/SignupForm";
import { Avatar } from "@mui/material";

const NavBar = ({ open, setOpen }) => {
  const drawerWidth = 240;
  const theme = useTheme();
  const auth = useAuth();
  const [openSearch, setOpenSearch] = useState(false);
  const [avatarPreview, setAvatarPreview] = useState(auth.avatar);
  const modal = useModal();

  const handleLoginModal = () => {
    modal.openModal(<LoginForm closeModal={modal.closeModal}></LoginForm>);
  };

  const handleSignupModal = () => {
    modal.openModal(<SignupForm closeModal={modal.closeModal}></SignupForm>);
  };

  const handleDrawer = () => {
    setOpen((open) => !open);
  };

  const handleOpenSearch = () => setOpenSearch(true);

  const handleCloseSearch = () => setOpenSearch(false);

  useEffect(() => {
    setAvatarPreview(auth.avatar);
  }, [auth.avatar]);

  return (
    <>
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
          <Link to='/timeline' style={{ textDecoration: "none" }}>
            <AnoteLogo />
          </Link>
          <SearchBar />
          <IconButton onClick={handleOpenSearch} sx={{ color: "white" }}>
            <FilterAltRoundedIcon />
          </IconButton>
          <Link to='/communities' style={{ textDecoration: "none" }}>
            <StyledIconButton>
              <Groups2RoundedIcon />
            </StyledIconButton>
          </Link>
          {!auth.isAuthenticated ? (
            <ButtonStack direction='row' spacing={3}>
              <NavButtons
                handleLoginModal={handleLoginModal}
                handleSignupModal={handleSignupModal}
              />
            </ButtonStack>
          ) : (
            <Avatar
              sx={{
                width: "50px",
                height: "50px",
                marginLeft: "auto",
                marginRight: "80px",
              }}
              alt='Avatar'
              src={avatarPreview}
            ></Avatar>
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
            handleDrawer={handleDrawer}
            avatarPreview={avatarPreview}
          />
        </Drawer>
      )}
    </>
  );
};

NavBar.propTypes = {
  open: PropTypes.bool.isRequired,
  setOpen: PropTypes.func.isRequired,
};

export default NavBar;
