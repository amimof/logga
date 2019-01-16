<template>
  <div id="container">
    <ol class="breadcrumb">
      <li class="breadcrumb-item active" aria-current="page" 
        v-for="(item, index) in $route.meta.breadcrumb" 
        :key="index">
        {{ item.name }}
      </li>
    </ol>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'Breadcrumb',
  data () {
    return {
      breadcrumbList: []
    }
  },
  computed: {
    ...mapState([
      'currentPath'
    ]),
  },
  watch: {
    '$route' () {
      this.breadcrumbList = this.$route.meta.breadcrumb
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


