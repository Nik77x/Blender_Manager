<script lang="ts">
export default {
  name: "Navbar",
};
</script>

<template>
  <div class="main">
    <h1>navbar</h1>
  </div>
</template>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.main {
  display: flex;
  flex-grow: 1;
  width: 100%;
  height: 100%;
  background-color: $background-color-darker;

  h1 {
    margin-left: 10px;
  }
}
</style>
