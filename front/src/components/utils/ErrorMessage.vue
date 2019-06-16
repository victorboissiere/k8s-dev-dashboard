<template>
  <div v-if="message !== ''">
    <v-alert :value="true" :type="type">
      {{ message }}
    </v-alert><br/>
  </div>
</template>

<script>
export default {
  props: {
    message: { type: String },
    type: { type: String, default: () => 'error' },
  },
};
</script>
