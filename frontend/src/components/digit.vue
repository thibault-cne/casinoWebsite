<template>
  <span>
    <span v-if="d == '.'">.</span>
    <span v-else-if="d == ' '">&nbsp;</span>
    <span v-else class="single-digit">
      <span :style="{ '--value': d }"></span>
    </span>
  </span>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "digit",
  props: {
    digit: { required: true, type: String },
  },
  watch: {
    digit: function (newVal) {
      this.d = newVal;
    },
  },
  data() {
    return {
      d: "",
    };
  },
  mounted() {
    setTimeout(() => {
      this.d = this.$props.digit;
    }, 200);
  },
});
</script>

<style scoped lang="scss">
.single-digit {
  line-height: 1em;
  display: inline-flex;
}

.single-digit > * {
  height: 1em;
  display: inline-block;
  overflow-y: hidden;
}
.single-digit > *:before {
  position: relative;
  content: "0\A 1\A 2\A 3\A 4\A 5\A 6\A 7\A 8\A 9";
  white-space: pre;
  top: calc(var(--value) * -1em);
  text-align: center;
  transition: all 1s cubic-bezier(1, 0, 0, 1);
}
</style>
