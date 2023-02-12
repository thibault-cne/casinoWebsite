<template>
  <div>
    <blackjackCard :isFaceDown="false" />
  </div>
</template>

<script>
import blackjackCard from "@/components/blackjack/blackjackCard.vue";

export default {
  name: "BlackJack",
  components: {
    blackjackCard,
  },
};
</script>

<style scoped lang="scss"></style>
