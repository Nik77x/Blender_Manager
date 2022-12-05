<script lang="ts">
import ListViewBar from "./ListViewBar.vue";

export default {
  name: "VersionsList",
  components: { ListViewBar },
};
</script>

<template>
  <div class="list-view">
    <ListViewBar></ListViewBar>
    <div class="list">
      <ul>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
        <li>yes</li>
      </ul>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.list-view {
  display: flex;
  max-height: 100%;
  overflow-y: hidden;
  flex-flow: column;
  background-color: $background-color;
  width: 100%;
}

.options-bar {
  background-color: $background-color-header;
  width: 100%;
  height: 50px;
  display: flex;
  align-content: center;
}

.list {
  background-color: $background-color-lighter;
  width: 100%;
  height: 100%;

  overflow-y: auto;
}

.option {
  background-color: #202020;
  width: 75px;

  height: 37px;

  margin: 5px;
  border-radius: 50px;
  display: flex;

  align-self: center;

  h1 {
    font-size: 15px;
    text-align: center;
    align-self: center;
    width: 100%;
    margin: 0;
  }
}
</style>
