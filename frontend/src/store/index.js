import Vue from "vue";
import Vuex from "vuex";
import mutations from "./mutations";
import actions from "./actions";

Vue.use(Vuex);

const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== "production",
  state: {
    // set an empty array to news data instead of undefined to trigger
    // Vue change mechanism
    news: {
      csdn: {
        logoUrl: "https://csdnimg.cn/cdn/content-toolbar/csdnlogo.png",
        logoAlt: "CSDN Logo",
        desc: "CSDN News",
        newsLink: "/news/csdn",
        data: [],
      },
      "css-tricks": {
        logoUrl:
          "https://css-tricks.com/wp-content/uploads/2013/06/CSS-Tricks-logo.png",
        logoAlt: "CSS-Tricks Logo",
        desc: "CSS-Tricks News",
        newsLink: "/news/css-tricks",
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
    newsKey: state => Object.keys(state.news),

    // get home news panel data without data fetched
    homeNewsPanelData: state =>
      Object.keys(state.news).map(k => ({ ...state.news[k], data: undefined })),
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
