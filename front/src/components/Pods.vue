<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <div v-if="warningMessage !== ''">
      <v-alert :value="true" type="warning">
        {{ warningMessage }}
      </v-alert><br/>
    </div>
    <Request
      :path="`namespaces/${namespace}/deployment/${deployment}/pods`"
      ref="request"
      @onDataReceived="onDataReceived"
    >
      <template v-slot:default="pods">
        <v-toolbar color="indigo" dark>
          <v-toolbar-title>
            Pods for deployment {{ deployment }} ({{ pods.data.length }})
          </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon @click="dialog = true">redo</v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon @click="refresh">refresh</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card v-for="pod in pods.data">
          <v-card-title>
            <v-icon large left>navigate_next</v-icon>
            <span class="title font-weight-light">{{ pod.name }}</span>
            <v-spacer></v-spacer>
            <v-chip label color="orange" text-color="white" small
                    v-if="pod.status.phase === 'Pending'">
              <v-icon left>autorenew</v-icon> Pending
            </v-chip>
            <v-chip label color="green" text-color="white" small
                    v-if="pod.status.phase === 'Running'">
              <v-icon left>cloud_done</v-icon> Running
            </v-chip>
            <v-chip small label>
              {{ prettyDatetime(pod.status.startTime) }}<br/>
            </v-chip>
          </v-card-title>

          <v-card-text>
            <div v-if="pod.status.message !== '' || pod.status.reason !== ''">
              {{ pod.status.message }}  {{ pod.status.reason }}
            </div>
            <div v-for="container in pod.status.containerStatuses" :key="container.name">
              <h3>{{ container.name }} </h3>
              <div v-if="!container.ready">
                <v-chip label color="red" text-color="white" small
                        v-if="true">
                  <v-icon left>error</v-icon> Not ready
                </v-chip><br/>
              </div>
              <b>Image:</b> {{ container.image }}<br/>
              <v-tooltip bottom v-if="container.restartCount > 0">
                <template v-slot:activator="{ on }">
                  <span v-on="on" class="orange--text text-md-center">
                    {{ container.restartCount }} restart(s)
                  </span>
                </template>
                <span>Memory exhaustion or healthcheck fail. Check logs or Grafana</span>
              </v-tooltip>
              <div v-for="(value, key) in container.state" :key="`${pod.name}-${key}`">
                <div v-if="key !== 'running'">
                  <p class="red--text text-md-center">
                    <br/>
                    {{ key }} {{ value.reason }}, {{ value.message }}
                  </p>
                </div>
              </div>
            </div>
          </v-card-text>
        </v-card>
      </template>
    </Request>

    <v-dialog v-model="isDialogLoading" hide-overlay persistent width="300">
      <v-card color="primary" dark>
        <v-card-text>
          Please stand by
          <v-progress-linear
            indeterminate
            color="white"
            class="mb-0"
          ></v-progress-linear>
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
          <v-btn
            color="green darken-1"
            flat="flat"
            @click="dialog = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="green darken-1"
            flat="flat"
            @click="arbitraryRollingUpdate"
          >
            Proceed
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import ta from 'time-ago';
import Request from './utils/Request.vue';
import request from '../request';

export default {
  components: { Request },
  props: {
    namespace: { type: String },
    deployment: { type: String },
  },
  methods: {
    prettyDatetime(datetime) {
      return ta.ago(datetime);
    },
    refresh() {
      this.$refs.request.refresh();
    },
    checkForContainerRestart(pods) {
      let restartCount = 0;
      pods.forEach((pod) => {
        pod.status.containerStatuses.forEach((container) => {
          restartCount += container.restartCount;
        });
      });

      if (restartCount > 0) {
        this.warningMessage = `${restartCount} restart(s) have occurred. It may be caused by a memory exhaustion (OOMKilled) so you may need to increase memory requests/limits. It can also be failed health checks`;
      }
    },
    onDataReceived(pods) {
      this.checkForContainerRestart(pods);
    },
    arbitraryRollingUpdate() {
      this.dialog = false;
      this.isDialogLoading = true;
      request.make('patch', `namespaces/${this.namespace}/deployment/${this.deployment}/arbitrary-rolling-update`).then(() => {
        setTimeout(() => {
          this.refresh();
          this.isDialogLoading = false;
        }, 1000);
      }).catch((error) => {
        console.error(error); // TODO: gui
      });
    },
  },
  data() {
    return {
      warningMessage: '',
      dialog: false,
      isDialogLoading: false,
    };
  },
};
</script>
