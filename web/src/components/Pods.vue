<template>
  <div id="container">

    <Breadcrumb/>

    <div class="d-flex">
      <h3 class="p-2">{{ signalChange().length }}</h3>
      <div class="p-2 input-group mb-3">
        <input type="text" v-model="searchString" @input="signalChange" class="form-control" placeholder="Search Pods" aria-label="Search Namespace" aria-describedby="search-button">
        <div class="input-group-append">
          <button class="btn btn-outline-primary" type="button" id="search-button">Search</button>
        </div>
      </div>
    </div>

    <div class="list-group">      
      <a v-bind:href="'#/namespaces/'+ns.metadata.namespace+'/pods/'+ns.metadata.name" class="list-group-item list-group-item-action" v-for="(ns, index) in signalChange()" :key="index">  
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1">{{ ns.metadata.name }}</h5>
          <small>{{ ns.status.phase }}</small>
          <small>{{ ns | numContainersReady }}/{{ ns.status.containerStatuses.length }}</small>
          <small>{{ ns | numContainerRestarts }}</small>
          <small>{{ ns.metadata.creationTimestamp }}</small>
        </div>
      </a>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
export default {
  name: 'Pods',
  components: {
    Breadcrumb
  },
  data() {
    return {
      searchString: ''
    }
  },
  mounted() {
    this.$store.dispatch('getPods', this.$route.params.id)
  },
  computed: {
    ...mapState([
      'pods'
    ]),
    ...mapGetters([
      'filterPods'
    ])
  },
  filters: {
    numContainersReady: function(pod) {
      if(!pod) {
        return 0;
      }
      var numReady = 0;
      for(let i = 0; i < pod.status.containerStatuses.length; i++) {
        if(pod.status.containerStatuses[i].ready) {
          numReady++;
        }
      }
      return numReady;
    },
    numContainerRestarts: function(pod) {
      if(!pod) {
        return 0;
      }
      var numRestarts = 0;
      for(let i = 0; i < pod.status.containerStatuses.length; i++)  {
        numRestarts += pod.status.containerStatuses[i].restartCount;
      }
      return numRestarts;
    }
  },
  methods: {
    signalChange() {
      let result = this.pods.items || [];
      if(this.searchString.length > 0) {
        result = this.filterPods(this.searchString);
      } 
      return result;
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


