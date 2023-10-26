const SET_USER = "SET_USER";
const LOGOUT = "LOGOUT";

const authReducer = (state, action) => {
  switch (action.type) {
    case SET_USER:
      return {
        ...state,
        user: {
          username: action.payload.username,
          email: action.payload.email,
        },
        isAuthenticated: true,
        isLoading: false,
        token: action.payload.token,
      };
    case LOGOUT:
      return {
        ...state,
        user: null,
        isAuthenticated: false,
        isLoading: false,
        token: null,
      };
    default:
      return state;
  }
};

export { authReducer, SET_USER, LOGOUT };
