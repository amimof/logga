<template>
  <div class="col" v-if="recentNamespaces.length > 0">
    <h4>Recent</h4>
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
  name: 'Namespaces',
  computed: {
    ...mapState([
      'recentNamespaces'
    ])
  }
}
</script>

<style lang="scss" scoped>
</style>


