<template>
  <form @submit.prevent="handleSubmit">
    <h3>Sign up</h3>
    <input type="text" placeholder="Username" v-model="username" />
    <input type="email" placeholder="Email" v-model="email" required />
    <input type="text" placeholder="First name" v-model="firstName" required />
    <input type="text" placeholder="Last name" v-model="lastName" required />
    <input type="date" placeholder="Date of birth" v-model="birthDate" required/>
    <input type="textarea" placeholder="About me" v-model="aboutMe" />
    <input type="password" placeholder="Password" v-model="password" required />
    <input
      type="password"
      placeholder="Confirm password"
      v-model="confirmPassword"
      required
    />
    <div v-if="validationErr" class="error">{{ validationErr }}</div>
    <div v-if="signupError" class="error">{{ signupError }}</div>
    <button v-if="!isPending">Sign up</button>
    <button v-if="isPending" disabled>Loading</button>
  </form>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";
import useValidators from "@/composables/useValidators";
import User from "@/composables/user";

export default {
  setup() {
    const { error: signupError, signup } = User();
    const { error: validationErr, validateUser } = useValidators();
    const router = useRouter();
    const username = ref("");
    const email = ref("");
    const firstName = ref("");
    const lastName = ref("");
    const birthDate = ref("");
    const aboutMe = ref("");
    const password = ref("");
    const confirmPassword = ref("");
    const isPending = ref(false);

    const handleSubmit = async () => {
      validateUser(
        username.value,
        email.value,
        password.value,
        confirmPassword.value
      );
      if (validationErr.value) {
        return;
      }

      await signup(
        username.value,
        email.value,
        firstName.value,
        lastName.value,
        birthDate.value,
        aboutMe.value,
        password.value,
        confirmPassword.value
      );

      if (signupError.value) {
        return;
      }

      router.push({name: "MainPage"})
    };

    return {
      username,
      email,
      firstName,
      lastName,
      birthDate,
      aboutMe,
      password,
      confirmPassword,
      validationErr,
      signupError,
      isPending,
      handleSubmit,
    };
  },
};
</script>

<style>
</style>