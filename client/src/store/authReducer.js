const initialState = {
  user: null,
  isAuthenticated: false,
  isLoading: false,
  token: null,
};

const SET_USER = "SET_USER";
const LOGOUT = "LOGOUT";

const authReducer = (state, action) => {
  switch (action.type) {
    case SET_USER:
      return {
        ...state,
        user: action.payload.user,
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

export { initialState, authReducer, SET_USER, LOGOUT };
