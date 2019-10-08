<template>
  <nav class="navbar-news-panel">
    <template v-for="newsLink in newsLinks">
      <router-link :key="newsLink.link" :to="newsLink.link">
        {{ newsLink.name.toUpperCase() }}
      </router-link>
    </template>
  </nav>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "navbar-news-panel",
  computed: {
    ...mapGetters(["newsLinks"]),
  },
};
</script>

<style lang="scss" scoped>
@import "../_variables.scss";

.navbar-news-panel {
  width: 90vw;
  height: 50vh;
  background-color: white;
  color: black;
  position: fixed;
  top: 8%;
  left: 50%;
  transform: translateX(-50%);
  box-shadow: $box-shadow;
  padding: 1rem 2rem;

  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  align-items: flex-start;

  a {
    padding: 0.8rem;
  }
}
</style>