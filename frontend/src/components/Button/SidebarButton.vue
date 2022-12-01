<template>
  <button class="main">
    <div class="flex">

      <img src="src/assets/images/logo-universal.png" alt="">
      <h1 class="content">{{content}}</h1>


    </div>

  </button>
</template>

<script>

export default {

  name: "SidebarButton",

  setup() {

  },

  props: {

    content: String

  }
}
</script>

<style lang="scss" scoped>

img {
  max-width: 50px;
  padding-left: 5px;
  align-self: center;
}

.flex {
  display: flex;
  align-content: center;
}

.main {
  height: 50px;
  width: 100%;
  background-color: black;

}
.content {
  padding: 0;
  margin: 0 auto 0 20px;
  align-self: center;

  color: white;

  max-font-size: 25px;
}
</style>