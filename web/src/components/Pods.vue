<template>
  <div class="container">

    <b-input-group :prepend="items.length.toString()">
      <b-form-input v-model="str" placeholder="Search Pods"></b-form-input>
      <b-input-group-append>
        <b-button v-on:click="str = ''" variant="outline-secondary">
          <i class="fas fa-times"></i>
        </b-button>
        <b-button v-on:click="isSortDown = !isSortDown" variant="outline-secondary">
          <i class="fas fa-sort-amount-down" v-if="isSortDown"></i>
          <i class="fas fa-sort-amount-up" v-if="!isSortDown"></i>
        </b-button>
      </b-input-group-append>
    </b-input-group>

    <p/>

    <Loader v-if="isLoading" />

    <h5 v-if="items.length == 0 && !isLoading && !isError">No pods found</h5>

    <PodsList :items="items" :active="activePod" />
    
    <ErrorCard title="Unable to load pods" :error="error" v-if="isError"/>

  </div>

</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
import PodsList from './PodsList.vue'

const keyup = 38;
const keydown = 40;
const keyenter = 13;
const keychars = [48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 109]

export default {
  name: 'Pods',
  components: {
    Loader,
    ErrorCard,
    PodsList
  },
  data() {
    return {
      isLoading: true,
      isError: false,
      error: null,
      str: '',
      activePod: 0
    }
  },
  mounted() {
    window.addEventListener("keydown", this.handlePress);
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
  beforeDestroy() {
    window.removeEventListener("keydown", this.handlePress);
  },
  methods: {
    handlePress() {

      // Up down
      if (event.keyCode == keyup && this.activePod > 0) {
        this.activePod--
      } else if (event.keyCode == keydown && this.activePod < (this.items.length-1)) {
        this.activePod++
      } else if (keychars.indexOf(event.keyCode) > -1) {
        this.activePod = 0;
        this.$el.querySelector('input').focus();
      }

      // Enter
      if (event.keyCode == keyenter) {
        let n = this.items[this.activePod].metadata.namespace;
        let p = this.items[this.activePod].metadata.name;
        this.$router.push({ path: `/namespaces/${n}/pods/${p}` })
      }

      // Blur and clear input
      if (event.keyCode == 27) {
        this.$el.querySelector('input').blur();
        this.str = "";
      }

    }
  },
  computed: {
    ...mapState([
      'pods',
      'podSort'
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


