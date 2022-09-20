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
                label="Wallet*"
                required
                v-model="this.u.wallet"
                type="number"
                class="small-field"
              ></v-text-field>
            </v-row>
            <v-row class="buttons">
              <v-btn-toggle tile group>
                <v-btn @click="this.addAmount(5)">+5</v-btn>
                <v-btn @click="this.addAmount(25)">+25</v-btn>
                <v-btn @click="this.addAmount(100)">+100</v-btn>
                <v-btn @click="this.multiplyAmount(0.5)">1/2</v-btn>
                <v-btn @click="this.multiplyAmount(2)">x2</v-btn>
              </v-btn-toggle>
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
  name: "addDialog",
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
    addAmount(value) {
      this.u.wallet += value;
    },
    multiplyAmount(value) {
      this.u.wallet = Math.floor(this.u.wallet * value);
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

.buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
}
</style>
