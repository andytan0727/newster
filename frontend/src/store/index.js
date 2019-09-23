import Vue from "vue";
import Vuex from "vuex";
import mutations from "./mutations";
import actions from "./actions";

Vue.use(Vuex);

const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== "production",
  state: {
    news: {
      // set an empty array to data instead of undefined to trigger Vue change mechanism
      csdn: {
        data: [],
      },
      "css-tricks": {
        data: [],
      },
    },
    loading: false,
    fetchError: false,
  },
  getters: {
    newsData: state => newsType => {
      return state.news[newsType].data;
    },
    loading: state => state.loading,
  },
  mutations,
  actions,
});

if (module.hot) {
  // accept actions and mutations as hot modules
  module.hot.accept(["./mutations"], () => {
    // require the updated modules
    // have to add .default here due to babel 6 module output
    const newMutations = require("./mutations").default;

    // swap in the new modules and mutations
    store.hotUpdate({
      mutations: newMutations,
    });
  });
}

export default store;
