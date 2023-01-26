<template>
  <div class="flex justify-center flex-col">
    <button v-bind:class="computeClass()" class="btn btn-wide" type="button">
      {{ this.$props.range }}
    </button>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "rouletteRow",
  props: {
    range: String,
  },
  data() {
    return {
      players: [],
    };
  },
  methods: {
    computeClass() {
      let map = {
        "1 - 7": "bg-rose-700",
        0: "bg-green-700",
        "8 - 14": "bg-slate-900",
      };
      return map[this.$props.range];
    },
  },
});
</script>
<style scoped lang="scss"></style>
