import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)
Vue.use(axios)

const apiUrl = 'http://localhost:8080/api'

export default new Vuex.Store({
  state: {
     namespaces: {},
     pods: {},
  },
  actions: {
    getNamespaces ({ commit }, searchString) {
      axios
        .get(`${apiUrl}/namespaces`)
        .then(r => r.data)
        .then(namespaces => {
          if(searchString) {
            let filtered = _.filter(this.state.namespaces.items, function(item){
              return item.metadata.name === searchString;
            });
            namespaces.items = filtered;
          }
          commit('SET_NAMESPACES', namespaces)
        })
    },
    getPods ({ commit }, namespace) {
      axios
        .get(`${apiUrl}/namespaces/${namespace}/pods`)
        .then(r => r.data)
        .then(pods => {
          commit('SET_PODS', pods)
        })
    }
  },
  mutations: {
    SET_NAMESPACES (state, namespaces) {
      state.namespaces = namespaces
    },
    SET_PODS (state, pods) {
      state.pods = pods
    }
  }
})
