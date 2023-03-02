<template>
  <div>
    <div class="p-10">
      <progressBar :time="time" />
      <wheelComponent :isSpinning="isSpinning" :outcome="outcome" />
    </div>
    <div>
      <div class="flex justify-around">
        <div class="form-control w-full max-w-xs">
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
        <div class="btn-group lg:btn-group-horizontal justify-center items-end">
          <button class="btn" @click="addWager(5000)">+5K</button>
          <button class="btn" @click="addWager(25000)">+25K</button>
          <button class="btn" @click="addWager(100000)">+100K</button>
          <button class="btn" @click="multiplyWager(1 / 2)">1/2</button>
          <button class="btn" @click="multiplyWager(2)">x2</button>
          <button
            class="btn"
            :onclick="
              () => {
                wager = user.wallet;
              }
            "
          >
            Max
          </button>
        </div>
      </div>
      <div
        class="flex items-center p-5 md:items-start flex-col md:flex-row md:justify-around"
      >
        <rouletteRow :range="'1 - 7'" :amount="wager" :rolling="isSpinning" />
        <rouletteRow :range="'0'" :amount="wager" :rolling="isSpinning" />
        <rouletteRow :range="'8 - 14'" :amount="wager" :rolling="isSpinning" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import wheelComponent from "../components/wheel.vue";
import rouletteRow from "../components/rouletteRow.vue";
import progressBar from "../components/progressBar.vue";
import { socket } from "../utils/websocket";
import { getRequest } from "../axios/getRequest";
import { User } from "../models/user";

export default {
  components: {
    wheelComponent,
    rouletteRow,
    progressBar,
  },
  props: {
    userProps: { type: Object as () => User, required: true },
  },
  watch: {
    userProps: function (newVal) {
      this.user = newVal;
    },
  },
  data() {
    return {
      user: {} as User,
      isSpinning: false,
      outcome: 9,
      wager: 0,
      time: 30,
    };
  },
  mounted() {
    getRequest("/roulette/time", "json").then((r) => {
      this.setProgress(r.data.time);
    });

    this.user = this.$props.userProps;
    socket.on("endgame", (res: any) => {
      this.outcome = res.body.number;
      this.roll();
    });
    socket.on(this.user.id, () => {});
    socket.on("start", (res: any) => {
      this.setProgress(res.body.time);
    });
    socket.on("bet", (res: any) => {
      this.user.wallet -= res.body.amount;
      this.$emit("update", this.user);
    });
  },
  methods: {
    roll() {
      this.isSpinning = !this.isSpinning;
      setTimeout(() => {
        this.isSpinning = !this.isSpinning;
        this.$emit("refresh");
      }, 6 * 1000);
    },
    addWager(n: number) {
      this.wager += n;
    },
    multiplyWager(i: number) {
      this.wager = Math.trunc(this.wager * i);
    },
    setProgress(endTime: number) {
      this.time = endTime - Date.now();
    },
  },
};
</script>
<style scoped lang="scss"></style>
