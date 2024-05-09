<template>
  <Header ></Header>
  <div class="app" >
     <RouterView/>
  </div>
</template>

<script lang="ts" setup>
import Header from './components/header.vue';
</script>

<style>
.app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e00;
  /*margin-top: 60px;*/
}
</style>
