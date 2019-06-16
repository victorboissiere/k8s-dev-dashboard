<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <div v-if="pod.status.message !== '' || pod.status.reason !== ''">
      {{ pod.status.message }}  {{ pod.status.reason }}
    </div>
    <div v-for="container in pod.status.containerStatuses" :key="container.name">
      <h3>{{ container.name }} </h3>
      <div v-if="!container.ready">
        <v-chip label color="red" text-color="white" small
                v-if="true">
          <v-icon left>error</v-icon> Not ready
        </v-chip><br/>
      </div>
      <b>Image:</b> {{ container.image }}<br/>
      <v-tooltip bottom v-if="container.restartCount > 0">
        <template v-slot:activator="{ on }">
                  <span v-on="on" class="orange--text text-md-center">
                    {{ container.restartCount }} restart(s)
                  </span>
        </template>
        <span>Memory exhaustion or healthcheck fail. Check logs or Grafana</span>
      </v-tooltip>
      <div v-for="(value, key) in container.state" :key="`${pod.name}-${key}`">
        <div v-if="key !== 'running'">
          <p class="red--text text-md-center">
            <br/>
            {{ key }} {{ value.reason }}, {{ value.message }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    pod: { type: Object },
  },
};
</script>
