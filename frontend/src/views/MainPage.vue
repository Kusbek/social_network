<template>
  <h2>Main Page</h2>
</template>

<script>
import { onMounted } from "vue";
import getUser from "@/composables/getUser";
export default {
  setup() {
    const { user } = getUser();
    onMounted(() => {

      console.log("mounted");
    });
  }
};
</script>

<style>
</style>