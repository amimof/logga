<template>
  <span v-if="isLoaded"><i class="fas fa-hashtag"></i> {{ podCount }}</span>
</template>

<script>
import { mapGetters } from 'vuex'
import _ from 'lodash'
export default {
  name: 'PodCount',
  data () {
    return {
      pods: {},
      isLoaded: false
    }
  }, 
  props: {
    namespace: String
  },
  methods: {
    getPods() {
      this.get(`${this.getAPIURL}/namespaces/${this.namespace}/pods`)
        .then(r => r.data)
        .then(pods => {
          this.pods = pods;
          this.isLoaded = true;
        })
    }
  },
  mounted() {
    this.getPods();
  },
  computed: {
    ...mapGetters([
      'get',
      'getAPIURL'
    ]),
    podCount() {
      if(_.has(this.pods, 'items')) {
        return this.pods.items.length;
      }
      return 0;
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


