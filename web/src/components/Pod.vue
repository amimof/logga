<template>
  <div class="h-100">
    <div class="container">
      <Loader v-if="isLoading" />
      <ErrorCard title="Unable to load pod" :error="error" v-if="isError"/>
    </div>
    <LogViewer v-if="!isError && !isLoading"/>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import LogViewer from './LogViewer.vue'
import ErrorCard from './ErrorCard.vue'
import Loader from './Loader.vue'
const keyesc = 27;
export default {
  name: 'Pod',
  components: {
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
  mounted() {  
    window.addEventListener("keydown", this.handlePress);
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
  beforeRouteLeave (to, from, next) {
    this.$store.dispatch('closeStream');
    next()
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this.handlePress);
  },
  methods: {
    handlePress() {
      if (event.keyCode == keyesc) {
        this.$router.push({ path: `/namespaces/${this.$route.params.namespace}/pods` })
      }
    }
  },
  computed: {
    ...mapState([
      'pod'
    ]),
  }
}
</script>

<style lang="scss" scoped>
</style>


