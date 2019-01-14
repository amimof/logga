<template>
  <div id="container">

  <div class="input-group mb-3">
    <input type="text" v-model="namespace" v-on:change="signalChange" class="form-control" placeholder="Search Namespace" aria-label="Search Namespace" aria-describedby="search-button">
    <div class="input-group-append">
      <button class="btn btn-outline-primary" type="button" id="search-button">Search</button>
    </div>
  </div>

    <div class="list-group">
      <a v-bind:href="'#/namespaces/'+ns.metadata.name+'/pods'" class="list-group-item list-group-item-action" v-for="(ns, index) in namespaces.items" :key="index"> 
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
import { mapState } from 'vuex'
export default {
  name: 'Namespaces',
  data: function() {
    return {
      namespace: ''
    }
  },
  mounted() {
    this.$store.dispatch('getNamespaces') 
  },
  computed: mapState([
    'namespaces'
  ]),
  methods: {
    signalChange: function(evt){
      this.$store.dispatch('getNamespaces', this.namespace)
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


