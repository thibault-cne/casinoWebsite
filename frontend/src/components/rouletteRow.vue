<template>
  <div class="flex justify-center flex-col">
    <button
      v-bind:class="computeClass()"
      class="btn btn-wide"
      type="button"
      :onclick="bet"
    >
      {{ this.$props.range }}
    </button>
    <playerList :listProps="players" />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { socket, sendMsg } from "@/websocket/roulette";
import playerList from "@/components/playerList.vue";

export default defineComponent({
  name: "rouletteRow",
  props: {
    range: String,
    amount: Number,
  },
  components: {
    playerList,
  },
  data() {
    return {
      players: [],
    };
  },
  mounted() {
    socket.on("newbet", (res) => {
      if (res.body.color == this.getColorFromRange()) {
        this.addPlayer(res.body);
      }
    });
    socket.on("endgame", () => {
      this.players = [];
    });
  },
  methods: {
    computeClass() {
      let map = {
        "1 - 7": "bg-rose-700",
        0: "bg-green-700",
        "8 - 14": "bg-slate-900",
      };
      return map[this.$props.range];
    },
    getColorFromRange() {
      let map = {
        "1 - 7": "red",
        0: "green",
        "8 - 14": "black",
      };
      return map[this.$props.range];
    },
    bet() {
      if (this.$props.amount <= 0) {
        return;
      }

      const data = {
        color: this.getColorFromRange(),
        amount: this.$props.amount,
      };
      sendMsg("bet", data);
    },
    addPlayer(bet) {
      for (let p of this.players) {
        if (p.user.username == bet.user.username) {
          p.wager += bet.amount;
          return;
        }
      }

      const p = {
        wager: bet.amount,
        user: bet.user,
      };
      this.players.push(p);
    },
  },
});
</script>
<style scoped lang="scss"></style>
