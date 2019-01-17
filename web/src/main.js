import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import router from './router'
import store from './store'
import moment from 'vue-moment'

import '../node_modules/jquery/dist/jquery.min.js';
import '../node_modules/bootstrap/scss/bootstrap.scss';
import '../node_modules/bootstrap/dist/js/bootstrap.min.js';
import '../node_modules/lodash/lodash.min.js';
import './styles/main.scss'

// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false
Vue.use(store)
Vue.use(BootstrapVue);
Vue.use(moment)

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
