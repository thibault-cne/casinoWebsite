<template>
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-4">
    <div class="flex items-center justify-end pb-4">
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
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-900 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-6 py-3">Username</th>
          <th scope="col" class="px-6 py-3">Wallet</th>
          <th scope="col" class="px-6 py-3">Status</th>
          <th scope="col" class="px-6 py-3">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr
          class="bg-gray-50 border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-slate-300 dark:hover:bg-gray-600"
          v-for="user in users"
          :key="user.id"
        >
          <td class="px-6 py-4">{{ user.username }}</td>
          <td class="px-6 py-4">
            <money :amount="user.wallet" />
          </td>
          <td class="px-6 py-4">{{ user.status }}</td>
          <td class="px-6 py-4">
            <edit-user :props-user="user" />
            <a
              type="button"
              :data-modal-target="user.id"
              :data-modal-show="user.id"
              class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
              >Edit user</a
            >
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
import { getRequest } from "@/axios/getRequest";
import { defineComponent } from "vue";
import { initModals } from "flowbite";
import editUser from "@/components/editUser.vue";
import money from "@/components/money.vue";

export default defineComponent({
  name: "userManagement",
  data() {
    return {
      users: [],
    };
  },
  mounted() {
    getRequest("/user/get/all", "json")
      .then((r) => {
        if (r.status == 200) {
          for (let element of r.data) {
            this.users.push(element);
          }
        }
      })
      .finally(() => {
        initModals();
      });
  },
  components: { editUser, money },
});
</script>
<style scoped lang="scss"></style>
