<template>
  <div class="container">

    <Nav/>
    
    <p/>

    <b-input-group :prepend="items.length.toString()">
      <b-form-input v-model="str" placeholder="Search Namespace"></b-form-input>
      <b-input-group-append>
        <b-button v-on:click="str = ''" variant="outline-primary">
          <i class="fas fa-times"></i>
        </b-button>
        <b-button v-on:click="isSortDown = !isSortDown" variant="outline-primary">
          <i class="fas fa-sort-amount-down" v-if="isSortDown"></i>
          <i class="fas fa-sort-amount-up" v-if="!isSortDown"></i>
        </b-button>
      </b-input-group-append>
    </b-input-group>

    <p/>

    <Loader v-if="isLoading" />

    <div class="alert alert-info" role="alert" v-if="items.length == 0 && !isLoading && !isError">
      No matching namespaces found
    </div>

    <div class="row" v-if="!isLoading && !isError">
      <div class="col" v-if="items.length > 0">
        <h4>All</h4>
        <NamespacesList :items="items"/>
      </div>
      <div class="col" v-if="recentNamespaces.length > 0">
        <h4>Recent</h4>
        <RecentNamespacesList class="col" />
      </div>
    </div>

    <ErrorCard title="Unable to load namespaces" :error="error" v-if="isError"/>

  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Nav from './Nav.vue'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
import NamespacesList from './NamespacesList.vue'
import RecentNamespacesList from './RecentNamespacesList.vue'

export default {
  name: 'Namespaces',
  components: {
    Nav,
    Loader,
    ErrorCard,
    NamespacesList,
    RecentNamespacesList
  },
  data() {
    return {
      isLoading: true,
      isError: false,
      error: null,
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
      'recentNamespaces',
      'nsSearchString',
      'nsSort',
      'theme'
    ]),
    ...mapGetters([
      'filterNamespaces',
      'sortNamespaces'
    ]),
    str: {
      set(str) {
        this.$store.dispatch('setNamespaceSearchString', str)
      },
      get() {
        return this.nsSearchString;
      }
    },
    isSortDown: {
      get() {
        return this.nsSort == 'asc' ? true : false
      },
      set(sort) {
        this.$store.dispatch('setNamespaceSort', sort ? 'asc' : 'desc')
      }
    },
    items() {
      let result = this.filterNamespaces(this.nsSearchString);
      if(this.isSortDown) {
        result = this.sortNamespaces('asc')
      } else {
        result = this.sortNamespaces('desc')
      }
      return result
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


