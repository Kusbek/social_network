<template>
  <div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="profile">
      <div
        v-if="parseInt(id) !== user.id && profileIsClosed"
        class="profile-closed"
      >
        <h1>This profile is closed. You can request to follow</h1>
        <FollowButton :id="id" />
      </div>
      <div v-else class="profile-details">
        <ProfileInfo :self="parseInt(id) === user.id" :profile="profile" />
        <div class="post-box">
          <PostList />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import User from "../composables/user.js";
import PostList from "../components/PostList";
import ProfileInfo from "../components/ProfileInfo";
import FollowButton from "../components/FollowButton";
import useSubscription from "../composables/subscription.js";
import useProfile from "@/composables/profile.js";
import { computed } from "@vue/runtime-core";

export default {
  props: ["id"],
  components: { PostList, ProfileInfo, FollowButton },
  setup(props) {
    const { user, getUser } = User();
    const { isFollowing, checkIfFollowing } = useSubscription();
    const { profile, error, load } = useProfile();
    load(props.id)
    getUser();
    checkIfFollowing(props.id);
    const profileIsClosed = computed(() => {
      if (profile.value.is_public) {
        return false
      }
      return !isFollowing.value;
    });

    return {
      profile,
      user,
      error,
      profileIsClosed,
    };
  },
};
</script>

<style>
.profile-details {
  display: grid;
  grid-template-columns: 1fr 3fr;
  column-gap: 80px;
}

.profile-closed {
  text-align: center;
}
</style>