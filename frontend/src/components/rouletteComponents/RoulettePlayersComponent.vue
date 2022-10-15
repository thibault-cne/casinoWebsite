<template>
  <v-card class="color-row" :style="isWinnerComp ? winningStyle : defaultStyle">
    <v-btn
      class="color-btn"
      v-bind:class="this.color"
      @click="this.playerBet()"
    >
      {{ this.displayedText }}
    </v-btn>
    <div class="best-bet">
      Best bet : {{ this.bestBet.username }} {{ this.bestBet.amount }}
    </div>
    <v-card class="players">
      <v-card-title>Players</v-card-title>
      <v-card-text><RoulettePlayersList :players="this.players" /></v-card-text>
    </v-card>
  </v-card>
</template>
<script>
import RoulettePlayersList from "./RoulettePlayersList.vue";
export default {
  props: {
    playerAmount: Number,
    color: String,
    displayedText: String,
    newBet: {},
    ws: WebSocket,
  },
  data() {
    return {
      players: [],
      bestBet: {
        username: "",
        amount: 0,
      },
      isWinnerComp: false,
      winningStyle: {
        "background-color": "green",
      },
      defaultStyle: {
        "background-color": "#967ffa",
      },
    };
  },
  watch: {
    newBet: function () {
      if (
        this.newBet.dataType === "newBet" &&
        this.newBet.data.betPosition === this.color
      ) {
        this.registerNewBet(this.newBet.data);
      }
      if (this.newBet.dataType === "endGame") {
        setTimeout(() => this.setColor(), 6000);
        setTimeout(() => this.resetValues(), 9000);
      }
    },
  },
  methods: {
    playerBet() {
      if (this.playerAmount <= 0) {
        return;
      }
      let data = { color: this.color, amount: this.playerAmount };
      console.log(data);
      // sendMsg(this.ws, JSON.stringify(data));
    },
    registerNewBet(data) {
      if (data.clientName === this.bestBet.username) {
        this.bestBet.amount += data.betAmount;
      } else {
        if (data.betAmount > this.bestBet.amount) {
          this.bestBet.username = data.clientName;
          this.bestBet.amount = data.betAmount;
        }
      }
      this.registerPlayer(data);
    },
    registerPlayer(data) {
      let playerAdded = false;
      this.players.forEach((player) => {
        if (player.clientName === data.clientName) {
          player.betAmount += data.betAmount;
          playerAdded = true;
        }
      });
      if (!playerAdded) {
        this.players.push(data);
      }
    },
    setColor() {
      if (this.newBet.data.color == this.color) {
        this.isWinnerComp = true;
      }
    },
    resetValues() {
      this.players = [];
      this.bestBet = {
        username: "",
        amount: 0,
      };
      this.isWinnerComp = false;
    },
  },
  components: { RoulettePlayersList },
};
</script>
<style scoped lang="scss">
.color-row {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 30%;
  padding: 10px;
  background-color: aqua;
}
.color-btn {
  color: #ffffff;
  width: 100px;
  height: 40px;
}

.best-bet {
  padding: 10px;
}

.players {
  width: 80%;
}

.red {
  background: #f95146;
}

.black {
  background: #2d3035;
}

.green {
  background: #00c74d;
}
</style>
