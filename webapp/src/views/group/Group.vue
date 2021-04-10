<template>
  <div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="group" class="group">
      <SubsList :title="'Owner'" :users="[group.owner]" />
      <div>
        <GroupInfo :group="group" />
        <PostList />
      </div>
      <div>Group Members</div>
    </div>
  </div>
</template>

<script>
import { ref } from "@vue/reactivity";
import GroupInfo from "../../components/GroupInfo";
import PostList from "../../components/PostList";
import SubsList from "../../components/SubsList";
import useGroup from "../../composables/group.js";

export default {
  props: ["id"],
  components: { GroupInfo, PostList, SubsList },
  setup(props) {
    // const error = ref(null);
    // const group = ref({
    //   id: props.id,
    //   owner: {
    //     id: 1,
    //     first_name: "Bekarys",
    //     last_name: "Kuspan",
    //     path_to_photo: "/img/avatars/2021-03-15 20.32.55.jpg",
    //   },
    //   title: "test group title",
    //   description: "test group descriptions",
    // });

    const { error, group, getGroup } = useGroup();
    getGroup(props.id);
    return {
      error,
      group,
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
</style>