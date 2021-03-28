<template>
  <div class="toggle" :class="stateClass" @click="onClick">
    <div class="slider" :class="stateClass"></div>
  </div>
</template>

<script>
import { computed, ref } from "@vue/runtime-core";
export default {
  props: ["state"],
  setup(props, context) {
    const stateClass = computed(() => {
      if (props.state) {
        return "active";
      }
    });
    const onClick = () => {
      context.emit("toggle")
    };

    return {
      onClick,
      stateClass,
    };
  },
};
</script>

<style scoped>
.toggle {
  width: 100px;
  height: 50px;
  background: #fff;
  border: 2px solid #ddd;
  border-radius: 200px;
  padding: 2px;
  transition: background 0.6s;
}

.toggle.active {
  background: #72d09c;
  transition: background 0.6s;
}

.slider {
  width: 50px;
  height: 50px;
  background: #ddd;
  border-radius: 100%;
  box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.6);
  transform: translateX(0%);
  transition: transform 0.05s ease-in-out;
}

.slider.active {
  transform: translateX(100%);
}
</style>