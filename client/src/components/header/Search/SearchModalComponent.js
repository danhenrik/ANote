import PropTypes from "prop-types";
import React from "react";
import Modal from "@mui/material/Modal";
import SearchForm from "./SearchForm";

const style = {
  display: "flex",
  alignItems: "flex-start",
  justifyContent: "center",
  overflow: "auto",
  maxHeight: "100vh",
  "@media (min-height: 700px)": {
    alignItems: "center",
  },
};

function SearchModalComponent({ open, handleClose }) {
  return (
    <div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby='modal-modal-search'
        aria-describedby='modal-modal-search'
        sx={style}
      >
        <>
          <SearchForm></SearchForm>
        </>
      </Modal>
    </div>
  );
}

SearchModalComponent.propTypes = {
  open: PropTypes.bool.isRequired,
  handleClose: PropTypes.func.isRequired,
};

export default SearchModalComponent;
