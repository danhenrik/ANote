import PropTypes from "prop-types";
import React from "react";
import Modal from "@mui/material/Modal";
import { Container, IconButton, Paper } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";

const ModalComponent = ({ handleClose, open, content, style }) => {
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
};

ModalComponent.propTypes = {
  handleClose: PropTypes.func.isRequired,
  content: PropTypes.node,
  open: PropTypes.bool.isRequired,
  style: PropTypes.node,
};

export default ModalComponent;
