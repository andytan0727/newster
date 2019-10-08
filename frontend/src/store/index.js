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
      "dev-to": {
        logoUrl:
          "https://res.cloudinary.com/practicaldev/image/fetch/s--7-9vabgh--/c_fill,f_auto,fl_progressive,h_320,q_auto,w_320/https://thepracticaldev.s3.amazonaws.com/uploads/user/profile_image/3/13d3b32a-d381-4549-b95e-ec665768ce8f.png",
        logoAlt: "DevTo Logo",
        desc: "DevTo News",
        newsLink: "/news/dev-to",
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

    // get newsLinks with its title name to be used in navbar
    newsLinks: state =>
      Object.keys(state.news).map(key => ({
        name: key,
        link: state.news[key].newsLink,
      })),

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
