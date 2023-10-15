import { Button } from "@mui/material";
import { styled } from "@mui/material/styles";

const LoginButton = styled(Button)({
  width: "130px",
  backgroundColor: "white",
  color: "black",
  whiteSpace: "nowrap",
});

const SignupButton = styled(Button)(({ theme }) => ({
  width: "130px",
  backgroundColor: theme.palette.primary.main,
  color: "white",
  whiteSpace: "nowrap",
}));

export { LoginButton, SignupButton };
