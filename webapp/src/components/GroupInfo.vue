<template>
  <div class="group-info">
    <h1>{{ group.title }}</h1>
    <p>{{ group.description }}</p>
  </div>
</template>

<script>
export default {
  props: ["group"],
};
</script>

<style scoped>
.group-info {
  background: white;
  border-radius: 8px;
  padding: 10px;
  margin-bottom: 20px;
}
</style>