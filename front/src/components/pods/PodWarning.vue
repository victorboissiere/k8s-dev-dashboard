<template>
  <div>
    <ErrorMessage :message="warningMessage" type="warning"/>
  </div>
</template>

<script>
import ErrorMessage from '../utils/ErrorMessage.vue';

export default {
  components: { ErrorMessage },
  props: {
    pods: { type: Array },
  },
  watch: {
    pods() {
      this.checkForContainerRestart();
    },
  },
  methods: {
    checkForContainerRestart() {
      let restartCount = 0;
      this.pods.forEach((pod) => {
        pod.status.containerStatuses.forEach((container) => {
          restartCount += container.restartCount;
        });
      });

      if (restartCount > 0) {
        this.warningMessage = `${restartCount} restart(s) have occurred. It may be caused by a memory exhaustion (OOMKilled) so you may need to increase memory requests/limits. It can also be failed health checks`;
      }
    },
  },
  mounted() {
    this.checkForContainerRestart();
  },
  data() {
    return {
      warningMessage: '',
    };
  },
};
</script>
