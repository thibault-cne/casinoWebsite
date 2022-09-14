<template>
  <v-card class="list">
    <v-table fixed-header>
      <thead>
        <tr>
          <th>ID</th>
          <th>Username</th>
          <th>AccessType</th>
          <th>Wallet</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="u in this.users" :key="u.id">
          <td>{{ u.userId }}</td>
          <td>{{ u.username }}</td>
          <td>{{ u.accessType }}</td>
          <td>{{ u.wallet }}</td>
          <td>
            <v-icon class="mr icon" @click="editDialog = true"
              >mdi-pencil
            </v-icon>
            <v-icon class="icon">mdi-delete</v-icon>
            <div class="text-center">
              <v-dialog v-model="editDialog">
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
                          v-model="u.username"
                          class="mr"
                        ></v-text-field>
                        <v-select
                          :items="['1', '2', '3']"
                          label="Access Type*"
                          required
                          v-model="u.accessType"
                          class=""
                        ></v-select>
                      </v-row>
                      <v-row>
                        <v-text-field
                          label="Wallet*"
                          required
                          v-model="u.wallet"
                          class="small-field"
                        ></v-text-field>
                      </v-row>
                    </v-container>
                    <small>*indicates required field</small>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="blue-darken-1"
                      text
                      @click="editDialog = false"
                    >
                      Close
                    </v-btn>
                    <v-btn
                      color="blue-darken-1"
                      text
                      @click="editDialog = false"
                    >
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </div>
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-card>
</template>
<script>
import { getRequest } from "@/axios/requests/getRequest";

export default {
  name: "AdminView",
  data() {
    return {
      users: [],
      editDialog: false,
      deleteDialog: false,
    };
  },
  created() {
    this.retriveUsers();
  },
  methods: {
    retriveUsers() {
      getRequest("/client/data/admin/all").then((r) => {
        this.users = r.data.clients;
      });
    },
  },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/mixin";
.icon {
  cursor: pointer;
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
