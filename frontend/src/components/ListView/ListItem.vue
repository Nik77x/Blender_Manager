<template>
  <div class="main">
    <div class="content">
      <h1 class="title">{{ props.title }}</h1>
      <hr class="vertical-separator"/>
      <div class="sub-info">
        <p class="build-branch">{{ props.buildBranch }}</p>
        <hr class="horizontal-separator"/>
        <p class="build-date">{{ props.buildDate }}</p>
        <hr class="horizontal-separator"/>
        <p class="build-commit">{{ props.commitHash }}</p>
      </div>
    </div>
    <button>Download</button>
  </div>
</template>

<script lang="ts">
export default {
  name: "ListItem",
};
</script>

<script lang="ts" setup>
const props = defineProps<{
  title: String;
  buildBranch: String;
  buildDate: String;
  commitHash: String;
  commitLink: String;
  downloadLink: String;
}>();

function download() {
  // TODO: Download blender, unpack, store, link, run, ??, profit
}
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.main {
  background-color: $background-color;
  border-radius: 20px;
  border-color: black;
  border-width: 10px;
  display: flex;

  button {
    --bright: 100%;

    background-color: #67be16;
    margin-left: auto;
    margin-right: 20px;
    align-self: center;
    height: 70%;
    min-width: 95px;
    width: 120px;
    border-radius: 10px;
    border-width: 0;
    color: $text-color;
    text-align: start;
    font-size: 1.2rem;
    font-family: "Nunito", serif;

    &:hover {
      filter: brightness(150%);
    }
  }
}

.content {
  width: 100%;
  height: 90%;
  background-color: transparent;
  align-self: center;
  display: flex;
  flex-flow: column;
  margin: 5px 15px;

  h1.title {
    align-self: start;
    margin-top: 7px;
    margin-bottom: 0;
    font-family: "Monsterrat Thin", serif;
    font-size: 27px;
  }

  .sub-info {
    display: flex;

    p {
      align-self: start;
      margin-bottom: auto;
      margin-top: 0;
      font-family: "Roboto Light", serif;
      font-size: 13px;
    }

    .build-branch {
      // TODO: If build is stable make this green
      color: red;
    }

    .build-date {
      color: lighten($background-color, 30%);
    }

    .build-commit {
      color: lighten($background-color, 30%);
    }
  }
}

.vertical-separator {
  width: 70%;
  margin: 3px;
  border: 1px solid $background-color-lighter;
}

.horizontal-separator {
  margin: 0 7px;
  width: 3px;
  height: 3px;
  background-color: lighten($background-color-lighter, 30%);
  align-self: center;
  border: 1px solid $background-color-lighter;
  border-radius: 5px;
}
</style>
