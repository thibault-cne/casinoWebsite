<template>
  <div class="pa-4">
    <ul
      class="max-w-md space-y-1 text-gray-500 list-none list-inside dark:text-gray-400"
    >
      <li v-for="item in listProps" :key="item.id">
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

<script>
import { defineComponent } from "vue";
import money from "@/components/money.vue";

export default defineComponent({
  name: "playerList",
  props: {
    listProps: { required: true },
  },
  components: {
    money,
  },
});
</script>
<style lang=""></style>
