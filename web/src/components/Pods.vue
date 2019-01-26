<template>
  <div class="container">

    <Nav/>
    
    <p/>

    <b-input-group :prepend="items.length.toString()">
      <b-form-input v-model="str" placeholder="Search Pods"></b-form-input>
      <b-input-group-append>
        <b-button v-on:click="str = ''" variant="outline-primary">
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
      <a v-bind:href="'#/namespaces/'+item.metadata.namespace+'/pods/'+item.metadata.name" class="list-group-item list-group-item-action" v-for="(item, index) in items" :key="index">  
        <div class="d-flex justify-content-between align-items-center">
          <h5 class="mb-1">{{ item.metadata.name }}</h5>
          <small><span>{{ item.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
        </div>
      </a>
    </div>

    <div class="alert alert-info" role="alert" v-if="items.length == 0 && !isLoading && !isError">
      No running pods found
    </div>
    
    <ErrorCard title="Unable to load pods" :error="error" v-if="isError"/>

  </div>

</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Nav from './Nav.vue'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
export default {
  name: 'Pods',
  components: {
    Nav,
    Loader,
    ErrorCard
  },
  data() {
    return {
      isLoading: true,
      isError: false,
      error: null,
      str: ''
    }
  },
  mounted() {
    this.$store.dispatch('getPods', this.$route.params.namespace).then(() => {
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
      'pods',
      'podSort'
    ]),
    ...mapGetters([
      'filterPods',
      'sortPods'
    ]),
    isSortDown: {
      get() {
        return this.podSort == 'asc' ? true : false
      },
      set(sort) {
        this.$store.dispatch('setPodSort', sort ? 'asc' : 'desc')
      }
    },
    items() {
      let result = this.filterPods(this.str);
      if(this.isSortDown) {
        result = this.sortPods('asc')
      } else {
        result = this.sortPods('desc')
      }
      return result
    }
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
  }
}
</script>

<style lang="scss" scoped>
</style>


