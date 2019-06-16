<template>
  <div>
    <p v-if="isLoading">
      <v-progress-linear :indeterminate="true" :class="`position: absolute;`"></v-progress-linear>
    </p>
    <slot v-else :data="data"></slot>
    <ErrorMessage :message="errorMessage"/>
  </div>
</template>

<script>
import request from '../../request';
import ErrorMessage from './ErrorMessage.vue';

export default {
  components: { ErrorMessage },
  props: {
    path: { type: String },
  },
  methods: {
    load() {
      request.make('get', this.path).then((response) => {
        if (response.status === 200) {
          this.isLoading = false;
          this.data = response.data;
          this.$emit('onDataReceived', response.data);
        } else {
          this.isLoading = false;
          this.errorMessage = 'Something went wrong';
        }
      }).catch(() => {
        this.isLoading = false;
        this.errorMessage = 'Something went wrong';
      });
    },
    refresh() {
      this.isLoading = true;
      this.load();
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
      errorMessage: '',
    };
  },
};
</script>

<style>
  .v-progress-linear {
    margin: 0;
  }
</style>
