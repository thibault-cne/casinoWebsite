<template>
  <v-sheet class="bg-deep-purple pa-12" height="90vh" width="100vw">
    <WheelComponent class="wheel" :isSpinning="isSpinning" :outcome="outcome" />
    <RouletteInputComponent :newMessage="this.newMessage" :ws="this.ws" />
  </v-sheet>
</template>
<script>
import { refreshToken } from "@/axios/requests/refreshRequests";
import { authStore } from "@/store/authStore";
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
import RouletteInputComponent from "@/components/rouletteComponents/RouletteInputComponent.vue";
export default {
  components: {
    WheelComponent,
    RouletteInputComponent,
  },
  created() {
    let uri = "";

    if (authStore.getters.loggedIn) {
      uri =
        "ws://localhost:5454/api/v1/roulette/connect?accessToken=" +
        authStore.getters.accessToken;
    } else {
      uri = "ws://localhost:5454/api/v1/roulette/connect";
    }

    this.ws = new WebSocket(uri);
    this.connect();
    this.ws.onmessage = (msg) => {
      let data = JSON.parse(msg.data);

      if (data.dataType === "endGame") {
        this.outcome = data.data.number;
        this.roll();
      }

      this.newMessage = data;
    };
  },
  data() {
    return {
      newMessage: "",
      isSpinning: false,
      outcome: 9,
      ws: null,
    };
  },
  methods: {
    roll() {
      this.isSpinning = !this.isSpinning;
      setTimeout(() => {
        this.isSpinning = !this.isSpinning;
      }, 6 * 1000);
    },
    connect() {
      console.log("Attempting Connection...");

      this.ws.onopen = () => {
        console.log("Successfully Connected");
      };

      this.ws.onmessage = (msg) => {
        let data = JSON.parse(msg.data);
        if (data.errorType === 401) {
          refreshToken().then(() => {
            this.connect();
          });
          return;
        }
      };

      this.ws.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
      };

      this.ws.onerror = (error) => {
        console.log("Socket Error: ", error);
      };
    },
  },
};
</script>
<style scoped lang="scss">
.wheel {
  padding: 10px;
}
</style>
