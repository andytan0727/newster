<template>
  <div class="navbar-news-panel">
    <router-link to="/news/csdn">CSDN</router-link>
    <router-link to="/news/css-tricks">CSS-Tricks</router-link>
  </div>
</template>

<script>
export default {
  name: "navbar-news-panel",
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