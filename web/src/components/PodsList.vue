<template>
  <b-list-group>
    <b-list-group-item 
      v-for="(item, index) in items" :key="index"
      v-bind:href="'#/namespaces/'+item.metadata.namespace+'/pods/'+item.metadata.name"
      :active="active == index"
      :v-if="item.status.phase == 'Running'">
      <div class="d-flex justify-content-between align-items-center">
        <span class="mb-1">{{ item.metadata.name }}</span>
        <small></small>
        <small><span><i class="far fa-clock"></i> {{ item.metadata.creationTimestamp | moment("from", "now", true) }} <i class="fas fa-long-arrow-alt-up"></i> {{ item.spec.containers.length }} <i class="fas fa-redo-alt"></i> {{ item | numContainerRestarts }}</span></small>
      </div>
    </b-list-group-item>
  </b-list-group>
</template>

<script>
export default {
  name: 'PodsList',
  props: {
    items: {
      type: Array,
      required: true,
      default: () => {
        return [];
      }
    },
    active: {
      type: Number,
      default: () => {
        return 0
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
.created {
  opacity: 0.5;
}
</style>


