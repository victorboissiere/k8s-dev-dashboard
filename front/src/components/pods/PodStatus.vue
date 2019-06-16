<template>
  <div>
    <v-chip label color="orange" text-color="white" small
            v-if="pod.status.phase === 'Pending'">
      <v-icon left>autorenew</v-icon> Pending
    </v-chip>
    <v-chip label color="green" text-color="white" small
            v-if="pod.status.phase === 'Running'">
      <v-icon left>cloud_done</v-icon> Running
    </v-chip>
    <v-chip small label>
      {{ time }}<br/>
    </v-chip>
  </div>
</template>

<script>
import ta from 'time-ago';

export default {
  props: {
    pod: { type: Object },
  },
  computed: {
    time() {
      return ta.ago(this.pod.status.startTime);
    },
  },
};
</script>
