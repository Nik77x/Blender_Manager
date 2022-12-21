<template>
  <main>
    <div class="nav-bar-div">
      <Nav-bar></Nav-bar>
    </div>

    <div class="main-content-div">
      <sidebar></sidebar>
      <ListView></ListView>
    </div>

    <div class="status-bar-div">
      <statusbar></statusbar>
    </div>

    <settings-view class="settings"></settings-view>
  </main>
</template>

<script lang="ts" setup>
import NavBar from "./components/Navbar.vue";
import Sidebar from "./components/Sidebar.vue";
import Statusbar from "./components/Statusbar.vue";
import ListView from "./components/ListView/ListView.vue";
import SettingsView from "./components/SettingsView.vue";
</script>

<style>
* {
  overflow-y: hidden;
  overflow-x: hidden;
  user-select: none;
}
</style>

<style scoped>
main {
  display: flex;
  flex-flow: column;
  height: 100vh;
  width: 100vw;
  background-color: navy;
}

.nav-bar-div {
  max-height: 90px;
  height: 150px;
  width: 100%;
}

.main-content-div {
  display: flex;
  flex-flow: row;
  flex-grow: 1;
  width: 100%;
}

.status-bar-div {
  max-height: 25px;
  min-height: 20px;
}

.settings {
  min-width: 700px;
  min-height: 400px;

  width: 60%;
  height: 50%;
}
</style>
