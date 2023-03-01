<template>
  <div class="card" v-bind:class="computeClass()">{{ number }}</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "cardComp",
  props: {
    color: String,
    number: Number,
  },
  methods: {
    computeClass(): string {
      let map = new Map<string, string>([
        ["red", "bg-[#f95146]"],
        ["green", "bg-[#00c74d]"],
        ["black", "bg-[#2d3035]"],
      ]);

      return map.get(this.$props.color as string) as string;
    },
  },
});
</script>

<style scoped lang="scss">
.roulette-wrapper .wheel .row .card {
  height: 75px;
  width: 75px;
  margin: 3px;
  border-radius: 8px;
  border-bottom: 3px solid rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.5em;
}
</style>
