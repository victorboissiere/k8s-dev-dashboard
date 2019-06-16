<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <ArbitraryRollingUpdate ref="arbitraryRollingUpdate"/>
    <Request
      :path="`namespaces/${namespace}/deployments/${deployment}/pods`"
      ref="request"
    >
      <template v-slot:default="pods">
        <div v-if="pods.data">
        <PodWarning :pods="pods.data"/>
        <v-toolbar color="indigo" dark>
          <v-toolbar-title>
            Pods for deployment {{ deployment }} ({{ pods.data.length }})
          </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon @click="arbitraryRollingUpdate">redo</v-icon>
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
import ArbitraryRollingUpdate from './ArbitraryRollingUpdate.vue';
import PodWarning from './pods/PodWarning.vue';
import PodStatus from './pods/PodStatus.vue';
import PodList from './pods/PodList.vue';

export default {
  components: {
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
    refresh() {
      this.$refs.request.refresh();
    },
    arbitraryRollingUpdate() {
      this.$refs.arbitraryRollingUpdate.show();
    },
  },
};
</script>
