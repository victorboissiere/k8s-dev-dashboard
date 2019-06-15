<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Request :path="`namespaces/${namespace}/deployment/${deployment}/pods`">

    <template v-slot:default="pods">
      <v-card v-for="pod in pods.data">
        <v-card-title>
          <v-icon large left>home</v-icon>
          <span class="title font-weight-light">{{ pod.name }}</span>
        </v-card-title>

        <v-card-text>
          {{ pod.status.startTime }}<br/>
          HostIP : {{ pod.hostIP }}<br/>
          PodIP : {{ pod.podIP }}<br/>
          Phase : {{ pod.status.phase }}<br/>
          Message : {{ pod.status.message }}<br/>
          Reason : {{ pod.status.reason }}<br/>
          <div v-for="container in pod.status.containerStatuses" :key="container.name">
            image: {{ container.image }} <br/>
            {{ container.name }}<br/>
            <p v-for="(value, key) in container.state" :key="`${pod.name}-${key}`">
              <b>{{ key }}</b>: {{ value.reason }}, {{ value.message }}
              <br/>
            </p>
            <br/>
            Restart: {{ container.restartCount }}
          </div>
        </v-card-text>
      </v-card>
      {{ JSON.stringify(pods.data)}}
    </template>
  </Request>
</template>

<script>
import Request from './utils/Request.vue';

export default {
  components: { Request },
  props: {
    namespace: { type: String },
    deployment: { type: String },
  },
};
</script>
