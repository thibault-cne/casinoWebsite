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

      <table class="w-full text-sm text-left">
        <thead
          class="text-xs text-accent-content bg-accent uppercase border-b border-accent-content"
        >
          <tr>
            <th scope="col" class="px-6 py-3">Code</th>
            <th scope="col" class="px-6 py-3">Uses left</th>
            <th scope="col" class="px-6 py-3">Delete</th>
          </tr>
        </thead>
        <tbody class="bg-accent text-accent-content">
          <tr v-for="claim in claims" :key="claim.code">
            <td class="px-6 py-4">{{ claim.code }}</td>
            <td class="px-6 py-4">{{ claim.use }}</td>
            <td>
              <a
                @click="deleteClaim(claim.code)"
                class="text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-800 cursor-pointer"
                ><svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  class="w-5 h-5 mr-2 -ml-1"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
                  />
                </svg>
                Delete
              </a>
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
    this.fetchClaims();
  },
  methods: {
    newClaim() {
      postRequest(null, "/admin/claims/create", "json").then((r) => {
        this.claims.push(r.data as Claim);
      });
    },
    deleteClaim(code: string) {
      deleteRequest("/admin/claims/delete", { code: code }).then((_) => {
        this.fetchClaims();
      });
    },
    fetchClaims() {
      getRequest("/admin/claims/get", "json").then((r) => {
        this.claims = r.data;
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
