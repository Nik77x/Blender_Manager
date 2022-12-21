<template>
  <div class="settings-main">
    <div class="header">
      <h1>Settings</h1>
      <button class="save" @click="saveSettings"></button>
      <button class="close" @click="closeSettings"></button>
    </div>
    <div class="content">
      <div class="entry">
        <p>Library Path</p>
        <input type="text" v-model="libraryPath" />
      </div>
      <div class="entry"><p>stuff</p></div>
      <div class="entry"><p>stuff</p></div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "SettingsView",
};
</script>

<script lang="ts" setup>
import { ref } from "vue";

const libraryPath = ref("");

function closeSettings() {
  // stuff
}

function saveSettings() {
  // other stuff
}
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.settings-main {
  background-color: $background-color-darker;
  position: absolute;
  top: 50%;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, -50%);
  box-sizing: border-box;

  border-radius: 10px;

  box-shadow: 0 0 17px 0 #000000;
}

.header {
  margin: 10px 20px;
  display: flex;
  flex-flow: row;

  h1 {
    font-size: 40px;
    font-family: "Monsterrat Light", serif;
    text-align: start;
    margin: 0;
  }

  button.close {
    height: 30px;
    width: 30px;
    border: transparent;
    border-radius: 210px;
    background-color: red;
    align-self: center;
    margin-left: auto;
    transition: background-color 0.3s;

    &:hover {
      background-color: lighten(red, 15%);
    }

    &:active {
      background-color: lighten(red, 30%);
    }
  }
}

.content {
  background-color: transparent;

  max-height: 90%;

  flex-grow: 1;

  margin-left: 20px;
  margin-right: 20px;
  // TODO: scroll when overflow
  overflow-y: auto;
  overflow-block: auto;

  .entry {
    display: flex;

    height: 35px;

    p {
      text-align: start;
      vertical-align: center;
      font-family: "Monsterrat Light", serif;
      font-size: 20px;
      margin-top: auto;
      margin-bottom: auto;
    }

    input[type="text"] {
      font-family: inherit;
      margin-left: 50px;
      margin-right: 10px;
      display: flex;
      flex-grow: 1;
      height: 25px;
      align-self: center;
      background-color: $background-color;
      border-radius: 20px;
      color: white;
      padding-left: 5px;
      border: 0 solid $blue-on;

      transition: border-width 70ms, background-color 0.2s;

      &:focus {
        border-width: 3px;
        outline-width: 0;
      }

      &:hover {
        background-color: lighten($background-color, 3%);
      }
    }
  }
}
</style>
