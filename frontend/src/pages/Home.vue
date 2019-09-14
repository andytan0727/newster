<template>
  <div class="main-content">
    <h1 class="top-title">Top News:</h1>
    <h2 class="fetch-error" v-if="fetchError">Error Fetching News...</h2>

    <div class="spinner" v-if="loading">
      <loading-spinner></loading-spinner>
    </div>
    <news-title-list v-else :news="news"></news-title-list>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import NewsTitleList from "@/components/NewsTitleList";
import LoadingSpinner from "@/components/LoadingSpinner";

export default {
  name: "home",
  components: {
    NewsTitleList,
    LoadingSpinner,
  },
  data() {
    return {
      fetchError: false,
    };
  },
  computed: {
    ...mapGetters(["loading"]),
    news() {
      return this.$store.getters.news("csdn");
    },
  },
  methods: {
    ...mapActions(["fetchNews"]),
  },
  mounted() {
    // fetch news on mount
    this.fetchNews("csdn").catch(err => {
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