import { SET_NEWS, SET_LOADING } from "./mutation-types";

const mutations = {
  [SET_NEWS](state, payload) {
    const { newsType, news } = payload;
    state.news[newsType].data = news;
  },
  [SET_LOADING](state, loading) {
    state.loading = loading;
  },
};

export default mutations;
