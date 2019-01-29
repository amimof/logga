<template>
  <div id="app" :style="{'background-color': variant}">
  <!-- <div id="app"> -->
    <div style="padding-top: 16px;">
      <router-view/>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'App',
  components: {},
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
