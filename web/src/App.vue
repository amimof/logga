<template>
  <div id="app" :style="{'background-color': variant}">
    <Nav/>
    <router-view/>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Nav from './components/Nav.vue'
export default {
  name: 'App',
  components: {
    Nav,
  },
  computed: {
    ...mapState([
      'theme'
    ]),
    variant() {
      if(this.theme == 'dark') {
        return '#212429'
      } else {
        return '#fff'
      }
    }  
  },
}
</script>

<style lang="scss">
</style>
