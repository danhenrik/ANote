import { styled, alpha, InputBase } from "@mui/material";

const Search = styled("div")`
  position: relative;
  border-radius: ${({ theme }) => theme.shape.borderRadius};
  background-color: ${({ theme }) => alpha(theme.palette.common.white, 0.15)};
  &:hover {
    background-color: ${({ theme }) => alpha(theme.palette.common.white, 0.25)};
  }
  margin-right: ${({ theme }) => theme.spacing(2)};
  margin-left: 0;
  width: 100%;
  ${({ theme }) => theme.breakpoints.up("sm")} {
    margin-left: ${({ theme }) => theme.spacing(3)};
    width: 70%;
  }
`;

const SearchIconWrapper = styled("div")`
  padding: ${({ theme }) => theme.spacing(0, 2)};
  height: 100%;
  position: absolute;
  pointer-events: none;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const StyledInputBase = styled(InputBase)`
  color: inherit;
  & .MuiInputBase-input {
    padding: ${({ theme }) => theme.spacing(1, 1, 1, 0)};
    padding-left: ${({ theme }) => `calc(1em + ${theme.spacing(4)})`};
    transition: ${({ theme }) => theme.transitions.create("width")};
    width: 100%;
    ${({ theme }) => theme.breakpoints.up("md")} {
      width: 100%;
    }
  }
`;

export { Search, SearchIconWrapper, StyledInputBase };
