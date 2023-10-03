import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";

const ButtonBox = styled(Box)(({ theme }) => ({
  [theme.breakpoints.up("sm")]: {
    display: "none",
  },
}));

export default ButtonBox;
