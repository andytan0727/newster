import { SET_NEWS, SET_LOADING } from "./mutation-types";

const actions = {
  fetchNews({ commit }, newsType) {
    const port = process.env.PORT || process.env.VUE_APP_API_PORT;

    // request over https if hosting on cloud, else http if locally
    const protocol = location.protocol;
    const hostname = location.hostname;
    const url =
      process.env.NODE_ENV === "production"
        ? `${window.location.origin}/api/${newsType}`
        : `${protocol}//${hostname}:${port}/api/${newsType}`;

    // spin up loader
    commit(SET_LOADING, true);

    return fetch(url)
      .then(res => res.json())
      .then(data => {
        commit(SET_NEWS, { news: data.news, newsType });
      })
      .finally(() => {
        // remove loader either promise resolved or rejected
        commit(SET_LOADING, false);
      });
  },
};

export default actions;
