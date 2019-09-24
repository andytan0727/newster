<template>
  <div class="home-news-panel-div">
    <img class="news-logo" :src="logoSrc" :alt="logoAlt" />
    <router-link :to="newsLink" class="news-desc">{{ desc }}</router-link>
  </div>
</template>

<script>
export default {
  name: "home-news-panel",
  props: {
    desc: {
      type: String,
      required: true,
    },
    logoSrc: {
      type: String,
      required: true,
    },
    logoAlt: {
      type: String,
      required: true,
    },
    newsLink: {
      type: String,
      required: true,
    },
  },
};
</script>

<style lang="scss" scoped>
@import "../_variables.scss";

.home-news-panel-div {
  width: 75%;
  display: flex;
  align-items: center;
  box-shadow: $box-shadow;
  padding: 1.5rem;

  @media screen and (min-width: $tablet-width) {
    width: 85%;
    padding: 2rem;
  }

  .news-logo {
    margin-right: 1.2rem;
    height: 30px;
    width: 110px;

    @media screen and (min-width: $tablet-width) {
      height: 50px;
      width: 200px;
    }
  }

  .news-desc {
    font-weight: bold;
    text-align: left;

    @media screen and (min-width: $tablet-width) {
      font-size: 1.4rem;
    }
  }
}
</style>