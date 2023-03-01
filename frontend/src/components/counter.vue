<template>
  <div class="flex text-neutral-content justify-center">
    <i-digit
      :digit="digit"
      v-for="digit in split(num as number)"
      :key="digit"
    />
    <span>&nbsp;$</span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { splitIntoDigits } from "../utils/splitIntoDigits";
import IDigit from "../components/digit.vue";

export default defineComponent({
  name: "i-counter",
  props: {
    num: { required: true, type: Number },
  },
  components: {
    IDigit,
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
