<template>
  <div class="pa-4">
    <ul
      class="max-w-md space-y-1 text-gray-500 list-none list-inside dark:text-gray-400"
    >
      <li v-for="user in listProps" :key="user.id">
        <div class="flex justify-around">
          <span>{{ user.user.username }}</span>
          <div>
            <span>{{ user.wager }}</span>
            <span>&nbsp;$</span>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "playerList",
  props: {
    listProps: { required: true },
  },
});
</script>
<style lang=""></style>
