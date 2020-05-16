<template>
  <div class="homeBox">
    <components :is="componentOn" />
  </div>
</template>

<script>
import adminComponent from "./admin";
import userComponent from "./user";
import { mapGetters } from "vuex";
export default {
  computed: {
    ...mapGetters(["userName"]),
    componentOn() {
      if (this.userName === "admin") {
        return "adminComponent";
      } else {
        return "userComponent";
      }
    },
  },
  components: {
    adminComponent,
    userComponent,
  },
};
</script>

<style lang="scss" scoped>
.homeBox {
  background-color: #ffc107;
  font-size: 14px;
  margin-bottom: 20px;
  i {
    margin-right: 5px;
    color: #ffc107;
    font-size: 18px;
  }
}
</style>
