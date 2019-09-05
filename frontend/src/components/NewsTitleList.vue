<template>
  <ul class="news-title-list">
    <news-title-list-item
      v-for="(item, idx) in news"
      :key="item.title"
      :title="item.title"
      :url="item.url"
      :idx="idx"
    >
    </news-title-list-item>
  </ul>
</template>

<script>
import NewsTitleListItem from "@/components/NewsTitleListItem";

export default {
  name: "news-title-list",
  components: {
    NewsTitleListItem,
  },
  props: {
    news: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style lang="scss" scoped>
.news-title-list {
  width: 98%;
  margin: 0;
  padding: 0;
}
</style>