<template>
  <li class="news-title-list-item">
    <a :href="url" rel="noopener noreferrer" target="_blank">
      {{ idx + 1 }}. {{ title }}
    </a>
  </li>
</template>

<script>
export default {
  name: "news-list-item",
  props: {
    title: {
      type: String,
      required: true,
    },
    url: {
      type: String,
      required: true,
    },
    idx: {
      type: Number,
      required: true,
    },
  },
};
</script>

<style lang="scss" scoped>
.news-title-list-item {
  display: block;
  padding: 1rem 1rem;
  text-align: left;

  &:not(:last-child) {
    border-bottom: 2px solid black;
  }

  a {
    color: black;
    text-decoration: none;
    font-size: 1.3rem;
    line-height: 2rem;
  }
}
</style>