import PropTypes from "prop-types";
import {
  FacebookLoginButton,
  GoogleLoginButton,
} from "react-social-login-buttons";

const SocialMediaAuth = ({ authType, googleHandler, facebookHandler }) => {
  const loginText = authType === "Login" ? "Login" : "Cadastre-se";
  return (
    <>
      <GoogleLoginButton
        style={{ height: "40px", borderRadius: "6px" }}
        onClick={googleHandler}
      >
        <span>{loginText} com Google</span>
      </GoogleLoginButton>
      <FacebookLoginButton
        style={{ height: "40px", borderRadius: "6px" }}
        onClick={facebookHandler}
      >
        <span>{loginText} com Facebook</span>
      </FacebookLoginButton>
    </>
  );
};
SocialMediaAuth.propTypes = {
  authType: PropTypes.string.isRequired,
  googleHandler: PropTypes.func.isRequired,
  facebookHandler: PropTypes.func.isRequired,
};
export default SocialMediaAuth;
