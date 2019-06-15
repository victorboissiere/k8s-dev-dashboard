<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <v-navigation-drawer
    permanent
    app
    enable-resize-watcher
  >
    <v-list>
      <v-list-tile :to="{ name: 'home' }">
        <v-list-tile-action>
          <v-icon>home</v-icon>
        </v-list-tile-action>

        <v-list-tile-content>
          <v-list-tile-title>Home</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>

      <v-list-tile :to="{ name: 'nodes' }">
        <v-list-tile-action>
          <v-icon>view_column</v-icon>
        </v-list-tile-action>

        <v-list-tile-content>
          <v-list-tile-title>Nodes</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>

      <v-list-group prepend-icon="folder" no-action v-if="namespaces.length > 0">
        <template v-slot:activator>
          <v-list-tile>
            <v-list-tile-content>
              <v-list-tile-title>Deployments</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </template>

        <v-list-tile
          v-for="namespace in namespaces"
          :key="namespace"
          :to="{ name: 'deployments', params: { namespace } }"
        >
          <v-list-tile-content>
            <v-list-tile-title>{{ namespace }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list-group>
    </v-list>
  </v-navigation-drawer>
</template>


<script>
import request from '../request';

export default {
  mounted() {
    request.make('get', 'namespaces').then((response) => {
      this.namespaces = response.data;
    }).catch((error) => {
      console.log('Error fetching namespaces', error); // TODO: gui
    });
  },
  data() {
    return {
      namespaces: [],
    };
  },
};
</script>

