<template>
  <div class="sidebar">
    <SidebarButton
      content="Downloads"
      icon="src/assets/icons/arrow-down-solid.svg"
      @clicked="changeCategory(Category.Downloads)"
    ></SidebarButton>

    <SidebarButton
      content="Library"
      icon="src/assets/icons/list-ul-solid.svg"
      @clicked="changeCategory(Category.Library)"
    ></SidebarButton>

    <SidebarButton
      class="bottom"
      content="Settings"
      icon="src/assets/icons/gear-solid.svg"
      @clicked="changeCategory(Category.Settings)"
    ></SidebarButton>
  </div>
</template>

<script lang="ts">
import VueResizable from "vue-resizable";
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
  components: { SidebarButton, VueResizable },

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
  height: 100%;
  width: 400px;

  background-color: $background-color;
  margin: 0;
}

.bottom {
  margin-top: auto;
  margin-right: auto;
}
</style>
