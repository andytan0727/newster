<template>
  <li class="news-title-list-item">
    <a :href="url" rel="noopener noreferrer" target="_blank">
      <span class="item-idx"> {{ idx + 1 }}. </span>
      <span class="title">{{ title }}</span>
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
@import "../_variables.scss";

.news-title-list-item {
  display: block;
  text-align: left;

  @media screen and (max-width: $mobile-width) {
    & a {
      display: flex;
      align-items: center;
    }
  }

  .item-idx {
    display: inline-block;
    width: 30px;
    margin-right: 1rem;
    text-align: right;

    @media screen and (max-width: $mobile-width) {
      margin-left: -10px;
      width: 30px;
    }
  }

  .title {
    display: inline-block;
    padding: 1.5rem 0;
    width: 95%;

    &:not(&:last-child) {
      border-bottom: 1px solid rgba($color: #000, $alpha: 0.5);
    }

    @media screen and (max-width: $large-display-width) {
      width: 90%;
    }

    @media screen and (max-width: $mobile-width) {
      font-size: 1rem;
      width: 240px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }

  a {
    color: black;
    text-decoration: none;
    font-size: 1.3rem;
    line-height: 2rem;
  }
}
</style>