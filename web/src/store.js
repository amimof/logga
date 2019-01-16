import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'
import _ from 'lodash'

Vue.use(Vuex)
Vue.use(VueAxios, axios)

const apiUrl = 'http://localhost:8080/api'

export default new Vuex.Store({
  state: {
     namespaces: {},
     pods: {},
     pod: {},
     podLog: {}
  },
  actions: {
    getNamespaces ({ commit }) {
      axios
        .get(`${apiUrl}/namespaces`)
        .then(r => r.data)
        .then(namespaces => {
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
    },
    getPod ({ commit }, {namespace, pod}) {
      axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}`)
        .then(r => r.data)
        .then(pod => {
          commit('SET_POD', pod)  
        })
    },
    getPodLog ({ commit }, { namespace, pod }) {
      axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}/log`)
        .then(r => r.data)
        .then(pod => {
          console.log(pod)
          commit('SET_POD_LOG', pod)  
        })
    }
  },
  getters: {
    filterNamespaces: (state) => (searchString) => {
      let filtered = _.filter(state.namespaces.items, function(item){
        return item.metadata.name.includes(searchString);
      });
      return filtered;
    },
    filterPods: (state) => (searchString) => {
      let filtered = _.filter(state.pods.items, function(item) {
        return item.metadata.name.includes(searchString);
      });
      return filtered
    }
  },
  mutations: {
    SET_NAMESPACES (state, namespaces) {
      state.namespaces = namespaces
    },
    SET_PODS (state, pods) {
      state.pods = pods
    },
    SET_POD (state, pod) {
      state.pod = pod
    },
    SET_POD_LOG (state, podLog) {
      state.podLog = podLog;
    }
  }
})
