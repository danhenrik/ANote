import { useTheme } from "@mui/material/styles";
import Drawer from "@mui/material/Drawer";
import Toolbar from "@mui/material/Toolbar";
import Groups2RoundedIcon from "@mui/icons-material/Groups2Rounded";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import FilterAltRoundedIcon from "@mui/icons-material/FilterAltRounded";
import SearchBar from "./Search/SearchBar";
import { DrawerHeader, AppBar, StyledIconButton } from "./NavBar.style";
import NavList from "./NavList";
import AuthModalComponent from "../AccessControl/login/AuthModalComponent";
import { Button, Stack, SvgIcon } from "@mui/material";
import { PropTypes } from "prop-types";
import { useState } from "react";
import { ReactComponent as LogoIcon } from "./logo.svg";
import SearchModalComponent from "./Search/SearchModalComponent";
import { Link } from "react-router-dom";

function NavBar({ open, setOpen }) {
  const drawerWidth = 240;
  const theme = useTheme();

  const [openAuth, setOpenAuth] = useState(false);
  const [openSearch, setOpenSearch] = useState(false);
  const [authType, setAuthType] = useState("");

  const handleOpenLogin = () => {
    setOpenAuth(true);
    setAuthType("Login");
  };

  const handleOpenSearch = () => setOpenSearch(true);

  const handleCloseSearch = () => setOpenSearch(false);

  const handleOpenSignup = () => {
    setOpenAuth(true);
    setAuthType("Signup");
  };

  const handleCloseLogin = () => setOpenAuth(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <>
      <AuthModalComponent
        open={openAuth}
        handleClose={handleCloseLogin}
        authType={authType}
      />
      <SearchModalComponent open={openSearch} handleClose={handleCloseSearch} />
      <AppBar position='fixed' open={open}>
        <Toolbar>
          <IconButton
            color='inherit'
            aria-label='open drawer'
            onClick={handleDrawerOpen}
            edge='start'
            sx={{ mr: 2, ...(open && { display: "none" }) }}
          >
            <MenuIcon />
          </IconButton>
          <Link to='/' style={{ textDecoration: "none" }}>
            <div style={{ display: "flex" }}>
              <SvgIcon
                sx={{ mr: "10px" }}
                component={LogoIcon}
                inheritViewBox
              />
              <Typography
                variant='h6'
                component='span'
                sx={{
                  color: "purple",
                  fontWeight: "bold",
                }}
              >
                ANote
              </Typography>
            </div>
          </Link>
          <SearchBar />
          <IconButton onClick={handleOpenSearch} sx={{ color: "white" }}>
            <FilterAltRoundedIcon />
          </IconButton>
          <StyledIconButton>
            <Groups2RoundedIcon />
          </StyledIconButton>
          <Stack direction='row' spacing={3}>
            <Button
              onClick={handleOpenLogin}
              variant='contained'
              sx={{
                width: "130px",
                backgroundColor: "white",
                color: "black",
                whiteSpace: "nowrap",
              }}
            >
              Login
            </Button>
            <Button
              onClick={handleOpenSignup}
              variant='contained'
              sx={{
                width: "130px",
                backgroundColor: "#31CEFF",
                color: "white",
                whiteSpace: "nowrap",
              }}
            >
              Cadastre-se
            </Button>
          </Stack>
        </Toolbar>
      </AppBar>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: drawerWidth,
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
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === "ltr" ? (
              <ChevronLeftIcon sx={{ color: "white" }} />
            ) : (
              <ChevronRightIcon />
            )}
          </IconButton>
        </DrawerHeader>
        <NavList />
      </Drawer>
    </>
  );
}

NavBar.propTypes = {
  open: PropTypes.bool.isRequired,
  setOpen: PropTypes.func.isRequired,
};

export default NavBar;
