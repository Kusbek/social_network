<template>
  <div v-if="error" class="error">{{ error }}</div>
  <div v-if="profile" class="profile-details">
    <div class="profile-info">
      <div class="avatar">
        <img :src="profile.pathToPhoto" />
      </div>
    </div>
    <PostList />
  </div>
</template>

<script>
import useProfile from "@/composables/profile.js";
import PostList from "../components/PostList";
export default {
  props: ["id"],
  components: { PostList },
  setup(props) {
    const { getProfile } = useProfile();
    const { profile, error } = getProfile();
    return {
      profile,
      error,
    };
  },
};
</script>

<style>
.profile-details {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 80px;
}
.profile-info {
  text-align: center;
}
.profile-info h2 {
  text-transform: capitalize;
  font-size: 28px;
  margin-top: 20px;
}
.profile-info p {
  margin-bottom: 20px;
}

.avatar {
  overflow: hidden;
  border-radius: 20px;
  position: relative;
  padding: 160px;
}
.avatar img {
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  min-width: 100%;
  min-height: 100%;
  max-width: 200%;
  max-height: 200%;
}
</style>