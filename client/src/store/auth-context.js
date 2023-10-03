import { createContext, useContext, useReducer, useEffect } from "react";
import { PropTypes } from "prop-types";

const initialState = {
  user: null,
  isAuthenticated: false,
  isLoading: false,
};

const AuthContext = createContext();

const SET_USER = "SET_USER";
const LOGOUT = "LOGOUT";

const authReducer = (state, action) => {
  switch (action.type) {
    case SET_USER:
      return {
        ...state,
        user: action.payload,
        isAuthenticated: true,
        isLoading: false,
      };
    case LOGOUT:
      return {
        ...state,
        user: null,
        isAuthenticated: false,
        isLoading: false,
      };
    default:
      return state;
  }
};

export function useAuth() {
  return useContext(AuthContext);
}

export function AuthProvider({ children }) {
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

  // eslint-disable-next-line no-unused-vars
  const loginWithEmail = (email, password) => {
    return true;
  };

  const login = (userData, authMethod) => {
    if (authMethod == "GOOGLE") {
      loginWithGoogle(userData);
    } else if (authMethod == "EMAIL") {
      loginWithEmail(userData.email, userData.password);
    }

    localStorage.setItem("user", JSON.stringify(userData));
    dispatch({ type: SET_USER, payload: userData });
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
}

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
