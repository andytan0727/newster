<template>
  <div class="main-content">
    <h1 class="top-title">Top News:</h1>
    <h2 class="fetch-error" v-if="fetchError">Error Fetching News...</h2>

    <div class="spinner" v-if="loading">
      <loading-spinner></loading-spinner>
    </div>
    <news-title-list v-else :news="newsData"></news-title-list>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import NewsTitleList from "@/components/NewsTitleList";
import LoadingSpinner from "@/components/LoadingSpinner";

export default {
  name: "news",
  components: {
    NewsTitleList,
    LoadingSpinner,
  },
  props: {
    newsType: {
      type: String,
      required: true,
      default: "csdn", // TODO: remove if new home page is made later
    },
  },
  data() {
    return {
      fetchError: false,
    };
  },
  computed: {
    ...mapGetters(["loading"]),
    newsData() {
      return this.$store.getters.newsData(this.newsType);
    },
  },
  methods: {
    ...mapActions(["fetchNews"]),
  },
  beforeRouteUpdate(to, from, next) {
    const { newsType } = to.params;

    // fetch again if newsType param changed in /news
    this.fetchNews(newsType).catch(err => {
      console.log("Error fetching news: ", err);
      this.fetchError = true;
    });

    // remember to call next()
    next();
  },
  created() {
    this.fetchNews(this.newsType).catch(err => {
      console.log("Error fetching news: ", err);
      this.fetchError = true;
    });
  },
};
</script>

<style lang="scss" scoped>
.top-title {
  align-self: flex-start;
}

.spinner {
  height: 75vh;
  display: flex;
  align-items: center;
}

.fetch-error {
  height: 65vh;
  display: flex;
  align-items: center;
  color: red;
}

.main-content {
  padding: 0 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>