<template>
  <div id="container">
    <Loader v-if="isLoading" />
    
    <nav class="navbar sticky-top navbar-dark bg-dark">
      <a class="navbar-brand" href="#">Sticky top</a>
    </nav>

    <div class="log-view">
      <span class="log-item" v-for="(line, index) in lines" :key="index">
        {{ line }}
      </span>
    </div>

    <div class="alert alert-danger" role="alert" v-if="isError">
      <h4 class="alert-heading">Oops! <span class="navbar-brand fas fa-sad-tear"></span></h4>
      <p>Unable to load pod logs</p>
      <hr/>
      <p class="mb-0"><pre>{{ errMsg }}</pre></p>
    </div>

  </div>
</template>

<script>
import { mapState } from 'vuex'
import Loader from './Loader.vue'
export default {
  name: 'LogViewer',
  components: {
    Loader
  },
  data () {
    return {
      isLoading: true,
      isError: false,
      errMsg: null,
      lines: []
    }
  }, 
  created() {
    this.$store.dispatch('getPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod }).then(() => {
      this.loading = false;
      this.getLogs()
    }).catch(err => {
      this.isError = true;
      this.errMsg = err;
    }).finally(() => {
      this.isLoading = false;
    });
  },
  methods: {
    getLogs () {
      let lines = this.podLog.split(/\r?\n/);
      this.lines = lines;
    }
  },
  computed: {
    ...mapState([
      'podLog'
    ]),
  }
}
</script>

<style lang="scss" scoped>
.log-view {
  background-color: black;
  padding: 8px;
}
.log-item {
  font-family: Menlo,Monaco,Consolas,monospace;
  font-size: 12px;
  display: block;
  color: rgb(209, 209, 209);
  line-height: 20px;
  word-break: break-word;
  overflow-wrap: break-word;
}
</style>


