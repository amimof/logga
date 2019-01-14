import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import '../node_modules/jquery/dist/jquery.min.js';
import '../node_modules/bootstrap/scss/bootstrap.scss';
import '../node_modules/bootstrap/dist/js/bootstrap.min.js';
import '../node_modules/lodash/lodash.min.js';
import './styles/main.scss'

Vue.config.productionTip = false
Vue.use(store)

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
