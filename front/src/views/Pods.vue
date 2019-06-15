<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Pods :deployment="$route.params.deployment" :namespace="$route.params.namespace"/>
</template>

<script>
import Pods from '../components/Pods.vue';

export default {
  components: {
    Pods,
  },
};
</script>
