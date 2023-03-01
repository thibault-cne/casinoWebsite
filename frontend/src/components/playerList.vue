<template>
  <div class="pa-4">
    <ul
      class="max-w-md space-y-1 text-gray-500 list-none list-inside dark:text-gray-400"
    >
      <li v-for="item in listProps" :key="item.user.id">
        <div class="flex justify-around">
          <span>{{ item.user.username }}</span>
          <div>
            <money :amount="item.wager" />
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import money from "../components/money.vue";
import { Players } from "../models/players";

export default defineComponent({
  name: "playerList",
  props: {
    listProps: { required: true, type: Object as () => Players[] },
  },
  components: {
    money,
  },
});
</script>

<style lang="scss"></style>
