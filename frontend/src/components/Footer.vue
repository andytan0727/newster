<template>
  <div class="footer">
    &copy; {{ new Date().getFullYear() }} | Built with
    <a href="https://golang.org">Go</a> and
    <a href="https://vuejs.org">Vue</a> by Andy Tan
  </div>
</template>

<script>
export default {
  // use uppercase because footer is a reserved word in html
  name: "Footer",
};
</script>

<style lang="scss" scoped>
.footer {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  border-top: 2px solid black;
  padding: 1rem 0;

  @media screen and (max-width: 40em) {
    font-size: 0.8rem;
  }

  a {
    font-weight: bold;
    padding: 0 0.5rem;
  }
}
</style>