import React from "react";
import { act, render, waitFor } from "@testing-library/react";
import axios from "axios";
import { AuthProvider, useAuth } from "./auth-context";

jest.mock("axios");

const ConsumerRef = React.createRef();

const Consumer = () => {
  const auth = useAuth();
  ConsumerRef.current = auth;
  return null;
};

const renderProvider = () =>
  render(
    <AuthProvider>
      <Consumer />
    </AuthProvider>
  );

describe("AuthProvider", () => {
  beforeEach(() => {
    localStorage.clear();
    jest.clearAllMocks();
    ConsumerRef.current = null;
  });

  it("does not persist the JWT token in localStorage after login", async () => {
    axios.post.mockResolvedValueOnce({
      data: {
        data: {
          User: { Id: "alice", Avatar: null },
          Jwt: "super-secret-jwt",
        },
      },
    });

    renderProvider();

    await act(async () => {
      await ConsumerRef.current.login(
        { login: "alice@example.com", password: "x" },
        "EMAIL"
      );
    });

    const stored = JSON.parse(localStorage.getItem("user"));
    expect(stored).not.toBeNull();
    expect(stored.username).toBe("alice");
    expect(stored).not.toHaveProperty("token");
    expect(JSON.stringify(stored)).not.toContain("super-secret-jwt");
  });

  it("sends withCredentials when calling the login endpoint", async () => {
    axios.post.mockResolvedValueOnce({
      data: {
        data: {
          User: { Id: "alice", Avatar: null },
          Jwt: "x",
        },
      },
    });

    renderProvider();

    await act(async () => {
      await ConsumerRef.current.login(
        { login: "alice@example.com", password: "x" },
        "EMAIL"
      );
    });

    expect(axios.post).toHaveBeenCalledWith(
      "/api/auth/login",
      expect.objectContaining({ login: "alice@example.com" }),
      expect.objectContaining({ withCredentials: true })
    );
  });

  it("calls the logout endpoint and clears local storage on logout", async () => {
    localStorage.setItem(
      "user",
      JSON.stringify({ username: "alice", avatar: null })
    );
    axios.post.mockResolvedValueOnce({ data: {} });

    renderProvider();

    await waitFor(() => expect(ConsumerRef.current).not.toBeNull());

    await act(async () => {
      await ConsumerRef.current.logout();
    });

    expect(axios.post).toHaveBeenCalledWith(
      "/api/auth/logout",
      null,
      expect.objectContaining({ withCredentials: true })
    );
    expect(localStorage.getItem("user")).toBeNull();
  });

  it("clears local state even when the logout request fails", async () => {
    localStorage.setItem(
      "user",
      JSON.stringify({ username: "alice", avatar: null })
    );
    axios.post.mockRejectedValueOnce(new Error("network"));

    renderProvider();

    await waitFor(() => expect(ConsumerRef.current).not.toBeNull());

    await act(async () => {
      await ConsumerRef.current.logout();
    });

    expect(localStorage.getItem("user")).toBeNull();
  });
});
