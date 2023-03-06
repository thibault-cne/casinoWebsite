<template>
  <div class="flex items-center justify-center w-full p-4 md:inset-0">
    <div class="relative w-full h-full max-w-2xl md:h-auto">
      <div class="flex items-start justify-between p-4">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
          Edit user
        </h3>
      </div>
      <div class="p-6 space-y-6">
        <div class="grid grid-cols-6 gap-6">
          <div class="col-span-6 sm:col-span-3">
            <label
              for="username"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Username</label
            >
            <input
              type="text"
              name="username"
              id="username"
              class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Vladimir Petrienko"
              v-model="user.username"
            />
          </div>
        </div>
      </div>
      <div class="p-6 space-y-6">
        <div class="grid grid-cols-6 gap-6">
          <div class="col-span-6 sm:col-span-3">
            <label
              for="picture"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Picture</label
            >
            <div class="avatar">
              <div class="w-14 rounded-full">
                <img
                  :src="
                    base_backend_url +
                    '/user/get/picture/' +
                    $props.userProps.id
                  "
                />
              </div>
            </div>
            <input
              class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
              aria-describedby="file_input_help"
              id="picture"
              type="file"
              accept="image/*"
              @change="handleFileUpdate($event)"
            />
            <p
              class="mt-1 text-sm text-gray-500 dark:text-gray-300"
              id="file_input_help"
            >
              SVG, PNG, JPG or GIF (MAX. 800x400px).
            </p>
          </div>
        </div>
      </div>
      <!-- Modal footer -->
      <div class="flex items-center p-6 space-x-2">
        <button
          type="button"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          :onclick="saveEdit"
        >
          Save all
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { base_backend_url } from "../axios/axios";
import { postRequest } from "../axios/postRequest";
import { User } from "../models/user";

export default {
  name: "userSettings",
  props: {
    userProps: { required: true, type: Object as () => User },
  },
  watch: {
    userProps: function (newVal) {
      this.user = newVal;
    },
  },
  data() {
    return {
      base_backend_url: base_backend_url,
      user: {} as User,
      file: {} as File,
    };
  },
  methods: {
    handleFileUpdate(event: Event) {
      const target = event.target as HTMLInputElement;
      this.file = (target.files as FileList)[0];
    },
    saveEdit() {
      var data = new FormData();
      data.append("user", JSON.stringify(this.user));
      data.append("picture", this.file);
      postRequest(data, "/user/modify/", "file")
        .then((r) => {
          if (r.status === 200) {
            console.log("emit");
            this.$emit("update", this.user);
          }
        })
        .catch(() => {});
    },
  },
  mounted() {
    this.user = this.$props.userProps;
  },
};
</script>

<style lang="scss"></style>
