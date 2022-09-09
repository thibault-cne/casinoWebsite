<template>
  <div class="card" v-bind:class="this.color">{{ number }}</div>
</template>
<script>
export default {
  props: {
    color: String,
    number: Number,
  },
};
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
.card.red {
  background: #f95146;
}

.card.black {
  background: #2d3035;
}

.card.green {
  background: #00c74d;
}
</style>
