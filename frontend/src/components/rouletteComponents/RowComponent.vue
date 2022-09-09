<template>
  <div class="row">
    <CardComponent
      v-for="card in rowCards"
      :color="card.type"
      :number="card.number"
      :key="card.id"
    />
  </div>
</template>
<script>
import CardComponent from "./CardComponent.vue";
export default {
  components: {
    CardComponent,
  },
  data() {
    return {
      rowCards: [
        { type: "red", number: 1 },
        { type: "black", number: 14 },
        { type: "red", number: 2 },
        { type: "black", number: 13 },
        { type: "red", number: 3 },
        { type: "black", number: 12 },
        { type: "red", number: 4 },
        { type: "green", number: 0 },
        { type: "black", number: 11 },
        { type: "red", number: 5 },
        { type: "black", number: 10 },
        { type: "red", number: 6 },
        { type: "black", number: 9 },
        { type: "red", number: 7 },
        { type: "black", number: 8 },
      ],
    };
  },
};
</script>
<style scoped lang="scss">
.roulette-wrapper .wheel .row {
  display: flex;
}
</style>
