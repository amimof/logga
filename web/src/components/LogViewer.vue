<template>
  <div>
    <div class="container">

      <Loader v-if="isLoading" />
      
      <div class="alert alert-danger" role="alert" v-if="isError">
        <h4 class="alert-heading">Oops! <span class="navbar-brand fas fa-sad-tear"></span></h4>
        <p>Unable to load pod logs</p>
        <hr/>
        <p class="mb-0"><pre>{{ errMsg }}</pre></p>
      </div>
      <div class="alert alert-info" role="alert" v-if="!podLog && !isLoading">
        Pod logs are no longer available
      </div>
      
    </div>

    <div v-if="!isLoading && podLog" style="background-color: black; height: 100%" class="border-top border-secondary">
      <div class="container">
        <nav class="navbar sticky-top border-bottom border-dark bg" style="background-color: black;">
          <a class="navbar-brand" href="#">Sticky top</a>
          <b-button-group>
            <b-button variant="outline-primary" v-on:click="reload()" :disabled="isReloading" v-b-tooltip.hover title="Reload"><i class="fas fa-redo"></i></b-button>
            <b-button variant="outline-primary" v-b-tooltip.hover title="Download log to file"><i class="fas fa-download"></i></b-button>
          </b-button-group>
          <b-button-group>
            <b-button variant="outline-primary" v-on:click="toggleLargeText()" :pressed="isLargeText" v-b-tooltip.hover title="Increased text size"><i class="fas fa-text-height"></i></b-button>
            <b-button variant="outline-primary" v-on:click="gotoTop()" v-b-tooltip.hover title="Go to top"><i class="fas fa-arrow-up"></i></b-button>
            <b-button variant="outline-primary" v-on:click="gotoBottom()" v-b-tooltip.hover title="Go to bottom"><i class="fas fa-arrow-down"></i></b-button> 
            <b-button variant="outline-primary" v-on:click="watch()" :pressed="isWatching" v-b-tooltip.hover title="Tail log"><i class="fas fa-eye"></i></b-button> 
          </b-button-group>
        </nav>
        <div class="log-view">
          <table style="width: 100%">
            <tbody>
              <tr class="log-line" v-for="(line, index) in podLog" :key="index">
                <td class="log-line-number" v-bind:class="{ 'log-line-large': isLargeText }">{{ index }}</td>
                <td class="log-line-text" v-bind:class="{ 'log-line-large': isLargeText }">{{ line }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
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
      isReloading: false,
      isLoading: true,
      isError: false,
      isLargeText: false,
      isWatching: false,
      errMsg: null,
    }
  }, 
  created() {
    this.getLogs()
  },
  methods: {
    reload() {
      this.closeStream()
      this.getLogs()
    },
    getLogs () {
      this.isReloading = true;
      this.$store.dispatch('getPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod }).then(() => {
        this.gotoBottom();
      }).catch(err => {
        this.isError = true;
        this.errMsg = err;
      }).finally(() => {
        this.isLoading = false;
        this.isReloading = false
      });
    },
    watch() {
      this.gotoBottom();
      if(this.isWatching) {
        this.closeStream()
        this.isWatching = false;
      } else {
        this.$store.dispatch('streamPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod });
        this.isWatching = true;
      }
    },
    closeStream() {
      this.$store.dispatch('closeStream')
    },
    gotoBottom () {
      window.scrollTo(0, document.body.scrollHeight);
    },
    gotoTop () {
      window.scrollTo(0, 109);
    },
    toggleLargeText() {
      this.isLargeText = !this.isLargeText;
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
  padding-top: 16px;
  font-family: Menlo,Monaco,Consolas,monospace;
}
.log-view table {
  table-layout: fixed;
  width: 100%;
}
.log-line:hover {
  background-color: #22262b;
  color: #ededed;
}
.log-line-number {
  border-right: 1px #272b30 solid;
  padding-right: 10px;
  vertical-align: top;
  white-space: nowrap;
  width: 60px;
  color: #72767b;
  text-align: right;
}
.log-line-text {
  padding: 0 10px;
  white-space: pre-wrap;
  width: 100%;
  color: rgb(209, 209, 209);
  line-height: 20px;
  word-break: break-word;
  overflow-wrap: break-word;
  min-width: 0;
  word-wrap: break-word;
}
.log-line-large {
  font-size: 14px;
}
.log-line {
  font-size: 12px;
}
</style>


