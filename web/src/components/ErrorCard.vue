<template>  
  <div class="alert alert-danger" role="alert">
    <h4 class="alert-heading">Oops! <span class="navbar-brand fas fa-sad-tear"></span></h4>
    <p>{{ title }} <i class="fas fa-arrow-circle-right" v-b-toggle.collapse1 v-b-tooltip.hover title="Show/Hide details"></i></p>
    <b-collapse id="collapse1" class="mt-2">
      <p class="mb-0"><pre>{{ error }}</pre></p>
    </b-collapse>
  </div>
</template>

<script>

export default {
  name: 'ErrorCard',
  props: {
    error: Error,
    title: String
  }
}
</script>

<style lang="scss" scoped>
pre {
  display: inline-block;
}

i {
  display: inline-block;
}
</style>


