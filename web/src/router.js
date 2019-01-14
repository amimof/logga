import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import NotFoundComponent from './components/NotFoundComponent.vue'
import Namespaces from './components/Namespaces.vue'
import Pods from './components/Pods.vue'

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
      path: '/namespaces',
      name: 'Namespaces',
      component: Namespaces
    },
    {
      path: '/namespaces/:id/pods',
      name: 'Pods',
      component: Pods
    },
    {
      path: '*',
      component: NotFoundComponent
    }
  ]
})
