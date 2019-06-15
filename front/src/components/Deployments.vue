<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <v-card flat>
      <v-card-title>
        <v-icon large left>dashboard</v-icon>
        <span class="title font-weight-light">Deployments in {{ namespace }}</span>
      </v-card-title>
      <v-card-text>
        <Request :path="`namespaces/${namespace}/deployments`">
          <template v-slot:default="deployments">
            <v-card>
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
                      <v-chip label color="grey" text-color="white" :small="true"
                              v-if="deployment.labels && `release` in deployment.labels">
                        <v-icon left>insert_link</v-icon>Managed by Helm
                      </v-chip>
                      <v-chip label color="green" text-color="white" :small="true"
                              v-if="deployment.replicas === deployment.availableReplicas">
                        <v-icon left>cloud_done</v-icon>Running
                      </v-chip>
                      <v-chip label color="orange" text-color="white" :small="true"
                              v-else-if="deployment.replicas > deployment.availableReplicas && deployment.unvailableReplicas == 0">
                        <v-icon left>autorenew</v-icon> Deployment in progress
                      </v-chip>
                      <v-chip label color="red" text-color="white" :small="true" v-else>
                        <v-icon left>error</v-icon> Error
                      </v-chip>
                    </v-list-tile-action>
                  </v-list-tile>
                  <v-divider
                    :key="deployment.name"
                  ></v-divider>
                </template>
              </v-list>
            </v-card>
          </template>
        </Request>
      </v-card-text>
    </v-card>
</template>

<script>
import Request from './utils/Request.vue';

export default {
  components: { Request },
  props: {
    namespace: { type: String },
  },
  data() {
    return {
      active: null,
    };
  },
};
</script>
