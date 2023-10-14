import { styled } from "@mui/material/styles";
import MuiAppBar from "@mui/material/AppBar";
import { IconButton, Stack } from "@mui/material";
const drawerWidth = 240;

const AppBar = styled(MuiAppBar)`
  background-color: black;
  transition: ${({ theme }) =>
    theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    })};
  ${({ open, theme }) =>
    open &&
    `
    width: calc(100% - ${drawerWidth}px);
    margin-left: ${drawerWidth}px;
    transition: ${theme.transitions.create(["margin", "width"], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    })};
  `}
`;

const DrawerHeader = styled("div")`
  display: flex;
  align-items: center;
  padding: ${({ theme }) => theme.spacing(0, 1)};
  ${({ theme }) => theme.mixins.toolbar};
  justify-content: flex-end;
`;

const StyledIconButton = styled(IconButton)`
  color: orange;
  padding-right: 30px;
`;

const ButtonStack = styled(Stack)`
  ${({ theme }) => theme.breakpoints.down("sm")} {
    display: none;
  }
`;

export { AppBar, DrawerHeader, StyledIconButton, ButtonStack };
