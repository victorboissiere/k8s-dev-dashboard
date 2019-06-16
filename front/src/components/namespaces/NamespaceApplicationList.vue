<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <request path="namespaces/applications">
      <template v-slot:default="namespaces">
        <div v-if="namespaces.data">

          <v-layout row>
            <v-flex xs12 sm12>
              <v-card>
                <v-toolbar color="indigo" dark>
                  <v-toolbar-title>applications namespaces</v-toolbar-title>
                  <v-spacer></v-spacer>
                </v-toolbar>
                <v-list>
                  <v-list-tile
                    v-for="namespace in namespaces.data"
                    :key="`${namespace}-application`"
                    :to="{ name: 'deployments', params: { namespace } }"
                  >
                    <v-list-tile-content>
                      <v-list-tile-title v-text="namespace"></v-list-tile-title>
                    </v-list-tile-content>
                  </v-list-tile>
                </v-list>
              </v-card>
            </v-flex>
          </v-layout>
        </div>
      </template>
    </request>
  </div>
</template>

<script>
import Request from '../utils/Request.vue';

export default {
  components: { Request },
};
</script>
