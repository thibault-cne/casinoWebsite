<template>
  <v-sheet class="bg-deep-purple pa-12" height="90vh" width="100vw">
    <WheelComponent class="wheel" :isSpinning="isSpinning" :outcome="outcome" />
    <RouletteInputComponent
      :newMessage="this.newMessage"
      :loggedIn="this.loggedIn"
    />
  </v-sheet>
</template>
<script>
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
import RouletteInputComponent from "@/components/rouletteComponents/RouletteInputComponent.vue";
import { socket } from "@/websocket/roulette";

export default {
  components: {
    WheelComponent,
    RouletteInputComponent,
  },
  data() {
    return {
      newMessage: "",
      isSpinning: false,
      outcome: 9,
      loggedIn: true,
    };
  },
  mounted() {
    socket.on("endgame", () => {
      this.roll();
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
<style scoped lang="scss">
.wheel {
  padding: 10px;
}
</style>
