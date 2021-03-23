<template>
  <div class="subscription" v-if="profile.id !== user.id">
    <button @click="handleSubscription">{{ subscriptionText }}</button>
  </div>
</template>

<script>
import User from "../composables/user.js";
import useProfile from "../composables/profile.js";
import { computed } from "@vue/runtime-core";
export default {
  props: ["profile"],
  setup(props) {
    const { user, getUser } = User();
    const { follow, unfollow } = useProfile();
    getUser();

    const subscriptionText = computed(() => {
      if (!props.profile.isFollowing) {
        return "Follow";
      }
      return "Unfollow";
    });

    const handleSubscription = async () => {
      if (!props.profile.isFollowing) {
        await follow(props.profile.id);
      } else {
        await unfollow();
      }
    };

    return {
      user,
      subscriptionText,
      handleSubscription,
    };
  },
};
</script>

<style>
.subscription {
  margin-top: 10px;
}
</style>