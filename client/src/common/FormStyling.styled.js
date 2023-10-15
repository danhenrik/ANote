import { styled } from "@mui/material/styles";
import InputLabelOriginal from "@mui/material/InputLabel";
import ButtonOriginal from "@mui/material/Button";
import TextFieldOriginal from "@mui/material/TextField";

const InputLabel = styled(InputLabelOriginal)(() => ({
  color: "black",
  fontFamily: "sans-serif",
  fontSize: "0.9em",
  padding: "8px 0px",
}));

const TextField = styled(TextFieldOriginal)(() => ({
  ".MuiOutlinedInput-root": {
    borderRadius: "8px",
    height: "42px",
  },
}));

const TextArea = styled(TextFieldOriginal)(() => ({
  ".MuiOutlinedInput-root": {
    borderRadius: "8px",
    height: "auto",
  },
}));

const Button = styled(ButtonOriginal)(() => ({
  marginTop: "15px",
}));

const LineSeparator = styled("hr")(() => ({
  border: "none",
  borderTop: "1px solid #ccc",
  marginTop: "20px",
  marginBottom: "20px",
}));

export { InputLabel, TextField, TextArea, Button, LineSeparator };
