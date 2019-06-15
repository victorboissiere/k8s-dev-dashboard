<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <Deployments :namespace="$route.params.namespace"/>
</template>

<script>
import Deployments from '../components/Deployments.vue';

export default {
  components: {
    Deployments,
  },
};
</script>
