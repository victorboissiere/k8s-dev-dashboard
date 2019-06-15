<template>
  <v-app>
    <Navigation/>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>K8S</span>
        <span class="font-weight-light">Developer dashboard</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>

    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Navigation from './views/Navigation.vue';

export default {
  name: 'App',
  components: { Navigation },
  data() {
    return {};
  },
};
</script>
