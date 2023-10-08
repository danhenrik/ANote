import PropTypes from "prop-types";
import React from "react";
import Modal from "@mui/material/Modal";
import LoginForm from "./login/LoginForm";
import SignupForm from "./Signup/SignupForm";

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

function AuthModalComponent({ open, handleClose, authType }) {
  return (
    <div>
      <Modal
        sx={style}
        open={open}
        onClose={handleClose}
        aria-labelledby='modal-modal-auth'
        aria-describedby='modal-modal-signup-login'
      >
        <>
          {authType === "Signup" ? (
            <SignupForm closeModal={handleClose} />
          ) : (
            <LoginForm closeModal={handleClose} />
          )}
        </>
      </Modal>
    </div>
  );
}

AuthModalComponent.propTypes = {
  open: PropTypes.bool.isRequired,
  handleClose: PropTypes.func.isRequired,
  authType: PropTypes.string.isRequired,
};

export default AuthModalComponent;
