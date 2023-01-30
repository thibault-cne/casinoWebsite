<template>
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-10">
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
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-6 py-3">ID</th>
          <th scope="col" class="px-6 py-3">Game ID</th>
          <th scope="col" class="px-6 py-3">Color</th>
          <th scope="col" class="px-6 py-3">Amount</th>
        </tr>
      </thead>
      <tbody>
        <tr
          class="bg-slate-50 border-b dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
          v-bind:class="b.win ? 'dark:bg-green-800' : 'dark:bg-red-800'"
          v-for="b in this.bets"
          :key="b.id"
        >
          <th
            scope="row"
            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
          >
            {{ b.id }}
          </th>
          <td class="px-6 py-4">{{ b.game_id }}</td>
          <td class="px-6 py-4">{{ b.color }}</td>
          <td class="px-6 py-4">
            <money :amount="b.amount" />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { getRequest } from "@/axios/getRequest";
import money from "@/components/money.vue";

export default {
  name: "earningsView",
  components: { money },
  mounted() {
    getRequest("/user/get/bets", "json", { game: this.table }).then((r) => {
      this.bets = r.data;
    });
  },
  data() {
    return {
      table: "roulette",
      bets: [],
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
<style lang=""></style>
