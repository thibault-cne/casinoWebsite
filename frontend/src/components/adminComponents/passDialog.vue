<template>
  <div class="text-center">
    <v-dialog v-model="t">
      <v-card class="dialog">
        <v-card-title>
          <span class="text-h5">New users pass</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <span>{{ this.pass }}</span>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue-darken-1" text @click="this.$emit('toggle')">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
export default {
  name: "passDialog",
  props: {
    toggle: Boolean,
    pass: String,
  },
  watch: {
    toggle: function () {
      this.t = this.toggle;
    },
  },
  data() {
    return {
      t: false,
    };
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
