<template>
  <div>
    <p v-if="isLoading">
      <v-progress-linear :indeterminate="true" :class="`position: absolute;`"></v-progress-linear>
    </p>
    <slot v-else :data="data"></slot>
  </div>
</template>

<script>
import request from '../../request';

export default {
  props: {
    path: { type: String },
  },
  methods: {
    load() {
      request.make('get', this.path).then((response) => {
        if (response.status === 200) {
          this.isLoading = false;
          this.data = response.data;
        } else {
          this.isLoading = false;
          console.error('Error', response); // TODO: Implement front error
        }
      });
    },
  },
  watch: {
    path() {
      this.load();
    },
  },
  mounted() {
    this.load();
  },
  data() {
    return {
      isLoading: true,
      data: null,
    };
  },
};
</script>

<style>
  .v-progress-linear {
    margin: 0;
  }
</style>
