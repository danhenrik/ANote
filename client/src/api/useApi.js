import axios from "axios";

import { useAuth } from "../store/auth-context";

const api = axios.create({
  baseURL: "/",
  headers: {
    ContentType: "application/json",
    timeout: 1000,
  },
});

const useApi = () => {
  const auth = useAuth();

  if (auth.isLoggedIn) {
    api.defaults.headers.common.Authorization = `Bearer ${auth.token}`;
  } else {
    delete api.defaults.headers.common.Authorization;
  }

  return api;
};

export default useApi;
