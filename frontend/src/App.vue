<script lang="ts" setup>
import NavBar from "./components/Navbar.vue";
import VersionsList from "./components/VersionsList.vue";
import Sidebar from "./components/Sidebar.vue";
</script>

<template>
  <main>

    <nav-bar></nav-bar>

    <div class="main-content">
      <sidebar></sidebar>
      <versions-list></versions-list>
    </div>

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

.main-content {
  display: flex;
  height: 100%;

}



</style>
