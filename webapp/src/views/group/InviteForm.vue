<template>
  <div>
    <form v-if="successfullyInvited" @submit.prevent="handleInviteAgain">
      <p>Invite request was sent to {{ nickmail }}</p>
      <button>Invite other users</button>
    </form>
    <form v-else @submit.prevent="handleSubmit">
      <input type="text" placeholder="Username or email" v-model="nickmail" />
      <div v-if="error" class="error">{{ error }}</div>
      <button>Invite</button>
    </form>
  </div>
</template>

<script>
import { ref } from "@vue/reactivity";
import useGroupSubscription from "../../composables/groupsubscription";
export default {
  setup() {
    const { error, invite } = useGroupSubscription();
    const nickmail = ref("");
    const successfullyInvited = ref(false);

    const handleSubmit = async () => {
      await invite(nickmail.value);
      if (!error.value) {
        successfullyInvited.value = true;
      }
    };

    const handleInviteAgain = () => {
      nickmail.value = "";
      successfullyInvited.value = false;
    };

    return {
      error,
      nickmail,
      successfullyInvited,
      handleInviteAgain,
      handleSubmit,
    };
  },
};
</script>

<style scoped>
p {
    margin: 10px;
}

</style>