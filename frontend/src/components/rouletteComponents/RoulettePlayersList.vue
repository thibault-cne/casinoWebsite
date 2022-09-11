<template>
  <li v-for="player in players" v-bind:key="player.index">
    {{ player.clientName }} - {{ player.betAmount }}
  </li>
</template>
<script>
export default {
  name: "RoulettePlayerList",
  props: {
    players: Array,
  },
};
</script>
<style lang=""></style>
