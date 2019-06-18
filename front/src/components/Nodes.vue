<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Request path="nodes">
    <template v-slot:default="nodes">
      <v-layout justify-center v-if="nodes.data">
        <v-flex xs12 sm12>
          <v-toolbar color="indigo" dark>
            <v-toolbar-title>{{ nodes.data.length }} Nodes</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>

          <v-card>
            <v-container
              fluid
              grid-list-md
            >
              <v-layout row wrap>
                <v-flex
                  v-for="node in nodes.data"
                  :key="node.name"
                  class="xs6"
                >
                  <v-card>
                    <v-card-title>
                      <v-icon left>cloud</v-icon>
                      <span class="title font-weight-light">{{ node.name }}</span>
                    </v-card-title>

                    <v-card-text>
                      <h3>Labels</h3>
                      <v-divider></v-divider><br/>
                      <div v-for="(value, key) in node.labels" :key="`${node.name}-${key}`">
                        {{ key }} : {{ value}}
                      </div>
                      <br/>
                      <h3>Specifications</h3>
                      <v-divider></v-divider><br/>
                      <p>
                        <b>Kubelet Version:</b> {{ node.status.nodeInfo.kubeletVersion }}<br/>
                        <b>OS:</b> {{ node.status.nodeInfo.osImage }}<br/>
                        <b>Kernel:</b> {{ node.status.nodeInfo.kernelVersion }}<br/>
                        <b>Runtime version:</b> {{ node.status.nodeInfo.containerRuntimeVersion }}
                        <br/>
                      </p>
                      <h3>Allocation capacity</h3>
                      <v-divider></v-divider><br/>
                      <p>
                        <b>VM: </b>{{ node.status.allocatable.cpu }} CPU //
                        {{ getHumanBytes(getAllocatableMemory(node)) }}
                        RAM<br/>
                        <b>Storage:</b>
                        {{ getHumanBytes(node.status.allocatable['ephemeral-storage']) }}<br/>
                        <b>Pods:</b> {{ node.status.allocatable.pods }}<br/>
                      </p>
                      <h3>Request</h3>
                      Memory: {{ getHumanBytes(node.request.memory) }} /
                      {{ getHumanBytes(getAllocatableMemory(node)) }}
                      <v-progress-linear
                        :color="getUsageColor(getMemoryUsage(node))"
                        height="5"
                        :value="getMemoryUsage(node)"
                      ></v-progress-linear>
                      <br/>
                      CPU: {{ (node.request.cpu / 1000).toFixed(2) }} /
                      {{ node.status.allocatable.cpu }}
                      <v-progress-linear
                        :color="getUsageColor(getCPUUsage(node))"
                        height="5"
                        :value="getCPUUsage(node)"
                      ></v-progress-linear>
                    </v-card-text>
                  </v-card>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card>
        </v-flex>
      </v-layout>
    </template>
  </Request>
</template>

<script>
import Request from './utils/Request.vue';
import units from '../mixins/units';

export default {
  components: { Request },
  mixins: [units],
  methods: {
    getAllocatableMemory(node) {
      return node.status.allocatable.memory.replace('Ki', '') * 1000;
    },
    getMemoryUsage(node) {
      return node.request.memory / this.getAllocatableMemory(node) * 100;
    },
    getCPUUsage(node) {
      return (node.request.cpu / 1000) / node.status.allocatable.cpu * 100;
    },
    getUsageColor(percent) {
      if (percent < 75) {
        return 'green';
      } if (percent < 90) {
        return 'orange';
      }

      return 'red';
    },
  },
};
</script>
