<template>
  <v-sheet class="bg-deep-purple pa-12" height="90vh" width="100vw">
    <WheelComponent :isSpinning="isSpinning" :outcome="outcome" />
    <button @click="roll">Click</button>
  </v-sheet>
</template>
<script>
import { connect, socket } from "@/websocket";
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
export default {
  components: {
    WheelComponent,
  },
  mounted() {
    connect();
    socket.onmessage = (msg) => {
      let data = JSON.parse(msg.data);
      if (data.dataType === "endGame") {
        this.outcome = data.data.number;
        this.roll();
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
