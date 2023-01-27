<template>
  <div class="flex text-neutral-content">
    <div v-for="digit in split(this.$props.amount)" :key="digit">
      <span v-if="digit === ' '">&nbsp;</span>
      <span v-else>{{ digit }}</span>
    </div>
    <span>&nbsp;$</span>
  </div>
</template>

<script>
import { splitIntoDigits } from "@/plugins/utils";
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-money",
  props: {
    amount: { type: Number, required: true },
  },
  methods: {
    split(num) {
      return splitIntoDigits(num);
    },
  },
});
</script>
<style lang=""></style>
