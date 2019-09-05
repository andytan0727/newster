<template>
  <div class="main-content">
    <h1 class="top-title">Top News:</h1>
    <div class="spinner" v-if="loading">
      <loading-spinner></loading-spinner>
    </div>
    <news-title-list v-else :news="news"></news-title-list>
  </div>
</template>

<script>
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
      loading: false,
      news: [],
    };
  },
  mounted() {
    // spin up loader
    this.loading = true;

    fetch("http://localhost:8000/api/csdn")
      .then(res => {
        return res.json();
      })
      .then(data => {
        this.news = data.news;
      })
      .catch(err => {
        console.log("Error fetching news: ", err);
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

.main-content {
  padding: 0 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>