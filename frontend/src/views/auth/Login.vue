<template>
  <form @submit.prevent="handleSubmit">
    <h3>Login</h3>
    <input type="text" placeholder="Username or Email" v-model="creds" />
    <input type="password" placeholder="Password" v-model="password" />
    <div v-if="error" class="error">{{ error }}</div>
    <button v-if="!isPending">Log in</button>
    <button v-if="isPending" disabled>Loading</button>
  </form>
</template>

<script>
import { ref } from "vue";
import useLogin from '@/composables/useLogin.js'
export default {
  setup() {
    const creds = ref("");
    const password = ref("");
    const isPending = ref(false);
    const {error, login} = useLogin()
    const handleSubmit = async () => {
      await login(creds.value, password.value)
    };


    return {
      creds,
      password,
      isPending,
      error,
      handleSubmit
    };
  },
};
</script>

<style>
</style>