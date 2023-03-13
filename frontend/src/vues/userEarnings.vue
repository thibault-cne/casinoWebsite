<template>
  <div class="flex justify-center">
    <div class="w-11/12 relative overflow-x-auto sm:rounded-lg mt-4">
      <!--
      <div class="flex items-center justify-between pb-4">
        <div>
          <select
            name="table"
            id="table"
            class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            v-model="table"
            @change="onChange"
          >
            <option>roulette</option>
          </select>
        </div>
      </div>
    -->
      <table class="w-full text-left">
        <thead class="text-xs uppercase bg-accent text-accent-content">
          <tr>
            <th scope="col" class="px-6 py-3">ID</th>
            <th scope="col" class="px-6 py-3">Game ID</th>
            <th scope="col" class="px-6 py-3">Color</th>
            <th scope="col" class="px-6 py-3">Amount</th>
          </tr>
        </thead>
        <tbody class="bg-accent text-accent-content">
          <tr
            class="hover:bg-accent-focus"
            v-bind:class="b.win ? 'bg-success' : 'bg-error'"
            v-for="b in bets"
            :key="b.id"
          >
            <td scope="row" class="px-6 py-4 whitespace-nowrap">
              {{ b.id }}
            </td>
            <td class="px-6 py-4">{{ b.game_id }}</td>
            <td class="px-6 py-4">{{ b.color }}</td>
            <td class="px-6 py-4">
              <Money :amount="b.amount" />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  <div class="w-11/12 flex flex-col justify-center mt-10"></div>
</template>

<script lang="ts">
import { getRequest } from "../axios/getRequest";
import { Money } from "../components";
import { RouletteResult } from "../models/bet";

export default {
  name: "userEarnings",
  components: { Money },
  mounted() {
    getRequest("/user/get/bets", "json", { game: this.table }).then((r) => {
      this.bets = r.data;
    });
  },
  data() {
    return {
      table: "roulette",
      bets: [] as RouletteResult[],
    };
  },
  methods: {
    onChange() {
      getRequest("/user/get/bets", "json", { game: this.table }).then((r) => {
        this.bets = r.data;
      });
    },
  },
};
</script>
<style scoped lang="scss">
/* top-left border-radius */
table tr:first-child th:first-child {
  border-top-left-radius: 6px;
}

/* top-right border-radius */
table tr:first-child th:last-child {
  border-top-right-radius: 6px;
}

/* bottom-left border-radius */
table tr:last-child td:first-child {
  border-bottom-left-radius: 6px;
}

/* bottom-right border-radius */
table tr:last-child td:last-child {
  border-bottom-right-radius: 6px;
}
</style>
