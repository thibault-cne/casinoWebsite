<template>
  <div
    v-for="digit in split(num)"
    :key="digit"
    class="flex flex-col text-neutral-content justify-center"
  >
    <i-digit :digit="digit"></i-digit>
  </div>
  <div class="flex flex-col text-neutral-content justify-center">&nbsp;$</div>
</template>

<script>
import { defineComponent } from "vue";
import { splitIntoDigits } from "@/plugins/utils";
import IDigit from "@/components/digit.vue";

export default defineComponent({
  name: "i-counter",
  props: {
    num: { required: true },
  },
  components: {
    IDigit,
  },
  methods: {
    split(num) {
      let arr = splitIntoDigits(num);
      return arr;
    },
  },
});
</script>

<style scoped lang="scss"></style>
