import PropTypes from "prop-types";
import { useGoogleLogin } from "@react-oauth/google";
import FacebookLogin from "@greatsumini/react-facebook-login";
import {
  FacebookLoginButton,
  GoogleLoginButton,
} from "react-social-login-buttons";

import { useAuth } from "../../../store/auth-context";

const SocialMediaAuth = ({ authType }) => {
  const auth = useAuth();

  const loginText = authType === "Login" ? "Login" : "Cadastre-se";

  const googleHandler = useGoogleLogin({
    onSuccess: (codeResponse) => auth.login(codeResponse, "GOOGLE"),
    onError: () => alert("error"),
  });

  return (
    <>
      <GoogleLoginButton onClick={googleHandler}>
        {loginText + " com Google"}
      </GoogleLoginButton>
      <FacebookLogin
        appId=''
        buttonText={loginText + " com Facebook"}
        render={({ onClick, logout }) => (
          <FacebookLoginButton onClick={onClick} onLogoutClick={logout}>
            {loginText + " com Facebook"}
          </FacebookLoginButton>
        )}
      />
    </>
  );
};
SocialMediaAuth.propTypes = {
  authType: PropTypes.string.isRequired,
};
export default SocialMediaAuth;
