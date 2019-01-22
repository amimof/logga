<template>
  <div>
    <div class="container">
      <Breadcrumb/>
      <h4 v-if="pod" >
        {{ pod.metadata.name }}
      </h4>
    </div>
    <LogViewer/>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
import LogViewer from './LogViewer.vue'
export default {
  name: 'Pod',
  components: {
    Breadcrumb,
    LogViewer
  },
  data () {
    return {
      loading: true
    }
  }, 
  created() {
    this.$store.dispatch('getPod', { namespace: this.$route.params.namespace, pod: this.$route.params.pod })
    
  },
  mounted () {
    this.$route.meta.breadcrumb[2].name = this.$route.params.pod
  },
  computed: {
    ...mapState([
      'pod'
    ]),
  }
}
</script>

<style lang="scss" scoped>
h4 {
  padding-bottom: 16px;
}
</style>


