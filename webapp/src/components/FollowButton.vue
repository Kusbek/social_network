<template>
  <div class="subscription">
    <button @click="handleSubscription">{{ subscriptionText }}</button>
  </div>
</template>

<script>
import useSubscription from "../composables/subscription.js";
import { computed } from "@vue/runtime-core";
export default {
  props: ["profile"],
  setup(props) {
    const { follow, unfollow, isFollowing, checkIfFollowing } = useSubscription();
    checkIfFollowing(props.profile.id);
    const subscriptionText = computed(() => {
      console.log(isFollowing)
      if (!isFollowing.value) {
        return "Follow";
      }
      return "Unfollow";
    });

    const handleSubscription = async () => {
      if (!isFollowing.value) {
        await follow(props.profile.id);
      } else {
        await unfollow(props.profile.id);
      }
    };

    return {
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