import PropTypes from "prop-types";
import React from "react";
import Modal from "@mui/material/Modal";
import { Container, IconButton, Paper } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";

const style = {
  display: "flex",
  alignItems: "flex-start",
  justifyContent: "center",
  overflow: "auto",
  maxHeight: "100vh",
  "@media (min-height: 400px)": {
    alignItems: "center",
  },
};

function ModalComponent({ handleClose, open, content }) {
  return (
    <Modal
      sx={style}
      onClose={handleClose}
      open={open}
      aria-labelledby='modal-modal-auth'
      aria-describedby='modal-modal-signup-login'
    >
      <Container maxWidth='xs'>
        <Paper
          elevation={3}
          sx={{
            padding: 2,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <IconButton onClick={handleClose} sx={{ alignSelf: "flex-end" }}>
            <CloseIcon />
          </IconButton>
          {content}
        </Paper>
      </Container>
    </Modal>
  );
}

ModalComponent.propTypes = {
  handleClose: PropTypes.func.isRequired,
  content: PropTypes.node,
  open: PropTypes.bool.isRequired,
};

export default ModalComponent;
