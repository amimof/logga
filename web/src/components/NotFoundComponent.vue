<template>
  <div class="container">
    <div class="alert alert-dark" role="alert">
      {{ msg }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFoundComponent',
  props: {
    msg: {
      type: String,
      default: () => {
        return "Ooops! Page not found"
      }
    }
  },
}
</script>

<style scoped>
</style>
