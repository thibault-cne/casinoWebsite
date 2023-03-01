<template>
  <button
    class="btn text-white border-solid border-r-0 border-t-0 border-l-0 border-b-4 border-black/[.2]"
    type="button"
    :onclick="this.onClick"
  >
    {{ this.$props.text }}
  </button>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "gameButton",
  props: {
    emit: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  },
  methods: {
    onClick() {
      this.$emit(this.$props.emit);
    },
  },
});
</script>
<style lang=""></style>
