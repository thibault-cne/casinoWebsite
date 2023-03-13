<template>
  <div class="flex text-neutral-content">
    <div v-for="digit in split($props.amount)" :key="digit">
      <span v-if="digit === ' '">&nbsp;</span>
      <span v-else>{{ digit }}</span>
    </div>
    <span>&nbsp;$</span>
  </div>
</template>

<script lang="ts">
import { splitIntoDigits } from "../utils/splitIntoDigits";
import { defineComponent } from "vue";

export default defineComponent({
  name: "money",
  props: {
    amount: { type: Number, required: true },
  },
  methods: {
    split(num: number) {
      return splitIntoDigits(num);
    },
  },
});
</script>
<style lang=""></style>
