<template>
  <div class="flex justify-center">
    <div class="w-11/12 justify-self-center shadow-md sm:rounded-lg mt-4">
      <div class="flex items-center justify-end">
        <button
          :onclick="newClaim"
          type="button"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
        >
          New
        </button>
      </div>

      <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead
          class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-900 dark:text-gray-400"
        >
          <tr>
            <th scope="col" class="px-6 py-3">Code</th>
            <th scope="col" class="px-6 py-3">Uses left</th>
            <th scope="col" class="px-6 py-3">Delete</th>
          </tr>
        </thead>
        <tbody>
          <tr
            class="bg-gray-50 border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-slate-300 dark:hover:bg-gray-600"
            v-for="claim in claims"
            :key="claim.code"
          >
            <td class="px-6 py-4">{{ claim.code }}</td>
            <td class="px-6 py-4">{{ claim.use }}</td>
            <td>
              <a @click="deleteClaim(claim.code)" class="cursor-pointer"
                >Delete</a
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { deleteRequest } from "../axios/deleteRequest";
import { getRequest } from "../axios/getRequest";
import { postRequest } from "../axios/postRequest";
import { Claim } from "../models/claim";

export default {
  name: "claimManagment",
  data() {
    return {
      claims: [] as Claim[],
    };
  },
  mounted() {
    getRequest("/admin/claims/get", "json").then((r) => {
      this.claims = r.data;
    });
  },
  methods: {
    newClaim() {
      postRequest(null, "/admin/claims/create", "json").then((r) => {
        this.claims.push(r.data as Claim);
      });
    },
    deleteClaim(code: string) {
      deleteRequest("/admin/claims/delete", { code: code });
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
