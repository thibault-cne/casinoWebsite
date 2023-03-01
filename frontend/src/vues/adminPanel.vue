<template>
  <div>
    <button
      :onclick="toggleRoulette"
      type="button"
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
    >
      Toggle roulette
    </button>
  </div>
</template>

<script lang="ts">
import { getRequest } from "../axios/getRequest";

export default {
  name: "adminPanel",
  methods: {
    toggleRoulette() {
      getRequest("/admin/games/roulette/toggle", "json");
    },
  },
};
</script>

<style scoped lang="scss"></style>
