<template>
  <div class="options-bar">
    <div class="option">
      <h1>Option</h1>
    </div>
    <div class="option">
      <h1>Option</h1>
    </div>
    <div class="option">
      <h1>Option</h1>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "ListViewBar",
};
</script>

<style lang="scss" scoped>
.options-bar {
  background-color: #373737;
  width: 100%;
  height: 50px;
  display: flex;
  align-content: center;
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
