// Manual Jest mock for axios so tests can run without parsing the ESM
// distribution shipped by axios v1. CRA enables `resetMocks: true`, so we
// can't rely on jest.fn(impl) keeping its implementation across tests.
// Wrap mock functions with `mockImplementation` accessors that stay valid.

const buildInstance = (config = {}) => ({
  defaults: {
    baseURL: config.baseURL,
    withCredentials: Boolean(config.withCredentials),
    headers: {
      common: {},
      ...(config.headers ? { ...config.headers } : {}),
    },
  },
  get: jest.fn(),
  post: jest.fn(),
  put: jest.fn(),
  patch: jest.fn(),
  delete: jest.fn(),
});

const create = jest.fn(buildInstance);
create.mockImplementation(buildInstance);

const axios = {
  create,
  get: jest.fn(),
  post: jest.fn(),
  put: jest.fn(),
  patch: jest.fn(),
  delete: jest.fn(),
  defaults: { headers: { common: {} } },
};

// Re-attach the implementation so that resetMocks: true (CRA default) does
// not strip it on the first beforeEach.
beforeEach(() => {
  create.mockImplementation(buildInstance);
});

module.exports = axios;
module.exports.default = axios;
