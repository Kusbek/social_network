<template>
  <div v-if="followRequestList" class="follow-request-list">
    <h3>Follow Requests</h3>
    <div v-if="error" class="error">{{ error }}</div>
    <div
      v-for="user of followRequestList"
      :key="user.id"
      class="follow-request-item"
    >
      <router-link :to="{ name: 'Profile', params: { id: user.id } }">
        <div class="follow-request-info">
          <div class="avatar-container">
            <img :src="user.path_to_photo" />
          </div>
          <p>{{ user.first_name }} {{ user.last_name }}</p>
        </div>
      </router-link>
      <button @click="handleAccept(user.id)">Accept</button>
    </div>
  </div>
</template>

<script>
import useSubscription from "../composables/subscription.js";
export default {
  setup() {
    const {
      error,
      followRequestList,
      acceptFollowRequest,
      getFollowRequests,
    } = useSubscription();
    getFollowRequests();
    const handleAccept = async (id) => {
      await acceptFollowRequest(id);
      if (!error.value) {
        followRequestList.value = followRequestList.value.filter(
          (user) => user.id != id
        );
      }
    };
    return {
      error,
      followRequestList,
      handleAccept,
    };
  },
};
</script>

<style scoped>
.follow-request-list h3 {
  margin-bottom: 5px;
}

.follow-request-item {
  display: flex;
  margin-bottom: 5px;
}
.follow-request-info {
  display: flex;
}

.follow-request-item button {
  margin-left: auto;
}
</style>