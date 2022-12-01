<script lang="ts" setup>
import NavBar from "./components/Navbar.vue";
import VersionsList from "./components/VersionsList.vue";
</script>

<template>
  <main>

    <nav-bar></nav-bar>
    <versions-list></versions-list>
    <div></div>
    <div></div>

  </main>


</template>

<style lang="scss">

main {

  display: flex;
  flex-flow: column;
  height: 100vh;
  width: 100vw;
  background-color: navy;

}



</style>
