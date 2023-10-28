import React, { createContext, useContext, useReducer, useEffect } from "react";
import PropTypes from "prop-types";
import axios from "axios";
import { authReducer, SET_USER, LOGOUT } from "./authReducer";

const AuthContext = createContext();
export function useAuth() {
  return useContext(AuthContext);
}

export const AuthProvider = ({ children }) => {
  const storedUser = JSON.parse(localStorage.getItem("user"));
  const initialState = {
    user: {
      username: storedUser ? storedUser.username : null,
    },
    isAuthenticated: Boolean(storedUser),
    isLoading: false,
    token: storedUser ? storedUser.token : null,
  };

  const [state, dispatch] = useReducer(authReducer, initialState);

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      dispatch({ type: SET_USER, payload: JSON.parse(storedUser) });
    }
  }, []);

  const loginWithGoogle = (userData) => {
    return userData;
  };

  const loginWithEmail = async (login, password) => {
    try {
      const userLoginData = {
        login: login,
        password: password,
      };
      const response = await axios.post("/auth/login", userLoginData);

      console.log("Login successful:", response.data);

      const userData = {
        username: response.data.data.UserId,
        token: response.data.data.Jwt,
      };

      localStorage.setItem("user", JSON.stringify(userData));
      dispatch({ type: SET_USER, payload: userData });
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  const login = (userData, authMethod) => {
    if (authMethod === "GOOGLE") {
      loginWithGoogle(userData);
    } else if (authMethod === "EMAIL") {
      loginWithEmail(userData.login, userData.password);
    }
  };

  const logout = () => {
    localStorage.removeItem("user");
    dispatch({ type: LOGOUT });
  };

  const contextValue = {
    ...state,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={contextValue}>
      {state.isLoading ? <p>Loading...</p> : children}
    </AuthContext.Provider>
  );
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
