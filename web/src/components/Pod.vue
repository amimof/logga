<template>
  <div>
    <div class="container">
      <Breadcrumb/>
      <Loader v-if="isLoading" />
      <h4 v-if="pod" >
        {{ pod.metadata.name }}
      </h4>

      <ErrorCard title="Unable to load pod" :error="error" v-if="isError"/>
    </div>
    <LogViewer v-if="!isError && !isLoading"/>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
import LogViewer from './LogViewer.vue'
import ErrorCard from './ErrorCard.vue'
import Loader from './Loader.vue'
export default {
  name: 'Pod',
  components: {
    Breadcrumb,
    LogViewer,
    ErrorCard,
    Loader
  },
  data () {
    return {
      isLoading: true,
      isError: false,
      error: null
    }
  }, 
  created() {
    this.$store.dispatch('getPod', { namespace: this.$route.params.namespace, pod: this.$route.params.pod }).then(() => {
      this.isLoading = false
    }).catch(err => {
      this.isError = true
      this.error = err 
    }).finally(() => {
      this.isLoading = false
    })
    this.$store.dispatch('addRecentNamespace', this.$route.params.namespace)
  },
  computed: {
    ...mapState([
      'pod'
    ]),
  },
  beforeRouteLeave (to, from, next) {
    this.$store.dispatch('closeStream');
    next()
  }
}
</script>

<style lang="scss" scoped>
h4 {
  padding-bottom: 16px;
}
</style>


