<template>
  <div class="flex justify-center h-full">
    <div class="w-full md:w-8/12 flex-col">
      <section
        class="flex justify-center transition-all ease-in-out duration-200 scale-90 mt-8"
      >
        <transition-group name="deal" tag="div" class="flex">
          <blackjackCard
            v-for="(val, i) in dealerCards"
            :isFaceDown="val"
            :key="i"
          />
        </transition-group>
      </section>
      <section
        class="flex justify-center transition-all ease-in-out duration-200 scale-125 mt-36"
      >
        <transition-group
          name="deal"
          tag="div"
          class="flex flex-wrap justify-center items-center"
        >
          <blackjackCard
            v-for="(val, i) in handCards"
            :isFaceDown="val"
            :key="i"
          />
        </transition-group>
      </section>
    </div>
    <div class="flex-col justify-center mt-44">
      <div class="form-control w-full max-w-xs justify-center">
        <label class="label">
          <span class="label-text">Amount</span>
        </label>
        <input
          type="number"
          placeholder="Amount"
          class="input input-bordered w-full max-w-xs dark:text-white"
          v-model="wager"
        />
      </div>
      <div class="flex justify-center mt-5">
        <div class="pa-2">
          <gameButton :text="'Hit'" :emit="'hit'" @hit="hit" />
        </div>
        <div class="pa-2">
          <gameButton :text="'Stand'" :emit="'stand'" @stand="stand" />
        </div>
      </div>
      <div class="flex justify-center mt-1">
        <gameButton
          :text="'Double down'"
          :emit="'doubledown'"
          @doubledown="doubleDown"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import blackjackCard from "../components/blackjack/blackjackCard.vue";
import gameButton from "../components/blackjack/gameButton.vue";

export default {
  name: "BlackJack",
  components: {
    blackjackCard,
    gameButton,
  },
  props: {
    userProps: { required: true, type: Object },
  },
  mounted() {
    this.dealerCards.push(true);
    this.dealerCards.push(false);

    this.handCards.push(false);
    this.handCards.push(false);
    this.handCards.push(false);
    this.handCards.push(false);
  },
  data() {
    return {
      dealerCards: [] as boolean[],
      handCards: [] as boolean[],
      wager: 0,
    };
  },
  methods: {
    hit() {},
    stand() {},
    doubleDown() {},
  },
};
</script>

<style scoped lang="scss">
.deal-enter-active {
  animation: deal 0.3s;
}
.deal-leave-active {
  animation: deal 0.3s reverse;
}

@keyframes deal {
  0% {
    transform: translateY(-100vh);
  }
  100% {
    transform: translateY(0);
  }
}
</style>
