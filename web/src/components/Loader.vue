<template>
  <div class="d-flex justify-content-center">
    <pulse-loader color="rgb(0, 123, 255)"></pulse-loader>
  </div>
</template>

<script>
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
export default {
  name: 'Loader',
  components: {
    PulseLoader
  }
}
</script>

<style lang="scss" scoped>
</style>


