<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Request path="nodes">
    <template v-slot:default="nodes">
  <v-layout justify-center>
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
              v-bind="{ [`xs6`]: true }"
            >
              <v-card>
                <v-card-title>
                  <v-icon large left>home</v-icon>
                  <span class="title font-weight-light">{{ node.name }}</span>
                </v-card-title>

                <v-card-text>
                  <div v-for="(value, key) in node.labels" :key="`${node.name}-${key}`">
                    {{ key }} : {{ value}}
                  </div>
                  <br/><br/>
                  Kernel: {{ node.status.nodeInfo.kernelVersion }}<br/>
                  containerRuntimeVersion: {{ node.status.nodeInfo.containerRuntimeVersion }}<br/>
                  kubeletVersion: {{ node.status.nodeInfo.kubeletVersion }}<br/>
                  OS: {{ node.status.nodeInfo.osImage }}<br/>
                  <br/>
                  {{ node.status.allocatable.cpu }} CPU // {{ getHumanBytes(node.status.allocatable.memory.replace('Ki', '') * 1000) }} RAM<br/>
                  Storage: {{ getHumanBytes(node.status.allocatable['ephemeral-storage']) }}<br/>
                  Pods: {{ node.status.allocatable.pods }}<br/>
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

export default {
  components: { Request },
  methods: {
    getHumanBytes(bytes, decimals = 2) {
      if (bytes === 0) return '0 Bytes';

      const k = 1024;
      const dm = decimals < 0 ? 0 : decimals;
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

      const i = Math.floor(Math.log(bytes) / Math.log(k));

      return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
    },
  },
};
</script>
