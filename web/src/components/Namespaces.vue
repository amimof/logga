<template>
  <div class="container">

    <breadcrumb/>

    <b-input-group :prepend="signalChange().length.toString()">
      <b-form-input v-model="searchString" @input="signalChange" placeholder="Search Namespace"></b-form-input>
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

    <div class="alert alert-info" role="alert" v-if="signalChange().length == 0 && !isLoading && !isError">
      No matching namespaces found
    </div>

    <div class="row" v-if="!isLoading">
      <div class="col" v-if="signalChange().length > 0">
        <h4>All</h4>
        <a v-bind:href="'#/namespaces/'+ns.metadata.name+'/pods'" class="list-group-item list-group-item-action" v-for="(ns, index) in signalChange()" :key="index"> 
          <div class="d-flex justify-content-between align-items-center">
            <h5 class="mb-1">{{ ns.metadata.name }}</h5>
            <small><span>{{ ns.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
          </div>
        </a>
      </div>

      <RecentNamespacesList v-if="recentNamespaces.length > 0 && !isError" />

    </div>

    <ErrorCard title="Unable to load namespaces" :error="error" v-if="isError"/>

  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
import RecentNamespacesList from './RecentNamespacesList.vue'

export default {
  name: 'Namespaces',
  components: {
    Breadcrumb,
    Loader,
    ErrorCard,
    RecentNamespacesList
  },
  data() {
    return {
      searchString: '',
      isLoading: true,
      isError: false,
      error: null,
      isSortDown: true,
      sortBy: "name"
    }
  },
  mounted() {
    this.$store.dispatch('getNamespaces').then(() => {
      this.isLoading = false
    }).catch(error => {
      this.isError = true
      this.error = error
    }).finally(() => {
      this.isLoading = false;
    })
  },
  computed: {
    ...mapState([
      'namespaces',
      'recentNamespaces'
    ]),
    ...mapGetters([
      'filterNamespaces'
    ])
  },
  methods: {
    signalChange() {
      let query = {
        str: this.searchString,
        sort: (this.isSortDown ? 'asc' : 'desc')
      }
      return this.filterNamespaces(query);
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


