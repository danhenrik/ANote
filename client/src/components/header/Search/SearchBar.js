import React from "react";
import SearchIcon from "@mui/icons-material/Search";
import { Search, SearchIconWrapper, StyledInputBase } from "./SearchBar.styled";

const SearchBar = () => {
  return (
    <Search>
      <SearchIconWrapper>
        <SearchIcon />
      </SearchIconWrapper>
      <StyledInputBase
        style={{ width: "100%" }}
        placeholder='Pesquisar...'
        inputProps={{ "aria-label": "search" }}
        name='search'
        id='search'
      />
    </Search>
  );
};

export default SearchBar;
