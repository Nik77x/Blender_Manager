<template>
  <div class="main">
    <div class="content">
      <h1 class="title">{{ props.blendInfo.version_title }}</h1>

      <hr class="vertical-separator" />
      <transition name="slide" mode="out-in">
        <div class="sub-info" v-if="!isDownloading">
          <p class="build-branch">{{ props.blendInfo.build_type }}</p>
          <hr class="horizontal-separator" />
          <p class="build-date">{{ props.blendInfo.upload_date }}</p>
          <hr class="horizontal-separator" />
          <p class="build-commit">{{ props.blendInfo.commit_hash }}</p>
          <hr class="horizontal-separator" />
        </div>
        <progressbar
          class="progress-bar"
          :value="progress === undefined ? 0 : progress"
          :max="100"
          v-else
        ></progressbar>
      </transition>
      <!--      <semipolar-spinner-->
      <!--          :animation-duration="2000"-->
      <!--          :size="65"-->
      <!--          color="#ff1d5e"-->
      <!--      />-->
    </div>
    <button @click="download">Download</button>
  </div>
</template>

<script lang="ts">
export default {
  name: "ListItem",
};
</script>

<script lang="ts" setup>
import { Data } from "../../../wailsjs/go/models";
import { DownloadBlender } from "../../../wailsjs/go/backend/App";
import { ref } from "vue";
import { EventsOn } from "../../../wailsjs/runtime";
import Progressbar from "../Progressbar.vue";

const props = defineProps<{
  blendInfo: Data.BlendInfo;
}>();

const progress = ref<number>();

const isDownloading = ref(false);

EventsOn(props.blendInfo.commit_hash + "_dl-progress", (progress_val) => {
  progress.value = progress_val * 100;
});

function download() {
  isDownloading.value = true;
  DownloadBlender(props.blendInfo);
}
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

// Transition things //
.slide-leave-active,
.slide-enter-active {
  transition: all 0.5s ease;
}

.slide-leave-to,
.slide-enter-from {
  transform: translateX(-500px);
}

// Main style //
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
    text-align: center;
    font-size: 1.2rem;
    font-family: "Nunito", serif;

    &:hover {
      filter: brightness(150%);
    }
  }
}

.progress-bar {
  width: 80%;
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
