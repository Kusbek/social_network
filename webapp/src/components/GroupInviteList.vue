<template>
  <div v-if="groupInviteList" class="group-request-list">
    <h3>Group Join Requests</h3>
    <div v-if="error" class="error">{{ error }}</div>
    <div
      v-for="group of groupInviteList"
      :key="group.id"
      class="group-request-item"
    >
      <router-link :to="{ name: 'Group', params: { id: group.id } }">
        <div class="group-request-info">
          <p>{{ group.title }}</p>
        </div>
      </router-link>
      <button @click="handleAccept(group.id)">Accept</button>
    </div>
  </div>
</template>

<script>
import useGroupSubscription from "../composables/groupsubscription";
export default {
  setup(props) {
    const { error, groupInviteList, getGroupInviteList } = useGroupSubscription();
    getGroupInviteList();
    const handleAccept = () => {
        console.log("ACPETING")
    }
    return {
      error,
      groupInviteList,
      handleAccept,
    };
  },
};
</script>

<style scoped>
.group-request-list h3 {
  margin-bottom: 5px;
}

.group-request-item {
  display: flex;
  margin-bottom: 5px;
}
.group-request-info {
  display: flex;
}

.group-request-item button {
  margin-left: auto;
}
</style>