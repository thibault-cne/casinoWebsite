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
      </div>
      <div class="buttons">
        <v-btn-toggle tile group>
          <v-btn @click="addAmount(5)">+5</v-btn>
          <v-btn @click="addAmount(25)">+25</v-btn>
          <v-btn @click="addAmount(100)">+100</v-btn>
          <v-btn @click="multiplyAmount(0.5)">1/2</v-btn>
          <v-btn @click="multiplyAmount(2)">x2</v-btn>
          <v-btn>MAX</v-btn>
        </v-btn-toggle>
      </div>
    </v-card>
    <div class="colorRow">
      <RoulettePlayersComponent
        :color="'red'"
        :playerAmount="this.amount"
        :displayedText="'1 - 7'"
        :newBet="this.newMessage"
      />
      <RoulettePlayersComponent
        :color="'green'"
        :playerAmount="this.amount"
        :displayedText="'0'"
        :newBet="this.newMessage"
      />
      <RoulettePlayersComponent
        :color="'black'"
        :playerAmount="this.amount"
        :displayedText="'8 - 14'"
        :newBet="this.newMessage"
      />
    </div>
  </div>
</template>
<script>
import RoulettePlayersComponent from "./RoulettePlayersComponent.vue";
export default {
  components: {
    RoulettePlayersComponent,
  },
  props: {
    newMessage: {},
  },
  data() {
    return {
      newBet: {},
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
      this.amount *= value;
    },
  },
};
</script>
<style scoped lang="scss">
.input {
  display: flex;
  justify-content: center;
  align-items: center;
}

.input > * {
  flex-grow: 0;
  padding: 10px;
  width: 25%;
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
