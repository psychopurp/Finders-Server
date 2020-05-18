<template>
  <div class="wrapper">
    <m-header></m-header>

    <div class="wrapper_con">
      <side-bar />
      <page-main />
    </div>
  </div>
</template>

<script>
import MHeader from "./components/Header";
import SideBar from "./components/SideBar";
import PageMain from "./components/PageMain";
export default {
  name: "layout",
  components: {
    MHeader,
    SideBar,
    PageMain,
  },
};
</script>

<style></style>
