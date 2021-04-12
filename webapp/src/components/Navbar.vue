<template>
  <div class="navbar">
    <nav>
      <img src="@/assets/ninja.jpg" alt="" />
      <h1>
        <router-link :to="{ name: 'MainPage' }">Social Network</router-link>
      </h1>
      <div class="links">
        <div v-if="user">
          <!-- <router-link :to="{ name: 'CreatePlaylist' }">Create Playlist</router-link>
          <router-link :to="{ name: 'UserPlaylists' }">My Playlists</router-link> -->
          <router-link class="btn" :to="{ name: 'Groups' }"
            >All Groups</router-link
          >
          <router-link class="btn" :to="{ name: 'CreateGroupForm' }"
            >New Group</router-link
          >
          <router-link :to="{ name: 'Profile', params: { id: user.id } }">
            <span>Hi there, {{ user.username }}</span>
          </router-link>
          <button @click="handleClick">Logout</button>
        </div>
        <div v-else>
          <router-link class="btn" :to="{ name: 'Signup' }"
            >Sign up</router-link
          >
          <router-link class="btn" :to="{ name: 'Login' }">Login</router-link>
        </div>
      </div>
    </nav>
  </div>
</template>
<script>
import User from "../composables/user";
import { useRouter } from "vue-router";
export default {
  setup() {
    const { logout } = User();
    const router = useRouter();
    const { user } = User();

    const handleClick = async () => {
      await logout();
      router.push({ name: "Login" });
    };
    return { user, handleClick };
  },
};
</script>

<style scoped>
.navbar {
  padding: 16px 10px;
  margin-bottom: 60px;
  background: white;
}
nav {
  display: flex;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
}
nav img {
  max-height: 60px;
}
nav h1 {
  margin-left: 20px;
}
nav .links {
  margin-left: auto;
}
nav .links a,
button {
  margin-left: 16px;
  font-size: 14px;
}
span {
  font-size: 14px;
  display: inline-block;
  margin-left: 16px;
  padding-left: 16px;
  border-left: 1px solid #eee;
}
</style>