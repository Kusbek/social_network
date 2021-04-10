<template>
  <div class="group-list">
    <h2>All Groups</h2>
    <div v-if="groups">
      <GroupList :groups="groups" />
    </div>
  </div>
</template>

<script>
import GroupList from "../../components/GroupList";
import useGroup from "../../composables/group";
export default {
  components: {GroupList},
  setup() {
    const { groups, getGroups } = useGroup();
    getGroups();
    return {
      groups,
    };
  },
};
</script>

<style>
</style>