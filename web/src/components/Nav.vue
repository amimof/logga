<template>
  <div class="row align-items-center">
    <div class="logo col">logga</div>
    
    <Breadcrumb class="col" />
    <div class="row text-right align-items-center">
      <switches class="col" v-model="isDark" color="default"></switches>
      <span class="app-ver col">v1.0.0</span>
      <span class="github-link col">
        <a href="https://github.com/amimof/logga"><i class="fab fa-github text-dark"></i></a>
      </span>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Breadcrumb from './Breadcrumb.vue'
import Switches from 'vue-switches';

export default {
  name: 'Nav',
  components: {
    Breadcrumb,
    Switches
  },
  computed: {
    ...mapState([
      'theme'
    ]),
    isDark: {
      get() {
        return this.theme == 'dark' ? true : false;
      },
      set(val) {
        this.$store.dispatch('setTheme', val ? 'dark' : 'light')
      }
    }
  }

}
</script>

<style lang="scss" scoped>
.nav {
  align-content: center;
}
.logo {
  font-family: Consolas,monaco,monospace;
  color: #1DE9B6;
  -ms-flex: 0 0 80px;
  flex: 0 0 80px;
}
.vue-switcher {
  margin: 0;
}
.app-ver {
  font-size: 12px;
  color: gray;
}
.text-right {
  padding-right: 15px;
}
</style>


