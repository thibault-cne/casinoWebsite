<template>
  <v-card class="list">
    <div class="row">
      <v-btn class="m-10" @click="this.newUserDialog = true">New User</v-btn>
      <NewUserDialog
        :toggle="this.newUserDialog"
        @toggle="(p) => this.togglePassDialog(p)"
      />
      <PassDialog
        :toggle="this.passToggle"
        :pass="this.newPass"
        @toggle="this.passToggle = false"
      />
    </div>
    <v-table>
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
            <v-icon class="mr icon" @click="this.toggle(u.userId)"
              >mdi-pencil
            </v-icon>
            <v-icon class="icon">mdi-delete</v-icon>
            <EditDialog
              :toggle="this.editDialog[u.userId]"
              :user="u"
              @save-user="(u) => this.saveUser(u)"
            />
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-card>
</template>
<script>
import { getRequest } from "@/axios/requests/getRequest";
import { postRequest } from "@/axios/requests/postRequest";
import EditDialog from "@/components/adminComponents/editDialog.vue";
import NewUserDialog from "@/components/adminComponents/newUserDialog.vue";
import PassDialog from "@/components/adminComponents/passDialog.vue";

export default {
  name: "AdminView",
  data() {
    return {
      users: [],
      editDialog: [],
      deleteDialog: [],
      newUserDialog: false,
      newPass: "",
      passToggle: false,
    };
  },
  created() {
    this.retriveUsers();
    for (let index = 0; index < this.editDialog.length; index++) {
      this.editDialog.push(false);
      this.deleteDialog.push(false);
    }
  },
  methods: {
    retriveUsers() {
      getRequest("/client/data/admin/all").then((r) => {
        this.users = r.data.clients;
      });
    },
    saveUser(u) {
      let data = {
        userId: u.userId,
        username: u.username,
        accessType: u.accessType,
        wallet: u.wallet,
      };
      this.editDialog[u.userId] = false;
      postRequest(data, "/client/data/admin/update");
    },
    toggle(id) {
      this.editDialog[id] = true;
    },
    togglePassDialog(p) {
      this.newUserDialog = false;
      this.newPass = p;
      this.passToggle = true;
      this.retriveUsers();
    },
  },
  components: { EditDialog, NewUserDialog, PassDialog },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/margin";

.row {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.icon {
  cursor: pointer;
}

.mr {
  margin-right: 8px;
}
</style>
