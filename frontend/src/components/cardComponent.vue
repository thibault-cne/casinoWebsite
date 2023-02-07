<template>
  <div class="card" v-bind:class="this.computeClass()">{{ number }}</div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "cardComp",
  props: {
    color: String,
    number: Number,
  },
  methods: {
    computeClass() {
      let map = {
        red: "bg-[#f95146]",
        green: "bg-[#00c74d]",
        black: "bg-[#2d3035]",
      };
      return map[this.$props.color];
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
