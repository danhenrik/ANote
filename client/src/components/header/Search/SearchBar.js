import React, { useState } from "react";
import SearchIcon from "@mui/icons-material/Search";
import { Search, SearchIconWrapper, StyledInputBase } from "./SearchBar.styled";
import { createSearchParams, useNavigate } from "react-router-dom";

const SearchBar = () => {
  const navigate = useNavigate();
  const [searchValue, setSearchValue] = useState("");

  const handleChange = (e) => {
    setSearchValue(e.target.value);
  };

  const handleSearch = (e) => {
    if (e.key === "Enter") {
      if (searchValue != "") {
        navigate({
          pathname: "/timeline",
          search: createSearchParams({
            search: "true",
            title: searchValue,
          }).toString(),
        });
      }
    }
  };
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
        onChange={handleChange}
        onKeyDown={handleSearch}
      />
    </Search>
  );
};

export default SearchBar;
