<template>
  <div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="group" class="group">
      <SubsList :title="'Owner'" :users="[group.owner]" />
      <div>
        <GroupInfo :group="group" />
        <PostList />
      </div>
      <div>
        <div v-if="isGroupMember">
          <router-link
            class="btn"
            :to="{ name: 'InviteForm', params: { id: id } }"
            >Invite User</router-link
          >
        </div>
        <div v-else-if="reqIsPending">
          <button disabled>Request is sent</button>
        </div>
        <div v-else>
          <button @click="handleJoinRequestClick">I want to join</button>
        </div>
        <SubsList :title="'Group Members'" :users="groupMemberList" />
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "@vue/reactivity";
import GroupInfo from "../../components/GroupInfo";
import PostList from "../../components/PostList";
import SubsList from "../../components/SubsList";
import useGroup from "../../composables/group.js";
import useGroupSubscription from "../../composables/groupsubscription";
import User from "../../composables/user";
export default {
  props: ["id"],
  components: { GroupInfo, PostList, SubsList },
  setup(props) {
    const { error, group, getGroup } = useGroup();
    const { user, getUser } = User();
    const {
      error: subsError,
      getGroupMemberList,
      groupMemberList,
      isGroupMember,
      reqIsPending,
      checkIfGroupMember,
      requestToJoin
    } = useGroupSubscription();
    getGroup(props.id);
    getGroupMemberList(props.id);
    getUser().then(() => {
      checkIfGroupMember(user.value.id, props.id);
    });

 
    const handleJoinRequestClick = async() => {
      await requestToJoin(props.id)
      if (!subsError.value){
        reqIsPending.value = true
      }
    };

    return {
      subsError,
      error,
      group,
      groupMemberList,
      isGroupMember,
      reqIsPending,
      handleJoinRequestClick,
    };
  },
};
</script>
<style scoped>
.group {
  display: grid;
  grid-template-columns: 1fr 3fr;
  column-gap: 80px;
}

.btn {
  margin: 10px;
}
</style>