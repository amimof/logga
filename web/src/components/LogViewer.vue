<template>
  <div class="h-100">
    
    <div class="container">
      <Loader v-if="isLoading" />
      <div class="alert alert-info" role="alert" v-if="!podLog && !isLoading">
        Pod logs are no longer available
      </div>
    </div>

    <div v-if="!isLoading && podLog" :style="{'background-color': variant, 'border-light': theme == 'light', 'border-dark': theme == 'dark'}" class="border-top h-100">
      <div>
        <nav class="navbar sticky-top border-bottom bg" :class="{'border-light': theme == 'light', 'border-dark': theme == 'dark'}" :style="{'background-color': variant}">
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
              <b-button variant="outline-primary" v-on:click="reload()" :disabled="isReloading" v-b-tooltip.hover title="Reload"><i class="fas fa-redo"></i></b-button>
              <b-button variant="outline-primary" v-b-tooltip.hover title="Download log to file"><i class="fas fa-download"></i></b-button>
            </b-button-group>
            <b-button-group>
              <b-button variant="outline-primary" v-on:click="toggleLargeText()" :pressed="isLargeText" v-b-tooltip.hover title="Increased text size"><i class="fas fa-text-height"></i></b-button>
              <b-button variant="outline-primary" v-on:click="gotoTop()" v-b-tooltip.hover title="Go to top"><i class="fas fa-arrow-up"></i></b-button>
              <b-button variant="outline-primary" v-on:click="gotoBottom()" v-b-tooltip.hover title="Go to bottom"><i class="fas fa-arrow-down"></i></b-button> 
              <b-button variant="outline-primary" v-on:click="toggleWatch()" :pressed="isWatching" v-b-tooltip.hover title="Tail log"><i class="fas fa-eye"></i></b-button> 
            </b-button-group>
          </div>
        </nav>
        <div class="log-view" :style="{'background-color': variant}">
          <div class="container">
            <table style="width: 100%">
              <tbody>
                <tr class="log-line" 
                  v-bind:class="{'log-line-light': theme == 'light', 'log-line-dark': theme == 'dark'}" 
                  v-for="(line, index) in podLog" :key="index">
                  <td class="log-line-number"  v-bind:class="{ 'log-line-large': isLargeText, 'log-line-number-light': theme == 'light', 'log-line-number-dark': theme == 'dark' }">{{ index+lineStart+1 }}</td>
                  <td class="log-line-text" v-bind:class="{ 'log-line-large': isLargeText, 'log-line-text-light': theme == 'light', 'log-line-text-dark': theme == 'dark' }">{{ line }}</td>
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
    this.getLogs()
    this.selectedContainer = 0;
  },
  updated() {
    this.gotoBottom();
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
    },
    gotoBottom () {
      window.scrollTo(0, document.body.scrollHeight);
    },
    gotoTop () {
      window.scrollTo(0, 49);
    },
    toggleLargeText() {
      this.isLargeText = !this.isLargeText;
    },
    setSelectedContainer(index) {
      this.selectedContainer = index;
      this.getLogs();
      if(this.isWatching) {
        this.closeStream();
        this.watch();
      }
    }

  },
  computed: {
    ...mapState([
      'pod',
      'podLog',
      'theme',
      'maxLines',
      'lineStart'
    ]),
    variant() {
      if(this.theme == 'dark') {
        return 'black'
      } else {
        return 'white'
      }
    }
  }
}
</script>

<style lang="scss">
.log-view {
  padding-top: 16px;
  font-family: Menlo,Monaco,Consolas,monospace;
}
.log-view table {
  table-layout: fixed;
  width: 100%;
}
.log-line-dark:hover {
  background-color: #22262b;
  color: #ededed;
}
.log-line-light:hover {
  background-color: #ededed;
  color: #22262b;
}
.log-line-number {
  padding-right: 10px;
  vertical-align: top;
  white-space: nowrap;
  width: 60px;
  color: #72767b;
  text-align: right;
}
.log-line-number-dark {
  border-right: 1px #272b30 solid;
}
.log-line-number-light {
  border-right: 1px #ededed solid;
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
}
.log-line-text-dark {
  color: rgb(209, 209, 209);
}
.log-line-text-light {
  color: #272b30;
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


