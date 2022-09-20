<template>
  <div class="text-center">
    <v-dialog v-model="dialogToggle">
      <v-alert v-model="this.alert" type="warning" closable>
        {{ this.alertMessage }}
      </v-alert>
      <v-progress-circular
        v-if="this.spinner"
        class="spinner"
        indeterminate
        color="primary"
        :width="6"
        :size="75"
      ></v-progress-circular>
      <v-card class="dialog" :class="{ blur: this.spinner }">
        <v-card-title>
          <span class="text-h5">Create a new user</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-text-field
                label="Username*"
                required
                v-model="this.newUser.username"
                class="mr"
              ></v-text-field>
              <v-select
                :items="['1', '2', '3']"
                label="Access Type*"
                required
                v-model="this.newUser.accessType"
              ></v-select>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue-darken-1" text @click="this.createNewUser()">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
import { postRequest } from "@/axios/requests/postRequest";

export default {
  name: "newUsetDialog",
  props: {
    toggle: Boolean,
  },
  watch: {
    toggle: function () {
      this.dialogToggle = this.toggle;
    },
  },
  data() {
    return {
      dialogToggle: false,
      newUser: {
        username: "",
        accessType: 1,
      },
      alert: false,
      spinner: false,
      alertMessage: "",
    };
  },
  methods: {
    createNewUser() {
      if (this.newUser.username.length <= 0) {
        this.alertMessage = "Please make sure to enter a username.";
        this.alert = true;
        return;
      }

      let data = {
        username: this.newUser.username,
        accessType: this.newUser.accessType,
      };
      this.spinner = true;
      postRequest(data, "/client/data/admin/create")
        .then((r) => {
          this.$emit("toggle", r.data.pass);
        })
        .catch((e) => {
          if (e.response) {
            if (e.response.status === 400) {
              this.alertMessage = e.response.data.error;
              this.alert = true;
            }
          }
        });
      this.spinner = false;
    },
    notEmpty(v) {
      console.log(v.length);
      return v.length > 0 || "Make sure to enter a username";
    },
  },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/mixin";

.spinner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
}

.blur {
  opacity: 60%;
}

.mr {
  margin-right: 8px;
}

.dialog {
  min-width: 100vw;

  @include mixin.lg {
    min-width: 45vw;
  }
}
</style>
