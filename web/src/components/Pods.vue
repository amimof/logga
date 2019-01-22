<template>
  <div class="container">

    <breadcrumb/>

    <b-input-group :prepend="signalChange().length.toString()">
      <b-form-input v-model="searchString" @input="signalChange" placeholder="Search Pods"></b-form-input>
      <b-input-group-append>
        <b-button v-on:click="searchString = ''" variant="outline-primary">
          <i class="fas fa-times"></i>
        </b-button>
        <b-button v-on:click="isSortDown = !isSortDown" variant="outline-primary">
          <i class="fas fa-sort-amount-down" v-if="isSortDown"></i>
          <i class="fas fa-sort-amount-up" v-if="!isSortDown"></i>
        </b-button>
      </b-input-group-append>
      <b-dropdown text="Dropdown" variant="outline-primary" slot="append">
        <b-dropdown-item>Action C</b-dropdown-item>
        <b-dropdown-item>Action D</b-dropdown-item>
      </b-dropdown>
    </b-input-group>

    <p/>

    <Loader v-if="isLoading" />

    <div class="list-group" v-if="!isLoading">      
      <a v-bind:href="'#/namespaces/'+ns.metadata.namespace+'/pods/'+ns.metadata.name" class="list-group-item list-group-item-action" v-for="(ns, index) in signalChange()" :key="index">  
        <div class="d-flex justify-content-between align-items-center">
          <h5 class="mb-1">{{ ns.metadata.name }}</h5>
          <small><span>{{ ns.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
        </div>
      </a>
    </div>

    <div class="alert alert-info" role="alert" v-if="signalChange().length == 0 && !isLoading">
      No running pods found
    </div>
    
    <div class="alert alert-danger" role="alert" v-if="isError">
      <h4 class="alert-heading">Oops! <span class="navbar-brand fas fa-sad-tear"></span></h4>
      <p>Unable to load pods</p>
      <hr/>
      <p class="mb-0"><pre>{{ errMsg }}</pre></p>
    </div>

  </div>

</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
import Loader from './Loader.vue'
export default {
  name: 'Pods',
  components: {
    Breadcrumb,
    Loader
  },
  data() {
    return {
      searchString: '',
      isLoading: true,
      isError: false,
      errMsg: null,
      isSortDown: true,
      sortBy: "name"
    }
  },
  mounted() {
    this.$store.dispatch('getPods', this.$route.params.namespace).then(() => {
      this.isLoading = false
    }).catch(err => {
      this.isError = true
      this.errMsg = err 
    }).finally(() => {
      this.isLoading = false
    })
    this.$store.dispatch('addRecentNamespace', this.$route.params.namespace)
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
      let query = {
        str: this.searchString,
        sort: (this.isSortDown ? 'asc' : 'desc'),
        phase: 'Running'
      }
      return this.filterPods(query);
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


