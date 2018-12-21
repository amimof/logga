import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import NotFoundComponent from './components/NotFoundComponent.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/home'
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    },
    {
      path: '*',
      component: NotFoundComponent
    }
  ]
})
