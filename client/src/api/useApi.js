import axios from "axios";
import { useMemo } from "react";

const useApi = () => {
  const api = useMemo(
    () =>
      axios.create({
        baseURL: "/api/",
        withCredentials: true,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    []
  );

  return api;
};

export default useApi;
