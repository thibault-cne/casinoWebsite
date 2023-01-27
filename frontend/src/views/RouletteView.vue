<template>
  <div>
    <wheelComponent class="pa-10" :isSpinning="isSpinning" :outcome="outcome" />
    <div>
      <div class="flex justify-around">
        <div class="form-control w-full max-w-xs">
          <label class="label">
            <span class="label-text">Amount</span>
          </label>
          <input
            type="number"
            placeholder="Amount"
            class="input input-bordered w-full max-w-xs"
            v-model="wager"
          />
        </div>
        <div
          class="btn-group lg:btn-group-horizontal justify-center items-center"
        >
          <button class="btn" @click="this.addWager(5000)">+5K</button>
          <button class="btn" @click="this.addWager(25000)">+25K</button>
          <button class="btn" @click="this.addWager(100000)">+100K</button>
          <button class="btn" @click="this.multiplyWager(1 / 2)">1/2</button>
          <button class="btn" @click="this.multiplyWager(2)">x2</button>
          <button
            class="btn"
            :onclick="
              () => {
                this.wager = this.user.wallet;
              }
            "
          >
            Max
          </button>
        </div>
      </div>
      <div
        class="flex items-center pa-5 md:items-start flex-col md:flex-row md:justify-around"
      >
        <rouletteRow :range="'1 - 7'" :amount="wager" />
        <rouletteRow :range="'0'" :amount="wager" />
        <rouletteRow :range="'8 - 14'" :amount="wager" />
      </div>
    </div>
  </div>
</template>

<script>
import wheelComponent from "@/components/wheelComponent.vue";
import rouletteRow from "@/components/rouletteRow.vue";
import { socket } from "@/websocket/roulette";

export default {
  components: {
    wheelComponent,
    rouletteRow,
  },
  props: {
    userProps: Object,
  },
  watch: {
    userProps: function (newVal) {
      this.user = newVal;
    },
  },
  data() {
    return {
      user: {},
      isSpinning: false,
      outcome: 9,
      wager: 0,
    };
  },
  mounted() {
    this.user = this.$props.userProps;
    socket.on("endgame", (res) => {
      this.outcome = res.body.number;
      this.roll();
    });
    socket.on(this.user.id, () => {
      console.log("User pinged");
    });
    socket.on("bet", (res) => {
      console.log(res);
      this.user.wallet -= res.body.amount;
      this.$emit("update", this.user);
    });
  },
  methods: {
    roll() {
      this.isSpinning = !this.isSpinning;
      setTimeout(() => {
        this.isSpinning = !this.isSpinning;
      }, 6 * 1000);
    },
    addWager(n) {
      this.wager += n;
    },
    multiplyWager(i) {
      this.wager = Math.trunc(this.wager * i);
    },
  },
};
</script>
<style scoped lang="scss"></style>
