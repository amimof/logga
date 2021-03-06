<template>
  <div class="h-100">
    
    <div class="container">
      <Loader v-if="isLoading" />
      <div class="alert alert-info" role="alert" v-if="!podLog && !isLoading">
        Pod logs are no longer available
      </div>
    </div>

    <div v-if="!isLoading && podLog" class="border-top border-dark h-100" style="background-color: black">
      <div>
        <nav class="navbar sticky-top border-bottom border-dark bg" style="background-color: black">
          <div class="container">
            <div class="row align-items-center">
              <b-dropdown class="navbar-brand" :text="pod.spec.containers[selectedContainer].name" variant="outline-primary" slot="append">
                <b-dropdown-header>containers</b-dropdown-header>
                <b-dropdown-item
                  variant="secondary"
                  v-for="(container, index) in pod.spec.containers" :key="index" 
                  :active="index == selectedContainer"
                  v-on:click="setSelectedContainer(index)">
                  {{ container.name }}
                </b-dropdown-item>
              </b-dropdown>
              <span class="log-length">{{ podLog.length }}/{{ maxLines }}</span>
            </div>

            <b-button-group>
              <b-button variant="outline-primary" v-on:click="reload()" :disabled="isReloading" v-b-tooltip.hover title="Reload (R)"><i class="fas fa-redo"></i></b-button>
              <b-button variant="outline-primary" v-on:click="download(pod.metadata.name+'.log', podLog)" v-b-tooltip.hover title="Download log (D)"><i class="fas fa-download"></i></b-button>
            </b-button-group>
            <b-button-group>
              <b-button variant="outline-primary" v-on:click="toggleLargeText()" :pressed="isLargeText" v-b-tooltip.hover title="Increased text size (L)"><i class="fas fa-text-height"></i></b-button>
              <b-button variant="outline-primary" v-on:click="gotoTop()" v-b-tooltip.hover title="Go to top (T)"><i class="fas fa-arrow-up"></i></b-button>
              <b-button variant="outline-primary" v-on:click="gotoBottom()" v-b-tooltip.hover title="Go to bottom (B)"><i class="fas fa-arrow-down"></i></b-button> 
              <b-button variant="outline-primary" v-on:click="toggleWatch()" :pressed="isWatching" v-b-tooltip.hover title="Watch log (W)"><i class="fas fa-eye"></i></b-button> 
            </b-button-group>
          </div>
        </nav>
        <div class="log-view">
          <div class="container">
            <table style="width: 100%">
              <tbody>
                <tr class="log-line" 
                  v-for="(line, index) in podLog" :key="index">
                  <td class="log-line-number"  v-bind:class="{ 'log-line-large': isLargeText }">{{ index+lineStart+1 }}</td>
                  <td class="log-line-text" v-bind:class="{ 'log-line-large': isLargeText }">{{ line }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Loader from './Loader.vue'
const keyw = 87;
const keyb = 66;
const keyt = 84;
const keyl = 76;
const keyr = 82;
const keyd = 68;
const keytab = 9;
export default {
  name: 'LogViewer',
  components: {
    Loader,
  },
  data () {
    return {
      isReloading: false,
      isLoading: true,
      isError: false,
      isLargeText: false,
      isWatching: false,
      error: null,
      selectedContainer: 0,
      modifier: 0
    }
  },
  mounted() {
    window.addEventListener("keydown", this.handlePress);
    this.getLogs();
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this.handlePress);
  },
  updated() {
    this.gotoBottom();
  },
  created() {
    if(this.$route.query.container) {
      for(let i = 0; this.pod.spec.containers.length > i; i++) {
        if(this.pod.spec.containers[i].name.toLowerCase() == this.$route.query.container.toLowerCase()) {
          this.selectedContainer = i;
          break;
        } 
      }
    }
    if(this.$route.query.watch && this.$route.query.watch.toLowerCase() == 'true') {
      this.toggleWatch();
    }
    if(this.$route.query.largetext && this.$route.query.largetext.toLowerCase() == "true") {
      this.toggleLargeText(); 
    }
  },
  methods: {
    reload() {
      this.closeStream()
      this.getLogs()
      this.$store.dispatch('resetLineStart');
      this.gotoBottom();
    },
    getLogs() {
      this.isReloading = true;
      this.$store.dispatch('getPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod, container: this.pod.spec.containers[this.selectedContainer].name }).then(() => {
        this.gotoBottom();
      }).catch(err => {
        this.isError = true;
        this.error = err;
      }).finally(() => {
        this.isLoading = false;
        this.isReloading = false
      });
    },
    watch() {
      this.$store.dispatch('streamPodLog', { namespace: this.$route.params.namespace, pod: this.$route.params.pod, container: this.pod.spec.containers[this.selectedContainer].name });
      this.isWatching = true;
      this.$router.push({ query: Object.assign({}, this.$route.query, { watch: "true" })});
      this.gotoBottom();
    },
    toggleWatch() {
      if(this.isWatching) {
        this.closeStream()
      } else {
        this.watch();
      }
    },
    closeStream() {
      this.$store.dispatch('closeStream')
      this.isWatching = false;
      this.$router.push({ query: Object.assign({}, this.$route.query, { watch: "false" })});
    },
    gotoBottom () {
      window.scrollTo(0, document.body.scrollHeight);
    },
    gotoTop () {
      window.scrollTo(0, 49);
    },
    toggleLargeText() {
      this.isLargeText = !this.isLargeText;
      this.$router.push({ query: Object.assign({}, this.$route.query, { largetext: this.isLargeText })});
    },
    setSelectedContainer(index) {
      this.selectedContainer = index;
      this.$router.push({ query: Object.assign({}, this.$route.query, { container: this.pod.spec.containers[this.selectedContainer].name })});
      this.getLogs();
      if(this.isWatching) {
        this.closeStream();
        this.watch();
      }
    },
    setNextContainer() {
      if(this.pod.spec.containers.length > 1) {
        if(this.selectedContainer < (this.pod.spec.containers.length-1)) {
          this.setSelectedContainer(this.selectedContainer+1);
          return
        }
        if(this.selectedContainer == (this.pod.spec.containers.length-1)) {
          this.setSelectedContainer(0);
          return
        }
      }
    },
    handlePress(event) {
      switch(event.keyCode) {
        case keyw:
          this.toggleWatch();
          break;
        case keyt:
          this.gotoTop();
          break;  
        case keyb:
          this.gotoBottom();
          break;
        case keyl:
          this.toggleLargeText();
          break;
        case keyr:
          this.reload();
          break;
        case keyd:
          this.download(`${this.pod.metadata.name}.log`, this.podLog);
          break;
        case keytab: 
          this.setNextContainer();
          event.preventDefault();
          break;
      }
    },
    download(filename, data) {
      const content = btoa(data.join('\r\n'));
      const link = document.createElement('a');
      link.setAttribute('download', filename); //or any other extension
      link.setAttribute("href", "data:text/plain;base64;charset=utf-8,"+content); //or any other extension
      document.body.appendChild(link);
      link.click();
    }
  },
  computed: {
    ...mapState([
      'pod',
      'podLog',
      'maxLines',
      'lineStart'
    ])
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
  line-height: 20px;
  word-break: break-word;
  overflow-wrap: break-word;
  min-width: 0;
  word-wrap: break-word;
  color: rgb(209, 209, 209);
}
.log-line-large {
  font-size: 14px;
}
.log-line {
  font-size: 12px;
}
.log-length {
  mix-blend-mode: difference;
  font-family: Menlo,Monaco,Consolas,monospace;
  color: rgb(108, 117, 125);
}
</style>


