<template>
  <div class="home-div">
    <h1>Categories:</h1>
    <home-news-panel
      v-for="data in newsPanelData"
      :key="data.title"
      :logoSrc="data.logoUrl"
      :logoAlt="data.logoAlt"
      :desc="data.desc"
      :newsLink="data.newsLink"
    ></home-news-panel>
  </div>
</template>

<script>
import HomeNewsPanel from "@/components/HomeNewsPanel";

export default {
  name: "home",
  components: {
    HomeNewsPanel,
  },
  computed: {
    newsPanelData() {
      return this.$store.getters.homeNewsPanelData;
    },
  },
};
</script>

<style lang="scss" scoped>
$marginTop: 1.5rem;

.home-div {
  margin-top: $marginTop;
  min-height: 75vh;
  display: flex;
  flex-direction: column;
  align-items: center;

  h1 {
    align-self: flex-start;
    margin: 0;
    margin-left: 3rem;
  }

  div {
    margin-top: $marginTop;
  }
}
</style>