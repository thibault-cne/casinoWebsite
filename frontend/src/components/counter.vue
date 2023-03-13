<template>
  <div class="flex text-neutral-content justify-center">
    <Digit :digit="digit" v-for="(digit, i) in split(num as number)" :key="i" />
    <span>&nbsp;$</span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { splitIntoDigits } from "../utils/splitIntoDigits";
import Digit from "./digit.vue";

export default defineComponent({
  name: "counter",
  props: {
    num: { required: true, type: Number },
  },
  components: {
    Digit,
  },
  methods: {
    split(num: number) {
      let arr = splitIntoDigits(num);
      return arr;
    },
  },
});
</script>

<style scoped lang="scss"></style>
