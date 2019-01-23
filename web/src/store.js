import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'
import _ from 'lodash'

Vue.use(Vuex)
Vue.use(VueAxios, axios)

const apiUrl = 'http://localhost:8080/api'
let eventSource;

export default new Vuex.Store({
  state: {
     namespaces: {},
     recentNamespaces: [],
     pods: {},
     pod: null,
     podLog: [],
     lines: []
  },
  actions: {
    getNamespaces({ commit }) {
      return axios
        .get(`${apiUrl}/namespaces`)
        .then(r => r.data)
        .then(namespaces => {
          commit('SET_NAMESPACES', namespaces)
        })
    },
    getPods({ commit }, namespace) {
      return axios
        .get(`${apiUrl}/namespaces/${namespace}/pods`)
        .then(r => r.data)
        .then(pods => {
          commit('SET_PODS', pods)
        })
    },
    getPod({ commit }, {namespace, pod}) {
      axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}`)
        .then(r => r.data)
        .then(pod => {
          commit('SET_POD', pod)  
        })
    },
    getPodLog({ commit }, { namespace, pod }) {
      return axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}/log?tailLines=1000`)
        .then(r => r.data)
        .then(pod => {
          let lines = pod.split(/\r?\n/)
          if(lines[lines.length-1] == "") {
            lines.splice(lines.length-1, 1)
          }
          commit('SET_POD_LOG', lines)  
        })
    },
    addRecentNamespace({commit}, namespace) {
      commit('ADD_RECENT_NAMESPACE', namespace)
    },
    streamPodLog({ commit }, { namespace, pod }) {
      eventSource = new EventSource(`${apiUrl}/namespaces/${namespace}/pods/${pod}/log?watch=true&tailLines=1000`)
      eventSource.onmessage = function (e) {
        commit("ADD_LINE", e.data)
      }
    },
    closeStream() {
      eventSource.close()
    }
  },
  getters: {
    filterList: () => (list, query) => {
      let q = query || { str: "", sort: "asc", phase: "" }
      
      // Remove whitespaces
      q.str = q.str.replace(/\s/g, "");

      // Filter by name
      let filtered = _.filter(list, function(item){
        return item.metadata.name.includes(q.str);
      });

      // Order by name
      filtered = _.orderBy(filtered, function(item) {
        return item.metadata.name
      }, [q.sort])

      // Filter by phase
      if(q.phase) {
        filtered = _.filter(filtered, function(item) {
          return item.status.phase === q.phase
        })
      }      
      return filtered
    },
    filterNamespaces: (state, getters) => (query) => {
      return getters.filterList(state.namespaces.items, query)
    },
    filterPods: (state, getters) => (query) => {
      return getters.filterList(state.pods.items, query)
    }
  },
  mutations: {
    SET_NAMESPACES(state, namespaces) {
      state.namespaces = namespaces
    },
    SET_PODS(state, pods) {
      state.pods = pods
    },
    SET_POD(state, pod) {
      state.pod = pod
    },
    SET_POD_LOG(state, podLog) {
      state.podLog = podLog;
    },
    ADD_LINE(state, line) {
      if(state.podLog.length >= 1001) {
        state.podLog.splice(0, 1)
      }
      state.podLog.push(line);
    },
    ADD_RECENT_NAMESPACE(state, namespace) {
      if(state.recentNamespaces.length >= 10) {
        state.recentNamespaces.splice(0, 1)
      }
      let index = state.recentNamespaces.indexOf(namespace)
      if(index > -1) {
        state.recentNamespaces.splice(index, 1)
      }
      state.recentNamespaces.unshift(namespace)
    }
  }
})
