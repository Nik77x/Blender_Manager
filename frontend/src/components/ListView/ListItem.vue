<template>
  <div class="main">
    <button>Download</button>
  </div>
</template>

<script>
export default {
  name: "ListItem",
};
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
    width: 100px;
    border-radius: 10px;
    border-width: 0;
    color: $text-color;
    font-size: 1.2rem;

    &:hover {

      filter: brightness(150%);

    }
  }

}


</style>