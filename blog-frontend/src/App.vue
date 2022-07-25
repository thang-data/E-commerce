<template>
  <div id="app">
    <RouterView />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "App",
});
</script>

<style lang="scss">
// @import "@/scss/commons.sass";

#app {
  font-family: $font;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
