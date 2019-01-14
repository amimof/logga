<template>
  <div id="container">
    {{ pods }}
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'Pods',
  mounted() {
    this.$store.dispatch('getPods', this.$route.params.id) 
  },
  computed: mapState([
    'pods'
  ])
}
</script>

<style lang="scss" scoped>
</style>


