<template>
  <div id="container">

    <breadcrumb/>

    <div class="d-flex">
      <h3 class="p-2">{{ signalChange().length }}</h3>
      <div class="p-2 input-group mb-3">
        <input type="text" v-model="searchString" @input="signalChange" class="form-control" placeholder="Search Namespace" aria-label="Search Namespace" aria-describedby="search-button">
        <div class="input-group-append">
          <button class="btn btn-outline-primary" type="button" id="search-button">Search</button>
        </div>
      </div>
    </div>

    <div class="list-group">
      <a v-bind:href="'#/namespaces/'+ns.metadata.name+'/pods'" class="list-group-item list-group-item-action" v-for="(ns, index) in signalChange()" :key="index"> 
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1">{{ ns.metadata.name }}</h5>
          <small>{{ ns.metadata.creationTimestamp }}</small>
        </div>
        <!-- <p class="mb-1">Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.</p>
        <small>Donec id elit non mi porta.</small> -->
      </a>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Breadcrumb from './Breadcrumb.vue'
export default {
  name: 'Namespaces',
  components: {
    Breadcrumb
  },
  data() {
    return {
      searchString: ''
    }
  },
  mounted() {
    this.$store.dispatch('getNamespaces')
  },
  computed: {
    ...mapState([
      'namespaces'
    ]),
    ...mapGetters([
      'filterNamespaces'
    ])
  },
  methods: {
    signalChange() {
      let result = this.namespaces.items || [];
      if(this.searchString.length > 0) {
        result = this.filterNamespaces(this.searchString);
      } 
      return result;
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


