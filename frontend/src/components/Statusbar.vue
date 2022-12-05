<template>
  <div class="status-bar">
    <p>sda</p>
  </div>
</template>

<script>
import "../assets/variables.scss";

export default {
  name: "Statusbar",
};
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.status-bar {
  background-color: $background-color-darker;
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: row;
  flex-grow: 1;

  p {
    align-self: center;
    font-size: 90%;
  }
}
</style>
