<template>
  <div class="list-group">
    <a v-bind:href="'#/namespaces/'+ns+'/pods'" class="list-group-item list-group-item-action" v-for="(ns, index) in recentNamespaces" :key="index"> 
      <div class="d-flex justify-content-between align-items-center">
        <h5 class="mb-1">{{ ns }}</h5>
      </div>
    </a>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'RecentNamespacesList',
  computed: {
    ...mapState([
      'recentNamespaces'
    ])
  }
}
</script>

<style lang="scss" scoped>
</style>


