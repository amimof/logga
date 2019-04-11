<template>  
  <b-alert show variant="danger">
    <h4 class="alert-heading">{{ title }}</h4>
    <p>{{ description }} <i class="fas fa-arrow-circle-right" v-b-toggle.collapse1 v-b-tooltip.hover title="Show/Hide details"></i></p>
    <b-collapse id="collapse1" class="mt-2">
      <p class="mb-0"><pre>{{ error }}</pre></p>
    </b-collapse>
  </b-alert>
</template>

<script>

import _ from 'lodash'
export default {
  name: 'ErrorCard',
  props: {
    error: {
      type: Error,
      required: true
    },
    title: {
      type: String,
      required: true
    }
  },
  computed: {
    status() {
      if (_.has(this.error, 'response.status')) {
        return this.error.response.status;
      }
      return 0;
    },
    statusText() {
      if (_.has(this.error, 'response.statusText'))  {
        return this.error.response.statusText;
      }
      return "Unknown error";
    },
    description() {
      if(this.status == 0) {
        return this.statusText;
      }
      return `${this.status}: ${this.statusText}`
    }
  }
}
</script>

<style lang="scss" scoped>
pre {
  display: inline-block; 
  white-space: pre-wrap;
}

i {
  display: inline-block;
}
</style>


