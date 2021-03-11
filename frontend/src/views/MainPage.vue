<template>
  <h2>Main Page</h2>
</template>

<script>
import User from "@/composables/user";
import { useRouter } from "vue-router";
import { onMounted } from "@vue/runtime-core";
export default {
  setup() {
    const { user, getUser } = User();
    const router = useRouter();
    
    onMounted(async () => {
      // await getUser()
      // console.log(user.value)
      // if (!user.value) {
      //   router.push({ name: "Login" });
      // }
    });
  },
};
</script>

<style>
</style>