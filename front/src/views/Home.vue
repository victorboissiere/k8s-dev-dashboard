<template>
  <div class="about">
    <NamespaceApplicationList/>
    <br/><br/>
    <Tracker/>
  </div>
</template>
<script>
import NamespaceApplicationList from '../components/namespaces/NamespaceApplicationList.vue';
import Tracker from '../components/Tracker.vue';

export default {
  components: { NamespaceApplicationList, Tracker },
};
</script>
