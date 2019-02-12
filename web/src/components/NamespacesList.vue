<template>
  <b-list-group>
    <b-list-group-item 
      v-for="(item, index) in items" :key="index"
      v-bind:href="'#/namespaces/'+item.metadata.name+'/pods'"
      :variant="variant" 
      :active="active == index">
      <div class="d-flex justify-content-between align-items-center">
        <h5 class="mb-1">{{ item.metadata.name }}</h5>
        <small><span>{{ item.metadata.creationTimestamp | moment("from", "now", true) }}</span></small>
      </div>
    </b-list-group-item>
  </b-list-group>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'NamespacesList',
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
  }
}
</script>

<style lang="scss" scoped>
</style>


