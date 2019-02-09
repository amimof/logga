<template>
  <div class="container">

    <Nav/>
    
    <p/>

    <b-input-group :prepend="items.length.toString()">
      <b-form-input v-model="str" placeholder="Search Pods"></b-form-input>
      <b-input-group-append>
        <b-button v-on:click="str = ''" variant="outline-primary">
          <i class="fas fa-times"></i>
        </b-button>
        <b-button v-on:click="isSortDown = !isSortDown" variant="outline-primary">
          <i class="fas fa-sort-amount-down" v-if="isSortDown"></i>
          <i class="fas fa-sort-amount-up" v-if="!isSortDown"></i>
        </b-button>
      </b-input-group-append>
    </b-input-group>

    <p/>

    <Loader v-if="isLoading" />

    <div class="alert alert-info" role="alert" v-if="items.length == 0 && !isLoading && !isError">
      No pods found
    </div>

    <PodsList :items="items" />
    
    
    <ErrorCard title="Unable to load pods" :error="error" v-if="isError"/>

  </div>

</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Nav from './Nav.vue'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
import PodsList from './PodsList.vue'
export default {
  name: 'Pods',
  components: {
    Nav,
    Loader,
    ErrorCard,
    PodsList
  },
  data() {
    return {
      isLoading: true,
      isError: false,
      error: null,
      str: ''
    }
  },
  mounted() {
    this.$store.dispatch('getPods', this.$route.params.namespace).then(() => {
      this.isLoading = false
    }).catch(err => {
      this.isError = true
      this.error = err 
    }).finally(() => {
      this.isLoading = false
    })
    this.$store.dispatch('addRecentNamespace', this.$route.params.namespace)
  },
  computed: {
    ...mapState([
      'pods',
      'podSort',
      'theme'
    ]),
    ...mapGetters([
      'filterPods',
      'sortPods'
    ]),
    isSortDown: {
      get() {
        return this.podSort == 'asc' ? true : false
      },
      set(sort) {
        this.$store.dispatch('setPodSort', sort ? 'asc' : 'desc')
      }
    },
    variant() {
      if(this.theme == 'dark') {
        return 'default'
      } else {
        return 'light'
      }
    },
    items() {
      let result = this.filterPods(this.str);
      if(this.isSortDown) {
        result = this.sortPods('asc')
      } else {
        result = this.sortPods('desc')
      }
      return result
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


