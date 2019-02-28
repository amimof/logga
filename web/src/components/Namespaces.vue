<template>
  <div class="container">
    <b-input-group :prepend="items.length.toString()">
      <b-form-input v-model="str" placeholder="Search Namespace"></b-form-input>
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

    <h5 v-if="items.length == 0 && !isLoading && !isError">No namespaces found</h5>

    <div class="row" v-if="!isLoading && !isError">
      <div class="col" v-if="items.length > 0">
        <h6 v-if="recentNamespaces.length > 0">All</h6>
        <NamespacesList :items="items" :active="activeNamespace"/>
      </div>
      <div class="col-4" v-if="recentNamespaces.length > 0">
        <h6>Recent</h6>
        <RecentNamespacesList class="col"/>
      </div>
    </div>

    <ErrorCard title="Unable to load namespaces" :error="error" v-if="isError"/>

  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Loader from './Loader.vue'
import ErrorCard from './ErrorCard.vue'
import NamespacesList from './NamespacesList.vue'
import RecentNamespacesList from './RecentNamespacesList.vue'

const keyup = 38;
const keydown = 40;
const keyenter = 13;
const keychars = [48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 109]

export default {
  name: 'Namespaces',
  components: {
    Loader,
    ErrorCard,
    NamespacesList,
    RecentNamespacesList
  },
  data() {
    return {
      isLoading: true,
      isError: false,
      error: null,
      activeNamespace: 0
    }
  },
  mounted() {
    window.addEventListener("keydown", this.handlePress);
    this.$store.dispatch('getNamespaces').then(() => {
      this.isLoading = false
    }).catch(error => {
      this.isError = true
      this.error = error
    }).finally(() => {
      this.isLoading = false;
    })
  },
  beforeDestroy() {
    window.removeEventListener("keydown", this.handlePress);
  },
  methods: {
    handlePress() {

      // Up down
      if (event.keyCode == keyup && this.activeNamespace > 0) {
        this.activeNamespace--
      } else if (event.keyCode == keydown && this.activeNamespace < (this.items.length-1)) {
        this.activeNamespace++
      } else if (keychars.indexOf(event.keyCode) > -1) {
        this.activeNamespace = 0;
        this.$el.querySelector('input').focus();
      }

      // Enter
      if (event.keyCode == keyenter) {
        let i = this.items[this.activeNamespace].metadata.name;
        this.$router.push({ path: `/namespaces/${i}/pods` })
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
      'namespaces',
      'recentNamespaces',
      'nsSearchString',
      'nsSort',
    ]),
    ...mapGetters([
      'filterNamespaces',
      'sortNamespaces'
    ]),
    str: {
      set(str) {
        this.$store.dispatch('setNamespaceSearchString', str)
      },
      get() {
        return this.nsSearchString;
      }
    },
    isSortDown: {
      get() {
        return this.nsSort == 'asc' ? true : false
      },
      set(sort) {
        this.$store.dispatch('setNamespaceSort', sort ? 'asc' : 'desc')
      }
    },
    items() {
      let result = this.filterNamespaces(this.nsSearchString);
      if(this.isSortDown) {
        result = this.sortNamespaces('asc')
      } else {
        result = this.sortNamespaces('desc')
      }
      return result
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


