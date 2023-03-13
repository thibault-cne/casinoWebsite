<template>
  <div class="flex justify-center flex-col">
    <button
      v-bind:class="computeClass()"
      class="btn btn-wide text-primary-content hover:bg-neutral-focus border-solid border-r-0 border-t-0 border-l-0 border-b-4 border-black/[.2]"
      type="button"
      :onclick="bet"
    >
      {{ $props.range }}
    </button>
    <PlayerList :listProps="players" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { socket, sendMsg } from "../utils/websocket";
import { getRequest } from "../axios/getRequest";
import PlayerList from "./playerList.vue";
import { RouletteBet } from "../models/bet";
import { Players } from "../models/players";

export default defineComponent({
  name: "rouletteRow",
  props: {
    range: { type: String, required: true },
    amount: { type: Number, required: true },
    rolling: { type: Boolean, required: true },
  },
  components: {
    PlayerList,
  },
  data() {
    return {
      players: [] as Players[],
    };
  },
  mounted() {
    this.fetchPlayers();
    socket.on("newbet", (res: any) => {
      if (res.body.color === this.getColorFromRange() && !this.$props.rolling) {
        this.addPlayer(res.body);
      }
    });
    socket.on("endgame", () => {
      setTimeout(() => {
        this.players = [];
        this.fetchPlayers();
      }, 6 * 1000);
    });
  },
  methods: {
    fetchPlayers() {
      getRequest("/roulette/bet", "json").then((r) => {
        if (r.data) {
          for (let elem of r.data) {
            if (elem.color === this.getColorFromRange()) {
              this.addPlayer(elem);
            }
          }
        }
      });
    },
    computeClass() {
      let map = new Map<string, string>([
        ["1 - 7", "bg-[#f95146]"],
        ["0", "bg-[#00c74d]"],
        ["8 - 14", "bg-[#2d3035]"],
      ]);
      return map.get(this.$props.range as string);
    },
    getColorFromRange() {
      let map = new Map<string, string>([
        ["1 - 7", "red"],
        ["0", "green"],
        ["8 - 14", "black"],
      ]);
      return map.get(this.$props.range as string);
    },
    bet() {
      if ((this.$props.amount as number) <= 0) {
        return;
      }

      const data = {
        color: this.getColorFromRange(),
        amount: this.$props.amount,
      };
      sendMsg("bet", data);
    },
    addPlayer(bet: RouletteBet) {
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
