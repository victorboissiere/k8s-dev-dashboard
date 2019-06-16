<template>
  <div>
    <ErrorMessage :message="errorMessage"/>
    <v-dialog v-model="isDialogLoading" hide-overlay persistent width="300">
      <v-card color="primary" dark>
        <v-card-text>
          Please stand by
          <v-progress-linear indeterminate color="white" class="mb-0"></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialog" max-width="600">
      <v-card>
        <v-card-title class="headline">Arbitrary rolling-update</v-card-title>
        <v-card-text>
          This action will try to match the parent deployment specification and trigger a
          rolling-update by adding the label <pre>LAST_FORCED_ROLLING_UPDATE</pre>
          <br/><br/>
          Kubernetes will then launch new pod(s) with a rolling-update strategy.
          It can be considered as a way to restart without any down time.
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat="flat" @click="hide">
            Cancel
          </v-btn>
          <v-btn color="green darken-1" flat="flat" @click="arbitraryRollingUpdate">
            Proceed
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import request from '../request';
import ErrorMessage from './utils/ErrorMessage.vue';

export default {
  components: { ErrorMessage },
  methods: {
    arbitraryRollingUpdate() {
      this.dialog = false;
      this.isDialogLoading = true;
      request.make('patch', `namespaces/${this.namespace}/deployments/${this.deployment}/arbitrary-rolling-update`).then(() => {
        setTimeout(() => {
          this.refresh();
          this.isDialogLoading = false;
        }, 1000);
      }).catch((error) => {
        this.isDialogLoading = false;
        if (error.response.status === 403) {
          this.errorMessage = 'You are not authorized to perform this action';
        } else {
          this.errorMessage = 'Something went wrong with the server';
        }
      });
    },
    show() {
      this.dialog = true;
    },
    hide() {
      this.dialog = false;
    },
  },
  data() {
    return {
      isDialogLoading: false,
      dialog: false,
      errorMessage: '',
    };
  },
};
</script>
