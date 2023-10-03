import { styled } from "@mui/system";
import InputLabelOriginal from "@mui/material/InputLabel";
import ButtonOriginal from "@mui/material/Button";
import TextFieldOriginal from "@mui/material/TextField";

export const InputLabel = styled(InputLabelOriginal)`
  color: black;
  font-family: sans-serif;
  font-size: 0.9em;
  padding: 8px 0px 8px 0px;
`;

export const TextField = styled(TextFieldOriginal)`
  .MuiOutlinedInput-root {
    border-radius: 8px;
    height: 42px;
  }
`;

export const TextArea = styled(TextFieldOriginal)`
  .MuiOutlinedInput-root {
    border-radius: 8px;
    height: auto;
  }
`;

export const Button = styled(ButtonOriginal)`
  margin-top: 15px;
`;

export const LineSeparator = styled("hr")`
  border: none;
  border-top: 1px solid #ccc;
  margin-top: 20px;
  margin-bottom: 20px;
`;
