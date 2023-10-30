const SET_USER = "SET_USER";
const LOGOUT = "LOGOUT";

const authReducer = (state, action) => {
  switch (action.type) {
    case SET_USER:
      return {
        ...state,
        username: action.payload.username,
        email: action.payload.email,
        isAuthenticated: true,
        isLoading: false,
        token: action.payload.token,
        avatar: action.payload.avatar,
      };
    case LOGOUT:
      return {
        ...state,
        username: null,
        email: null,
        isAuthenticated: false,
        isLoading: false,
        token: null,
        avatar: null,
      };
    default:
      return state;
  }
};

export { authReducer, SET_USER, LOGOUT };
