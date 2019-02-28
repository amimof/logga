<template>
  <div class="container">
    <div class="row align-items-center">
      <div class="logo col">logga</div>
      
      <Breadcrumb class="col" />
      <div class="row text-right align-items-center">
        <span class="app-ver col">v1.0.0</span>
        <span class="github-link col">
          <a href="https://github.com/amimof/logga" target="_blank"><i class="fab fa-github text-dark"></i></a>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import Breadcrumb from './Breadcrumb.vue'

let body;

export default {
  name: 'Nav',
  components: {
    Breadcrumb
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
.app-ver {
  font-size: 12px;
  color: gray;
}
.text-right {
  padding-right: 15px;
}
</style>


