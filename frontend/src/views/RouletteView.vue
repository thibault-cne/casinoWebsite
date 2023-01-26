<template>
  <div>
    <WheelComponent class="pa-10" :isSpinning="isSpinning" :outcome="outcome" />
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
        <div></div>
      </div>
      <div class="flex justify-around pa-5">
        <rouletteRow :range="'1 - 7'" />
        <rouletteRow :range="0" />
        <rouletteRow :range="'8 - 14'" />
      </div>
    </div>
  </div>
</template>

<script>
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
import rouletteRow from "@/components/rouletteRow.vue";
import { socket } from "@/websocket/roulette";

export default {
  components: {
    WheelComponent,
    rouletteRow,
  },
  props: {
    userProps: Object,
  },
  watch: {
    userProps: function () {
      this.user = this.$props.userProps;
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
    socket.on("endgame", () => {
      this.roll();
    });
    socket.on(this.user.id, () => {
      console.log("User pinged");
    });
  },
  methods: {
    roll() {
      this.isSpinning = !this.isSpinning;
      setTimeout(() => {
        this.isSpinning = !this.isSpinning;
      }, 6 * 1000);
    },
  },
};
</script>
<style scoped lang="scss"></style>
