import PropTypes from "prop-types";

import { Button } from "@mui/material";

function NavButtons({ handleLoginModal, handleSignupModal }) {
  return (
    <>
      <Button
        onClick={handleLoginModal}
        variant='contained'
        sx={{
          width: "130px",
          backgroundColor: "white",
          color: "black",
          whiteSpace: "nowrap",
        }}
      >
        Login
      </Button>
      <Button
        onClick={handleSignupModal}
        variant='contained'
        sx={{
          width: "130px",
          backgroundColor: "#31CEFF",
          color: "white",
          whiteSpace: "nowrap",
        }}
      >
        Cadastre-se
      </Button>
    </>
  );
}

NavButtons.propTypes = {
  handleLoginModal: PropTypes.func.isRequired,
  handleSignupModal: PropTypes.func.isRequired,
};

export default NavButtons;
