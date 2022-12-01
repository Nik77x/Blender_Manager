<template>
  <div class="main">
      <SidebarButton content="Downloads"></SidebarButton>
      <SidebarButton content="Library"></SidebarButton>
      <SidebarButton class="bottom" content="Settings"></SidebarButton>
  </div>

</template>

<script>
import SidebarButton from "./Button/SidebarButton.vue";

export default {
  name: "Sidebar",
  components: {SidebarButton}
}
</script>

<style scoped>

.main {
  width: 400px;
  height: 100%;
  background-color: #343434;
  margin: 0;
  display: flex;
  flex-flow: column;

}

.bottom {
  margin-top: auto;
  align-self: end;
}


</style>