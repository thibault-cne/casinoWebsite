<template>
  <div class="text-center">
    <v-dialog v-model="t">
      <v-card class="dialog">
        <v-card-title>
          <span class="text-h5">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-text-field
                label="Username*"
                required
                v-model="this.u.username"
                class="mr"
              ></v-text-field>
              <v-select
                :items="['1', '2', '3']"
                label="Access Type*"
                required
                v-model="this.u.accessType"
              ></v-select>
            </v-row>
            <v-row>
              <v-text-field
                label="Wallet*"
                required
                v-model="this.u.wallet"
                type="number"
                class="small-field"
              ></v-text-field>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue-darken-1" text @click="this.saveUser()">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
export default {
  name: "editDialog",
  props: {
    user: undefined,
    toggle: Boolean,
  },
  watch: {
    toggle: function () {
      this.t = this.toggle;
    },
  },
  data() {
    return {
      u: this.user,
      t: this.toggle,
    };
  },
  methods: {
    saveUser() {
      this.$emit("saveUser", this.user);
    },
  },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/mixin";

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
