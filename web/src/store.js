import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'
import _ from 'lodash'

Vue.use(Vuex)
Vue.use(VueAxios, axios)

let apiUrl =  process.env.NODE_ENV === 'dev' ? `http://${process.env.VUE_APP_API_HOST}:${process.env.VUE_APP_API_PORT}/api` : '/api'
let eventSource;

export default new Vuex.Store({
  state: {
    namespaces: {},
    recentNamespaces: [],
    pods: {},
    pod: null,
    podLog: [],
    lines: [],
    nsSort: 'asc',
    podSort: 'asc',
    nsSearchString: '',
    podSearchString: '',
    theme: 'dark',
    maxLines: 1000,
    lineStart: 0
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
      return axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}`)
        .then(r => r.data)
        .then(pod => {
          commit('SET_POD', pod)  
        })
    },
    getPodLog({ commit }, { namespace, pod, container }) {
      return axios
        .get(`${apiUrl}/namespaces/${namespace}/pods/${pod}/log?container=${container}&tailLines=1000`)
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
    streamPodLog({ commit }, { namespace, pod, container }) {
      eventSource = new EventSource(`${apiUrl}/namespaces/${namespace}/pods/${pod}/log?watch=true&container=${container}&tailLines=1000`)
      eventSource.onmessage = function (e) {
        commit("ADD_LINE", e.data)
      }
    },
    closeStream() {
      if(eventSource) {
        eventSource.close()
      }
    },
    setNamespaceSearchString({ commit }, searchString) {
      commit('SET_NAMESPACE_SEARCH_STRING', searchString)
    },
    setPodSearchString({ commit }, searchString) {
      commit('SET_POD_SEARCH_STRING', searchString)
    },
    setNamespaceSort({ commit }, sort) {
      commit('SET_NAMESPACE_SORT', sort)
    },
    setPodSort({ commit }, sort) {
      commit('SET_POD_SORT', sort)
    },
    setTheme({ commit }, theme) {
      commit('SET_THEME', theme)
    },
    resetLineStart({ commit }) {
      commit('SET_LINE_START', 0);
    }
  },
  getters: {
    get: () => (url) => {
      return axios.get(url)
    },
    getAPIURL: () => {
      return apiUrl;
    },
    sortList: () => (list, sort) => {
      return _.orderBy(list, function(item) {
        return item.metadata.name
      }, [sort])
    },
    filterList: () => (list, searchString) => {
      let q = searchString || ""
      
      // Remove whitespaces
      q = q.replace(/\s/g, "")

      // Filter by name
      let filtered = _.filter(list, function(item){
        return item.metadata.name.includes(q);
      });

      return filtered
    },
    filterNamespaces: (state, getters) => (searchString) => {
      state.nsSearchString = searchString;
      return getters.filterList(state.namespaces.items, searchString)
    },
    filterPods: (state, getters) => (searchString) => {
      state.podSearchString = searchString;
      return getters.filterList(state.pods.items, searchString)
    },
    sortNamespaces: (state, getters) => (sort) => {
      state.nsSort = sort;
      let list = getters.filterNamespaces(state.nsSearchString)
      return getters.sortList(list, sort)
    },
    sortPods: (state, getters) => (sort) => {
      state.podSort = sort;
      let list = getters.filterPods(state.podSearchString);
      return getters.sortList(list, sort)
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
      if(state.podLog.length >= state.maxLines) {
        state.podLog.splice(0, 1)
        state.lineStart += 1;
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
    },
    SET_NAMESPACE_SEARCH_STRING(state, searchString) {
      state.nsSearchString = searchString;
    },
    SET_NAMESPACE_SORT(state, sort) {
      state.nsSort = sort
    },
    SET_POD_SEARCH_STRING(state, searchString) {
      state.podSearchString = searchString;
    },
    SET_POD_SORT(state, sort) {
      state.podSort = sort
    },
    SET_THEME(state, theme) {
      state.theme = theme;
    },
    SET_LINE_START(state, num) {
      state.lineStart = num;
    }
  }
})
