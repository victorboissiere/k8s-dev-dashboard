<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <Request path="tracker" ref="request" @onDataReceived="onDataReceived">
      <template v-slot:default="tracker">
        <div v-if="tracker.data">

          <v-layout row>
            <v-flex xs6 sm6 class="pr-4">
              <v-card>
                <v-toolbar color="indigo" dark>
                  <v-toolbar-title>Deployment Tracker</v-toolbar-title>
                  <v-spacer></v-spacer>
                  <v-btn icon>
                    <v-icon @click="refresh">refresh</v-icon>
                  </v-btn>
                </v-toolbar>
                <v-list v-if="tracker.data.deployments.length > 0">
                  <v-list-tile
                    v-for="deployment in tracker.data.deployments"
                    :key="`${deployment.name}-${deployment.namespace}-tracker`"
                    :to="{ name: 'pods',
                    params: { namespace: deployment.namespace , deployment: deployment.name } }"
                  >
                    <v-list-tile-content>
                      <v-list-tile-title>
                        {{ deployment.namespace }} | Deployment {{ deployment.name }}
                      </v-list-tile-title>
                      <v-list-tile-sub-title
                        class="red--text"
                       v-if="deployment.status === 'Failing'"
                      >
                        Replicas not matching specifications
                      </v-list-tile-sub-title>
                      <v-list-tile-sub-title class="red--text" v-else>
                        Deployment is failing
                      </v-list-tile-sub-title>
                    </v-list-tile-content>
                  </v-list-tile>
                </v-list>
                <div class="text-md-center" v-else>
                  <v-card-title primary-title>
                    <div class="text-md-center">
                      Applications namespaces are healthy
                    </div>
                  </v-card-title>
                </div>
              </v-card>
            </v-flex>
            <v-flex xs6 sm6 class="pl-4">
              <v-card>
                <v-toolbar color="indigo" dark>
                  <v-toolbar-title>Cluster tracker</v-toolbar-title>
                </v-toolbar>
                  <v-card-title primary-title>
                    <div>
                      <h3 class="headline mb-0">Cluster resources</h3><br/>
                      <div>
                        <b>MEM requested:</b> {{ Math.floor(memoryUsage()) }} %<br/>
                        <b>CPU requested:</b> {{ Math.floor(cpuUsage()) }} %<br/>
                        <br/>
                        <b>Memory:</b> {{ getHumanBytes(clusterSpec.request.memory )}}
                        / {{ getHumanBytes(clusterSpec.limit.memory) }}
                        (Free:
                        {{ getHumanBytes(clusterSpec.limit.memory - clusterSpec.request.memory) }})
                        <br/>
                        <b>CPU:</b> {{ clusterSpec.request.cpu }}
                        / {{ clusterSpec.limit.cpu }}
                        (Free:
                        {{ clusterSpec.limit.cpu - clusterSpec.request.cpu }})
                        <br/>
                      </div>
                    </div>
                  </v-card-title>
              </v-card>
            </v-flex>
          </v-layout>
        </div>
      </template>
    </Request>
  </div>
</template>

<script>
import Request from './utils/Request.vue';
import units from '../mixins/units';

export default {
  components: { Request },
  mixins: [units],
  methods: {
    refresh() {
      this.$refs.request.refresh();
    },
    onDataReceived(data) {
      this.clusterSpec = data.topNodes.spec;
    },
    memoryUsage() {
      if (this.clusterSpec.limit.memory === 0) {
        return 0;
      }

      return this.clusterSpec.request.memory / this.clusterSpec.limit.memory * 100;
    },
    cpuUsage() {
      if (this.clusterSpec.limit.cpu === 0) {
        return 0;
      }

      return this.clusterSpec.request.cpu / this.clusterSpec.limit.cpu * 100;
    },
  },
  data() {
    return {
      clusterSpec: {
        request: {
          memory: 0,
          cpu: 0,
        },
        limit: {
          memory: 0,
          cpu: 0,
        },
      },
    };
  },
};
</script>
