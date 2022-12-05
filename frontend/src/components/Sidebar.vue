<template>
  <div class="sidebar">
    <SidebarButton
      content="Downloads"
      @clicked="changeCategory(Category.Downloads)"
    ></SidebarButton>

    <SidebarButton
      content="Library"
      @clicked="changeCategory(Category.Library)"
    ></SidebarButton>

    <SidebarButton
      class="bottom"
      content="Settings"
      @clicked="changeCategory(Category.Settings)"
    ></SidebarButton>
  </div>
</template>

<script lang="ts">
import SidebarButton from "./Button/SidebarButton.vue";
import Category from "../helpers/Category.js";
import { useCurrentCategory } from "../stores/stores";

export default {
  name: "Sidebar",
  computed: {
    Category() {
      return Category;
    },
  },
  components: { SidebarButton },

  setup() {
    let currentCategory = useCurrentCategory();

    let changeCategory = (category: Category) => {
      currentCategory.category = category;
    };

    return { changeCategory };
  },
};
</script>

<style lang="scss" scoped>
@import "../../src/assets/variables.scss";

.sidebar {
  display: flex;
  flex-flow: column;
  width: 400px;

  background-color: $background-color;
  margin: 0;
}

.bottom {
  margin-top: auto;
  margin-right: auto;
}
</style>
