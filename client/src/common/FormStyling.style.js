import { styled } from "@mui/system";
import InputLabelOriginal from "@mui/material/InputLabel";
import ButtonOriginal from "@mui/material/Button";
import TextFieldOriginal from "@mui/material/TextField";

const InputLabel = styled(InputLabelOriginal)`
  color: black;
  font-family: sans-serif;
  font-size: 0.9em;
  padding: 8px 0px 8px 0px;
`;

const TextField = styled(TextFieldOriginal)`
  .MuiOutlinedInput-root {
    border-radius: 8px;
    height: 42px;
  }
`;

const TextArea = styled(TextFieldOriginal)`
  .MuiOutlinedInput-root {
    border-radius: 8px;
    height: auto;
  }
`;

const Button = styled(ButtonOriginal)`
  margin-top: 15px;
`;

const LineSeparator = styled("hr")`
  border: none;
  border-top: 1px solid #ccc;
  margin-top: 20px;
  margin-bottom: 20px;
`;
export { InputLabel, Button, TextField, LineSeparator, TextArea };
