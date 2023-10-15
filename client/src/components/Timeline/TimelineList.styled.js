import { Button } from "@mui/material";
import { styled } from "@mui/material/styles";

const GridWrapper = styled("div")(() => ({
  display: "flex",
  justifyContent: "center",
  width: "90%",
  margin: "0 auto",
}));

const CreateButton = styled(Button)(({ theme }) => ({
  backgroundColor: theme.palette.primary.main,
  color: "white",
  whiteSpace: "nowrap",
  "&:hover": {
    backgroundColor: theme.palette.primary.hover,
  },
}));

export { GridWrapper, CreateButton };
