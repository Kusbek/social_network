<template>
  <form @submit.prevent="handleSubmit">
    <h3>Create a group</h3>
    <input type="text" placeholder="Title" v-model="title" />
    <input
      type="textarea"
      placeholder="Group description"
      v-model="description"
    />
    <div v-if="error" class="error">{{ error }}</div>
    <button>Create</button>
  </form>
</template>

<script>
import { ref } from "@vue/reactivity";
import useGroup from "../../composables/group";
import { useRouter } from "vue-router";
export default {
  setup() {
    const { error, createGroup } = useGroup();
    const router = useRouter();
    const title = ref("");
    const description = ref("");
    const handleSubmit = async () => {
      let newGroup = await createGroup(title.value, description.value);
      console.log(newGroup);
      if (!error.value) {
        router.push({ name: "Group", params: { id: newGroup.id } });
      }
    };

    return {
      error,
      title,
      description,
      handleSubmit,
    };
  },
};
</script>

<style>
</style>