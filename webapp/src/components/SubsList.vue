<template>
  <div class="subscription-list">
    <h3>{{ title }}</h3>
    <router-link
      v-for="user of users"
      :key="user.id"
      :to="{ name: 'Profile', params: { id: user.id } }"
    >
      <div class="subscription-item">
        <div class="avatar-container">
          <img :src="user.path_to_photo" />
        </div>
        <p>{{ user.first_name }} {{ user.last_name }}</p>
      </div>
    </router-link>
  </div>
</template>

<script>
export default {
  props: ["title", "users"],
  setup(props) {
    
  } 
};
</script>

<style scoped>
.subscription-list h3 {
  margin-bottom: 5px;
}

.subscription-item {
  display: flex;
  margin-bottom: 5px;
}

/* .avatar-container {
  overflow: hidden;
  height: 20px;
  width: 20px;
  margin-left: 5px;
  margin-right: 5px;
} */

/* .avatar-container img {
  top: 0;
  left: 0;
  min-width: 100%;
  min-height: 100%;
  max-width: 100%;
  max-height: 100%;
} */
</style>