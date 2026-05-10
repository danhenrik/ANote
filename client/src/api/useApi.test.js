import React from "react";
import { render } from "@testing-library/react";
import useApi from "./useApi";

const captureApi = (sink) => {
  const Probe = () => {
    sink.current = useApi();
    return null;
  };
  render(<Probe />);
};

describe("useApi", () => {
  it("returns an axios instance with withCredentials enabled", () => {
    const sink = { current: null };
    captureApi(sink);

    expect(sink.current).toBeDefined();
    expect(sink.current.defaults.withCredentials).toBe(true);
    expect(sink.current.defaults.baseURL).toBe("/api/");
  });

  it("does not share an axios instance across separate hook callers", () => {
    const first = { current: null };
    const second = { current: null };
    captureApi(first);
    captureApi(second);

    expect(first.current).not.toBe(second.current);
  });

  it("does not attach a default Authorization header", () => {
    const sink = { current: null };
    captureApi(sink);

    expect(
      sink.current.defaults.headers.common.Authorization
    ).toBeUndefined();
  });
});
