<template>
  <v-sheet class="bg-deep-purple pa-12" height="90vh" width="100vw">
    <WheelComponent :isSpinning="isSpinning" :outcome="outcome" />
    <RouletteInputComponent :newMessage="this.newMessage" />
  </v-sheet>
</template>
<script>
import { connect, socket } from "@/websocket";
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
import RouletteInputComponent from "@/components/rouletteComponents/RouletteInputComponent.vue";
export default {
  components: {
    WheelComponent,
    RouletteInputComponent,
  },
  created() {
    connect();
    socket.onmessage = (msg) => {
      let data = JSON.parse(msg.data);
      switch (data.dataType) {
        case "endGame":
          this.outcome = data.data.number;
          this.roll();
          break;
        case "newBet":
          this.newMessage = data.data;
          break;
        default:
          break;
      }
    };
  },
  data() {
    return {
      newMessage: "",
      isSpinning: false,
      outcome: 9,
    };
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
<style lang="scss"></style>
