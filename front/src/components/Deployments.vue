<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Request :path="`namespaces/${namespace}/deployments`" ref="request">
    <template v-slot:default="deployments">
      <div v-if="deployments.data">
        <v-toolbar color="indigo" dark>
          <v-toolbar-title>Deployments in {{ namespace }}</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon @click="refresh">refresh</v-icon>
          </v-btn>
        </v-toolbar>
        <v-card v-if="deployments.data.length > 0">
          <v-list two-line>
            <template v-for="(deployment, index) in deployments.data">
              <v-list-tile
                :key="`${deployment.name}-${index}`"
                :to="{ name: 'pods', params: { namespace, deployment: deployment.name } }">
                <v-list-tile-content>
                  <v-list-tile-title v-html="deployment.name"/>
                  <v-list-tile-sub-title>
                    Replicas: {{ deployment.replicas }} / {{ deployment.availableReplicas }}
                  </v-list-tile-sub-title>
                </v-list-tile-content>
                <v-list-tile-action>
                  <v-chip label color="red" text-color="white" :small="true"
                          v-if="deployment.status.hasErrors">
                    <v-icon left>error</v-icon> Error
                  </v-chip>
                  <v-chip label color="orange" text-color="white" :small="true"
                          v-else-if="deployment.status.isInProgress">
                    <v-icon left>autorenew</v-icon> Deployment in progress
                  </v-chip>
                  <v-chip label color="green" text-color="white" :small="true" v-else>
                    <v-icon left>cloud_done</v-icon>Running
                  </v-chip>
                  <v-chip label color="grey" text-color="white" :small="true"
                          v-if="deployment.labels && `release` in deployment.labels">
                    <v-icon left>insert_link</v-icon>Managed by Helm
                  </v-chip>
                </v-list-tile-action>
              </v-list-tile>
              <v-divider
                :key="deployment.name"
                v-if="index !== deployments.data.length - 1"
              ></v-divider>
            </template>
          </v-list>
        </v-card>
        <v-card v-else>
          <v-card-title>No deployments in this namespace</v-card-title>
        </v-card>
      </div>
    </template>
  </Request>
</template>

<script>
import Request from './utils/Request.vue';

export default {
  components: { Request },
  props: {
    namespace: { type: String },
  },
  methods: {
    refresh() {
      this.$refs.request.refresh();
    },
  },
  data() {
    return {
      active: null,
    };
  },
};
</script>
