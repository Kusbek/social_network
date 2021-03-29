<template>
  <div class="switch-container">
    <div>{{ name }}</div>
    <div class="toggle" :class="stateClass" @click="onClick">
      <div class="slider" :class="stateClass"></div>
    </div>
  </div>
</template>

<script>
import { computed, ref } from "@vue/runtime-core";
export default {
  props: ["state", "name"],
  setup(props, context) {
    const stateClass = computed(() => {
      if (props.state) {
        return "active";
      }
    });
    const onClick = () => {
      context.emit("toggle");
    };

    return {
      onClick,
      stateClass,
    };
  },
};
</script>

<style scoped>
.switch-container {
  width: 100%;
  display: flex;
  align-items: stretch;
  margin-top: 5px;
  margin-bottom: 5px;
  max-width: 1200px;
}
.toggle {
  width: 40px;
  height: 20px;
  background: #fff;
  border: 2px solid #ddd;
  border-radius: 200px;
  padding: 2px;
  transition: background 0.6s;
  margin-left: auto;
}

.toggle.active {
  background: #72d09c;
  transition: background 0.6s;
}

.slider {
  width: 20px;
  height: 20px;
  background: #ddd;
  border-radius: 100%;
  /* box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.6); */
  transform: translateX(0%);
  transition: transform 0.05s ease-in-out;
}

.slider.active {
  transform: translateX(100%);
}
</style>