<template>
  <button @click="buttonClick" class="main-button">
    <span class="content">
      {{ content }}

      <img src="../../assets/images/logo-universal.png" alt="" />
    </span>
  </button>
</template>

<script setup>
const emit = defineEmits(["clicked"]);
const props = defineProps({
  content: { type: String, required: true },
});

function buttonClick() {
  emit("clicked");
}
</script>

<style lang="scss" scoped>
@import "src/assets/variables.scss";

img {
  max-width: 50px;
  padding-left: 10px;
  padding-right: 0;
  align-self: center;
}

.main-button {
  height: 50px;
  width: 90%;
  background-color: $BG-button;
  display: flex;
  color: $text-color;
  border-width: 0;
  margin: 10px auto;
  border-radius: 5px;

  img {
    margin-left: auto;
    margin-right: 10px;
  }

  span {
    width: 100%;
  }

  &:hover {
    background-color: $action-color;
  }

  &:checked {
    background-color: $action-color;
  }
}

.content {
  display: flex;
  align-items: center;
  justify-content: start;

  width: 100%;
  height: 100%;
  margin: 0;

  padding-left: 10px;
  color: white;

  font-size: 25px;

  max-font-size: 25px;
}
</style>
