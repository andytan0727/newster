<template>
  <div id="app">
    <navbar></navbar>
    <router-view></router-view>
    <Footer></Footer>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar";
import Footer from "@/components/Footer";

export default {
  name: "app",
  components: {
    Navbar,
    Footer,
  },
};
</script>

<style lang="scss">
%box-shadow {
  box-shadow: 5px 5px 20px -5px rgba(black, 0.3);
}

body {
  min-height: 100vh;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

h1 {
  font-weight: 100;
}

a {
  text-decoration: none;
  color: black;
}
</style>
