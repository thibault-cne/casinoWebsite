<template>
  <div class="pa-4">
    <transition-group
      name="list"
      tag="ul"
      class="max-w-md space-y-1 text-gray-500 list-none list-inside dark:text-gray-400"
    >
      <li v-for="item in listProps" :key="item.user.id">
        <div class="flex justify-around">
          <span>{{ item.user.username }}</span>
          <div>
            <Money :amount="item.wager" />
          </div>
        </div>
      </li>
    </transition-group>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Money } from ".";
import { Players } from "../models/players";

export default defineComponent({
  name: "playerList",
  props: {
    listProps: { required: true, type: Object as () => Players[] },
  },
  components: {
    Money,
  },
});
</script>

<style scoped lang="scss">
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
