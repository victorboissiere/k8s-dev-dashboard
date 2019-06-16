<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <Request path="tracker" ref="request">
      <template v-slot:default="tracker">
        <div v-if="tracker.data">

          <v-layout row>
            <v-flex xs12 sm12>
              <v-card>
                <v-toolbar color="indigo" dark>
                  <v-toolbar-title>Tracker</v-toolbar-title>
                  <v-spacer></v-spacer>
                  <v-btn icon>
                    <v-icon @click="refresh">refresh</v-icon>
                  </v-btn>
                </v-toolbar>
                <v-list>
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

export default {
  components: { Request },
  methods: {
    refresh() {
      this.$refs.request.refresh();
    },
  },
};
</script>
