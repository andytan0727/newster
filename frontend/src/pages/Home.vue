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
import NewsTitleList from "@/components/NewsTitleList";
import LoadingSpinner from "@/components/LoadingSpinner";

const port = process.env.PORT || process.env.VUE_APP_API_PORT;

// request over https if hosting on cloud, else http if locally
const protocol = location.protocol;
const hostname = location.hostname;
const url =
  process.env.NODE_ENV === "production"
    ? `${window.location.origin}/api/csdn`
    : `${protocol}//${hostname}:${port}/api/csdn`;

export default {
  name: "home",
  components: {
    NewsTitleList,
    LoadingSpinner,
  },
  data() {
    return {
      loading: false,
      fetchError: false,
      news: [],
    };
  },
  mounted() {
    // spin up loader
    this.loading = true;

    fetch(url)
      .then(res => {
        return res.json();
      })
      .then(data => {
        this.news = data.news;
      })
      .catch(err => {
        console.log("Error fetching news: ", err);
        this.fetchError = true;
      })
      .finally(() => {
        // remove loader either promise resolved or rejected
        this.loading = false;
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