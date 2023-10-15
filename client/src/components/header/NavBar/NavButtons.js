import PropTypes from "prop-types";
import { LoginButton, SignupButton } from "./NavButtons.styled";

const NavButtons = ({ handleLoginModal, handleSignupModal }) => {
  return (
    <>
      <LoginButton onClick={handleLoginModal} variant='contained'>
        Login
      </LoginButton>
      <SignupButton onClick={handleSignupModal} variant='contained'>
        Cadastre-se
      </SignupButton>
    </>
  );
};

NavButtons.propTypes = {
  handleLoginModal: PropTypes.func.isRequired,
  handleSignupModal: PropTypes.func.isRequired,
};

export default NavButtons;
