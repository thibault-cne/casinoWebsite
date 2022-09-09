const mutations = {
  updateStorage(state, credentials) {
    state.accessToken = credentials.accessToken;
    state.refreshToken = credentials.refreshToken;
  },
  destroyToken(state) {
    state.accessToken = null;
    state.refreshToken = null;
  },
  setAccessToken(state, token) {
    state.accessToken = token;
  },
};

export { mutations };
