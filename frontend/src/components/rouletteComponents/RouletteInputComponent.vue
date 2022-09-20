<template>
  <div>
    <v-card class="bg-deep-purple">
      <div class="bg-white input">
        <v-text-field
          v-model="amount"
          label="Amount"
          class=""
          :hide-details="true"
        >
          <template v-slot:append>
            <v-btn @click="clear">clear</v-btn>
          </template>
        </v-text-field>
        <v-card class="wallet" v-if="this.loggedIn">
          <div>{{ this.playerWallet }}</div>
          <div class="currency">$</div></v-card
        >
      </div>
      <div class="buttons">
        <v-btn-toggle tile group>
          <v-btn @click="this.addAmount(5)">+5</v-btn>
          <v-btn @click="this.addAmount(25)">+25</v-btn>
          <v-btn @click="this.addAmount(100)">+100</v-btn>
          <v-btn @click="this.multiplyAmount(0.5)">1/2</v-btn>
          <v-btn @click="this.multiplyAmount(2)">x2</v-btn>
          <v-btn @click="this.addAmount(this.playerWallet)">MAX</v-btn>
        </v-btn-toggle>
      </div>
    </v-card>
    <div class="colorRow">
      <RoulettePlayersComponent
        :color="'red'"
        :playerAmount="this.amount"
        :displayedText="'1 - 7'"
        :newBet="this.newMessage"
        :ws="this.ws"
      />
      <RoulettePlayersComponent
        :color="'green'"
        :playerAmount="this.amount"
        :displayedText="'0'"
        :newBet="this.newMessage"
        :ws="this.ws"
      />
      <RoulettePlayersComponent
        :color="'black'"
        :playerAmount="this.amount"
        :displayedText="'8 - 14'"
        :newBet="this.newMessage"
        :ws="this.ws"
      />
    </div>
  </div>
</template>
<script>
import { getRequest } from "@/axios/requests/getRequest";
import RoulettePlayersComponent from "./RoulettePlayersComponent.vue";
export default {
  components: {
    RoulettePlayersComponent,
  },
  created() {
    this.updateWallet();
  },
  props: {
    newMessage: {},
    ws: WebSocket,
    loggedIn: Boolean,
  },
  watch: {
    newMessage: function () {
      if (this.newMessage.dataType === "newBet") {
        this.updateWallet();
      }
      if (this.newMessage.dataType === "endGame") {
        setTimeout(() => this.updateWallet(), 6000);
      }
    },
  },
  data() {
    return {
      amount: 0,
      redPlayers: {
        maxAmount: 0,
        players: [],
      },
      greenPlayers: {
        maxAmount: 0,
        players: [],
      },
      blackPlayers: {
        maxAmount: 0,
        players: [],
      },
      playerWallet: 0,
    };
  },
  methods: {
    clear() {
      this.amount = 0;
    },
    addAmount(value) {
      this.amount += value;
    },
    multiplyAmount(value) {
      this.amount = Math.floor(this.amount * value);
    },
    updateWallet() {
      if (this.loggedIn) {
        getRequest("/client/data/wallet").then((r) => {
          this.playerWallet = r.data.Wallet;
        });
      }
    },
    setColor(color) {
      switch (color) {
        case "green":
          this.winningRow.green = true;
          break;
        case "black":
          this.winningRow.black = true;
          break;
        case "red":
          this.winningRow.red = true;
          break;
      }
    },
  },
};
</script>
<style scoped lang="scss">
.input {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
}

.input > * {
  flex-grow: 0;
  padding: 10px;
  width: 25%;
}

.wallet {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  width: 5%;
}

.currency {
  padding-left: 5px;
  overflow: hidden;
}

.buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
}

.colorRow {
  margin-top: 4vh;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
</style>
