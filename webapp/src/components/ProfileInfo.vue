<template>
  <div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="profile" class="info-box">
      <div class="profile-info">
        <div class="avatar">
          <img :src="profile.path_to_photo" />
        </div>
        <FollowButton v-if="!self" :id="profile.id" />
        <h2>{{ profile.first_name }} {{ profile.last_name }}</h2>
        <p class="info">username: {{ profile.username }}</p>
        <p class="info">email: {{ profile.email }}</p>
        <p class="info">birth date: {{ profile.birth_date }}</p>
        <p class="info">About me: {{ profile.about_me }}</p>
        <Switch
          v-if="self"
          :state="profile.is_public"
          :name="'Public account'"
          @toggle="handleToggle"
        />
      </div>

      <SubsList :title="'Followers'" :users="followersList" />
      <SubsList :title="'Following'" :users="followingList" />
    </div>
  </div>
</template>

<script>
import FollowButton from "@/components/FollowButton";
import Switch from "@/components/Switch";
import SubsList from "@/components/SubsList";
import useSubscription from "@/composables/subscription.js";
import useProfile from "@/composables/profile.js";

export default {
  components: { FollowButton, Switch, SubsList},
  props: ["self", "profile"],
  setup(props) {
    const {error, setPublicity } = useProfile();
    const {
      getFollowers,
      followersList,
      getFollowing,
      followingList,
    } = useSubscription();
    getFollowers(props.profile.id);
    getFollowing(props.profile.id);

    const handleToggle = async () => {
      await setPublicity(!props.profile.is_public);
      if (!error.value) {
        props.profile.is_public = ! props.profile.is_public;
      }
    };

    return {
      error,
      handleToggle,
      followersList,
      followingList,
    };
  },
};
</script>

<style scoped>
.info-box {
  display: grid;
  grid-template-rows: 2fr 1fr 1fr;
  row-gap: 40px;
}
.profile-info {
  text-align: center;
}
.profile-info h2 {
  text-transform: capitalize;
  font-size: 28px;
  margin-top: 10px;
  margin-bottom: 10px;
}
.profile-info p {
  margin-bottom: 10px;
}

.subscription {
  margin-top: 10px;
}
.avatar {
  overflow: hidden;
  border-radius: 20px;
  position: relative;
  padding: 160px;
}
.avatar img {
  object-fit: contain;
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  min-width: 100%;
  min-height: 100%;
  max-width: 100%;
  max-height: 100%;
}

.info {
  text-align: left;
  color: #999;
}
</style>