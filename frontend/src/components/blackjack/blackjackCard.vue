<template>
  <div
    class="relative block w-32 h-48 bg-slate-50 border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 m-1"
  >
    <div class="absolute top-0 left-0 w-12 h-24">
      <div v-if="!this.$props.isFaceDown">
        <cardValue :suit="'h'" :value="'Q'" />
        <cardSuit :suit="'h'" />
      </div>
      <div v-else></div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import cardSuit from "./cardSuit.vue";
import cardValue from "./cardValue.vue";

export default defineComponent({
  name: "blackjackCard",
  props: {
    isFaceDown: {
      type: Boolean,
      required: true,
    },
  },
  components: {
    cardSuit,
    cardValue,
  },
});
</script>

<style scoped lang="scss"></style>
