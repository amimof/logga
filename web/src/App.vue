<template>
  <div id="app">
    <!-- Fixed navbar -->
    <!-- <nav class="navbar sticky-top navbar-expand-lg navbar-dark bg-dark">
      <div class="container">
        <router-link to ="/home">
          <i class="navbar-brand fab fa-linux"></i>
        </router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item active">
              <router-link to="/namespaces">
                <a class="nav-link">Namespaces</a>
              </router-link>
            </li>
          </ul>
        </div>
      </div>
    </nav> -->
    <div style="padding-top: 16px;">
      <router-view/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {}
}
</script>

<style>
</style>
