import { Select, TextareaAutosize } from "@mui/material";
import { styled } from "@mui/material/styles";

const CustomTextArea = styled(TextareaAutosize)({
  width: "100%",
  border: "1px solid #ced4da",
  borderRadius: "4px",
  padding: "10px",
  fontSize: "16px",
  font: "inherit",
  "&::placeholder": {
    color: "lightgray",
  },
});

const CustomSelect = styled(Select)({
  height: "40px",
});

export { CustomTextArea, CustomSelect };
