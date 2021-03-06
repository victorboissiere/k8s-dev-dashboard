<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <Request
      :path="`namespaces/${namespace}/deployments/${deployment}/pods`"
      ref="request"
    >
      <template v-slot:default="pods">
        <div v-if="pods.data">
        <ArbitraryRollingUpdate
          ref="arbitraryRollingUpdate"
          :namespace="namespace"
          @onPatched="onPatched"
          :deployment="deployment"
        />
        <PodEvents ref="events" :namespace="namespace" :deployment="deployment"/>
        <PodWarning :pods="pods.data"/>
        <v-toolbar color="indigo" dark>
          <v-toolbar-title>
            {{ namespace}} | Pods for deployment {{ deployment }} ({{ pods.data.length }})
          </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-btn icon>
                <v-icon @click="arbitraryRollingUpdate" v-on="on">rotate_90_degrees_ccw</v-icon>
              </v-btn>
            </template>
            <span>Rolling update</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-btn icon><v-icon @click="events" v-on="on">notifications</v-icon></v-btn>
            </template>
            <span>Events</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-btn icon><v-icon @click="refresh" v-on="on">refresh</v-icon></v-btn>
            </template>
            <span>Refresh</span>
          </v-tooltip>
        </v-toolbar>
        <v-card v-for="pod in pods.data">
          <v-card-title>
            <v-icon large left>folder_open</v-icon>
            <span class="title font-weight-light">{{ pod.name }}</span>
            <v-spacer></v-spacer>
            <PodStatus :pod="pod"/>
          </v-card-title>

          <v-card-text>
            <PodList :pod="pod"/>
          </v-card-text>
        </v-card>
        </div>
      </template>
    </Request>
  </div>
</template>

<script>
import Request from './utils/Request.vue';
import ArbitraryRollingUpdate from './pods/ArbitraryRollingUpdate.vue';
import PodWarning from './pods/PodWarning.vue';
import PodStatus from './pods/PodStatus.vue';
import PodList from './pods/PodList.vue';
import PodEvents from './pods/PodEvents';

export default {
  components: {
    PodEvents,
    PodList,
    PodStatus,
    PodWarning,
    ArbitraryRollingUpdate,
    Request,
  },
  props: {
    namespace: { type: String },
    deployment: { type: String },
  },
  methods: {
    onPatched() {
      this.refresh();
    },
    refresh() {
      this.$refs.request.refresh();
    },
    events() {
      this.$refs.events.show();
    },
    arbitraryRollingUpdate() {
      this.$refs.arbitraryRollingUpdate.show();
    },
  },
};
</script>
