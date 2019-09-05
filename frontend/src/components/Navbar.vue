<template>
  <div class="navbar">
    <router-link to="/" class="brand">Newster</router-link>

    <ul class="navbar-items">
      <li>
        <div class="dropdown" @click="handleShowNewsPanel()">
          News <dropdown-icon width="1.1rem" height="1.1rem"></dropdown-icon>
          <template v-if="showPanel">
            <navbar-news-panel></navbar-news-panel>
          </template>
        </div>
      </li>
      <li>
        <router-link to="/about">About</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import DropdownIcon from "@/components/icons/Dropdown";
import NavbarNewsPanel from "@/components/NavbarNewsPanel";

export default {
  name: "navbar",
  components: {
    DropdownIcon,
    NavbarNewsPanel,
  },
  data() {
    return {
      showPanel: false,
    };
  },
  methods: {
    handleShowNewsPanel() {
      this.showPanel = !this.showPanel;
    },
  },
};
</script>

<style lang="scss" scoped>
@import "../_variables.scss";

.brand {
  font-size: 2rem;

  @media screen and (max-width: 40em) {
    font-size: 1.4rem;
  }
}

.dropdown {
  display: flex;
  align-items: center;
  cursor: pointer;
  position: relative;

  & svg {
    padding-left: 0.3rem;
  }
}

.navbar {
  font-size: 1.3rem;
  padding: 0rem 1rem;
  display: flex;
  align-items: center;
  box-shadow: $box-shadow;

  @media screen and (max-width: 40em) {
    font-size: 1rem;
  }

  .navbar-items {
    display: flex;

    li {
      list-style-type: none;
      padding-left: 1.5rem;

      @media screen and (max-width: 40em) {
        padding-left: 0.8rem;
      }
    }
  }
}
</style>