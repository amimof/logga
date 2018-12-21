import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './styles/main.less'
import '../node_modules/jquery/dist/jquery.min.js';
import '../node_modules/bootstrap/scss/bootstrap.scss';
import '../node_modules/bootstrap/dist/js/bootstrap.min.js';

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
