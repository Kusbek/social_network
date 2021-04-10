<template>
  <div>
    <router-link
      v-for="group in groups"
      :key="group.id"
      :to="{ name: 'Group', params: { id: group.id } }"
    >
      <div class="group-item">
        <h3>{{ group.title }}</h3>
        <p>owned by {{ group.owner.first_name }} {{ group.owner.last_name }}</p>
      </div>
    </router-link>
  </div>
</template>

<script>
export default {
  props: ["groups"],
  setup(props) {
    console.log(props.groups);
  },
};
</script>

<style>
.group-item {
  display: flex;
  align-items: center;
  padding: 20px;
  border-radius: 10px;
  background: white;
  margin: 16px 0;
  transition: all ease 0.2s;
}

.group-item p {
  margin-left: auto;
}
</style>