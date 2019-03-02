<template>
<div class="h-100">
  <b-list-group>
    <b-list-group-item 
      v-for="(item, index) in items" :key="index"
      v-bind:href="'#/namespaces/'+item.metadata.name+'/pods'" 
      :active="active == index">
      <div class="d-flex justify-content-between align-items-center">
        <span class="mb-1">{{ item.metadata.name }}</span>
        <small><span class="namespace-icons"><i class="far fa-clock"></i> {{ item.metadata.creationTimestamp | moment("from", "now", true) }}</span> <span class="namespace-icons"><PodCount :namespace="item.metadata.name"/></span></small>
      </div>
    </b-list-group-item>
  </b-list-group>
  </div>
</template>

<script>
import PodCount from './PodCount.vue'
export default {
  name: 'NamespacesList',
  components: {
    PodCount
  },
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
  }
}
</script>

<style lang="scss">
.namespace-icons {
  padding-left: 8px;
  padding-right: 8px;
}
</style>


