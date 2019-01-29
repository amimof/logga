<template>
  <b-list-group>
    <b-list-group-item
      v-for="(ns, index) in recentNamespaces" :key="index"
      v-bind:href="'#/namespaces/'+ns+'/pods'"
      :variant="variant"> 
      <div class="d-flex justify-content-between align-items-center">
        <h5 class="mb-1">{{ ns }}</h5>
      </div>
    </b-list-group-item>
  </b-list-group>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'RecentNamespacesList',
  computed: {
    ...mapState([
      'recentNamespaces',
      'theme'
    ]),
    variant() {
      if(this.theme == 'dark') {
        return 'default'
      } else {
        return 'light'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


