<template>
  <div class="w-full bg-gray-200 rounded-full h-2.5 mb-4 dark:bg-gray-700">
    <div
      class="bg-teal-500 h-2.5 rounded-full dark:bg-teal-500 transition-all ease-out duration-1000"
      :style="{ width: this.width + '%' }"
    ></div>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "progressBar",
  props: {
    time: { required: true },
  },
  watch: {
    time: function (newVal) {
      this.timeRemaining = newVal;
      this.set();
    },
  },
  data() {
    return {
      width: 100,
      intervalId: null,
      timeRemaining: 30,
    };
  },
  methods: {
    set() {
      if (this.intervalId != null) {
        clearInterval(this.intervalId);
      }

      this.intervalId = setInterval(() => {
        this.timeRemaining -= 100;
        this.width = (this.timeRemaining * 100) / 30000;
      }, 100);
    },
  },
});
</script>
<style scoped lang="scss"></style>
