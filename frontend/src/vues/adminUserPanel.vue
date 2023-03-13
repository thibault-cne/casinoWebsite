<template>
  <div class="flex justify-center">
    <div class="w-11/12 relative overflow-x-auto sm:rounded-lg mt-4">
      <div class="flex items-center justify-end p-4">
        <label for="table-search" class="sr-only">Search</label>
        <div class="relative">
          <div
            class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"
          >
            <svg
              class="w-5 h-5 text-gray-500 dark:text-gray-400"
              aria-hidden="true"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </div>
          <input
            type="text"
            id="table-search-users"
            class="block p-2 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Search for users"
          />
        </div>
      </div>
      <table class="w-full text-sm text-left">
        <thead
          class="text-xs uppercase bg-accent text-accent-content border-b border-accent-content"
        >
          <tr>
            <th scope="col" class="px-6 py-3">Username</th>
            <th scope="col" class="px-6 py-3">Wallet</th>
            <th scope="col" class="px-6 py-3">Status</th>
            <th scope="col" class="px-6 py-3">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-accent text-accent-content">
          <tr v-for="user in users" :key="user.id">
            <td class="px-6 py-4">{{ user.username }}</td>
            <td class="px-6 py-4">
              <money :amount="user.wallet" />
            </td>
            <td class="px-6 py-4">{{ user.status }}</td>
            <td class="px-6 py-4 flex">
              <div v-if="userProps.status === 'super admin'">
                <edit-user :props-user="user" />
                <a
                  type="button"
                  :data-modal-target="user.id"
                  :data-modal-show="user.id"
                  class="font-medium text-primary hover:underline pr-4"
                  >Edit user</a
                >
              </div>
              <editWalletUser :props-user="user" />
              <a
                type="button"
                :data-modal-target="user.id + 'wallet'"
                :data-modal-show="user.id + 'wallet'"
                class="font-medium text-primary hover:underline pr-4"
                >Edit user wallet</a
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { getRequest } from "../axios/getRequest";
import { defineComponent } from "vue";
import { initModals } from "flowbite";
import editUser from "../components/editUser.vue";
import editWalletUser from "../components/editUserWallet.vue";
import money from "../components/money.vue";
import { User } from "../models/user";

export default defineComponent({
  name: "userManagement",
  props: {
    userProps: { required: true, type: Object as () => User },
  },
  data() {
    return {
      users: [] as User[],
    };
  },
  mounted() {
    getRequest("/user/get/all", "json")
      .then((r) => {
        if (r.status == 200) {
          for (let element of r.data) {
            this.users.push(element as User);
          }
        }
      })
      .finally(() => {
        initModals();
      });
  },
  components: { editUser, editWalletUser, money },
});
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
