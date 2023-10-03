import { SvgIcon, Typography } from "@mui/material";
import { ReactComponent as LogoIcon } from ".././logo.svg";

function AnoteLogo() {
  return (
    <div style={{ display: "flex" }}>
      <SvgIcon sx={{ mr: "10px" }} component={LogoIcon} inheritViewBox />
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
  );
}

export default AnoteLogo;
