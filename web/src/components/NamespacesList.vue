<template>
  <div class="list-group">
    <a v-bind:href="'#/namespaces/'+item.metadata.name+'/pods'" class="list-group-item list-group-item-action" v-for="(item, index) in items" :key="index"> 
      <div class="d-flex justify-content-between align-items-center">
        <h5 class="mb-1">{{ item.metadata.name }}</h5>
        <small><span>{{ item.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
      </div>
    </a>
  </div>
</template>

<script>
export default {
  name: 'NamespacesList',
  props: {
    items: {
      type: Array,
      required: true,
      default: () => {
        return [];
      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


