<template>
  <b-list-group>
    <b-list-group-item 
      v-for="(item, index) in items" :key="index"
      v-bind:href="'#/namespaces/'+item.metadata.namespace+'/pods/'+item.metadata.name"
      :variant="variant"
      :v-if="item.status.phase == 'Running'">
      <div class="d-flex justify-content-between align-items-center">
        <h5 class="mb-1">{{ item.metadata.name }}</h5>
        <small></small>
        <small><span>{{ item.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
      </div>
    </b-list-group-item>
  </b-list-group>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'PodsList',
  props: {
    items: {
      type: Array,
      required: true,
      default: () => {
        return [];
      }
    }
  },
  computed: {
    ...mapState([
      'theme'
    ]),
    variant() {
      if(this.theme == 'dark') {
        return 'default'
      } else {
        return 'light'
      }
    }
  },
  filters: {
    numContainersReady: function(pod) {
      if(!pod) {
        return 0;
      }
      var numReady = 0;
      for(let i = 0; i < pod.status.containerStatuses.length; i++) {
        if(pod.status.containerStatuses[i].ready) {
          numReady++;
        }
      }
      return numReady;
    },
    numContainerRestarts: function(pod) {
      if(!pod) {
        return 0;
      }
      var numRestarts = 0;
      for(let i = 0; i < pod.status.containerStatuses.length; i++)  {
        numRestarts += pod.status.containerStatuses[i].restartCount;
      }
      return numRestarts;
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


