<template>
  <div class="container">
    <div class="progress">
      <semipolar-spinner
        class="spinner"
        :animation-duration="2000"
        :size="15"
        color="#fff"
      />
      <h1>{{ progress_percent }}% {{ message }}</h1>
    </div>
  </div>
</template>

<script lang="ts">
import { SemipolarSpinner } from "epic-spinners";

export default {
  name: "Progressbar",
  components: { SemipolarSpinner },
};
</script>

<script lang="ts" setup>
import { computed } from "vue";

const props = defineProps<{
  max: number;
  value: number;
  message?: string;
}>();

const progress_percent = computed(() => {
  // calculate percentage
  let val = (props.value / props.max) * 100;

  // round to max 2 decimal places
  val = parseFloat(val.toFixed(2));

  // clamp between 0 - 100
  return Math.min(Math.max(val, 0), 100);
});
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

.container {
  background-color: $background-color-darker;

  height: 20px;
  border: solid 3px lighten($background-color-darker, 7%);
  border-radius: 100px;
  padding: 4px;
  display: flex;

  .progress {
    min-width: 10%;
    max-width: 100%;
    // val : max = x : 100
    // width: calc((v-bind(value) / v-bind(max) * 100) * 100%);
    width: calc(v-bind(progress_percent) * 1%);
    background-color: $blue-on;
    border-radius: 100px;
    display: flex;
    padding: 5px;
  }

  .spinner {
    align-self: center;
    margin-right: 3px;
  }

  h1 {
    font-size: 10px;
    align-self: center;
    font-family: "Monsterrat Thin", sans-serif;
  }
}
</style>
