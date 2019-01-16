<template>
  <div id="container">
    <Breadcrumb/>
    <h4 v-if="pod">
      {{ pod.metadata.name }}
    </h4>
    {{ podLog }}
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
export default {
  name: 'Pod',
  components: {
    Breadcrumb
  },
  created() {
    this.$store.dispatch('getPod', { namespace: this.$route.params.namespace, pod: this.$route.params.pod })
    this.$store.dispatch('getPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod })
  },
  mounted () {
    this.$route.meta.breadcrumb[2].name = this.$route.params.pod
  },
  computed: {
    ...mapState([
      'pod',
      'podLog'
    ]),
  }
}
</script>

<style lang="scss" scoped>
</style>


