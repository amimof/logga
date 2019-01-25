import Vue from 'vue'
import Router from 'vue-router'
import NotFoundComponent from './components/NotFoundComponent.vue'
import Namespaces from './components/Namespaces.vue'
import Pods from './components/Pods.vue'
import Pod from './components/Pod.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/namespaces'
    },
    {
      path: '/namespaces',
      name: 'Namespaces',
      component: Namespaces,
    },
    {
      path: '/namespaces/:namespace/pods',
      name: 'Pods',
      component: Pods,
    },
    {
      path: '/namespaces/:namespace/pods/:pod',
      name: 'Pod',
      component: Pod     
    },
    {
      path: '*',
      component: NotFoundComponent
    }
  ]
})
