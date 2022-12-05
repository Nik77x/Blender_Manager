<script lang="ts">
import ListViewBar from "./ListViewBar.vue";
import ListItem from "./ListItem.vue";
export default {
  name: "VersionsList",
  components: { ListItem, ListViewBar },
};
</script>

<template>
  <div class="list-view">
    <ListViewBar></ListViewBar>
    <div class="list">
      <ul>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
        <li>
          <ListItem></ListItem>
        </li>
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

ul {

  flex-grow: 1;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-y: auto;
}

li {
  width: 90%;
  height: 80px;
  margin-left: auto;
  margin-right: auto;
  margin-top: 10px;
  margin-bottom: 10px;

  div {
    display: flex;
    flex-grow: 1;
    width: 100%;
    height: 100%;
  }
}
</style>