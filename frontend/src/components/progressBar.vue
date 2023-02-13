<template>
  <div class="w-full bg-gray-200 rounded-full h-2.5 mb-4 dark:bg-gray-700">
    <div
      class="h-2.5 rounded-full dark:bg-teal-500 transition-all ease-out duration-1000"
      :style="{ width: this.width + '%' }"
      :class="[this.getColor(this.width)]"
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

    getColor: function (item) {
      let colors = [
        "bg-[#BB0B0B]",
        "bg-[#DB1702]",
        "bg-[#ED0000]",
        "bg-[#FF0921]",
        "bg-[#FF6347]",
        "bg-[#FF7F50]",
        "bg-[#FFE436]",
        "bg-[#FEE347]",
        "bg-[#F7E35F]",
        "bg-[#FAEA73]",
        "bg-[#FFFF00]",
        "bg-[#3AF24B]",
        "bg-[#57D53B]",
        "bg-[#34C924]",
        "bg-[#3A9D23]",
        "bg-[#228B22]",
        "bg-[#008000]",
        "bg-[#096A09]",
        "bg-[#006400]",
        "bg-[#00561B]",
        "bg-[#095228]",
      ];
      let res = Math.floor(item / 10);
      let down = 0;
      if (item % 10 >= 5) {
        down = 1;
      }
      return colors[res * 2 + down];
    },
  },
});
</script>
<style scoped lang="scss"></style>
