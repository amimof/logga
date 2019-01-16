import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import NotFoundComponent from './components/NotFoundComponent.vue'
import Namespaces from './components/Namespaces.vue'
import Pods from './components/Pods.vue'
import Pod from './components/Pod.vue'

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
      component: Namespaces,
      meta: {
        breadcrumb: [{
          name: 'Namespaces',
        }]
      }
    },
    {
      path: '/namespaces/:id/pods',
      name: 'Pods',
      component: Pods,
      meta: {
        breadcrumb: [{
          name: 'Namespaces',
        },{
          name: 'Pods',
        }]
      }
    },
    {
      path: '/namespaces/:namespace/pods/:pod',
      name: 'Pod',
      component: Pod,
      meta: {
        breadcrumb: [{
          name: 'Namespaces',
        },{
          name: 'Pods',
        },{
          name: ':pod',
        }]
      }      
    },
    {
      path: '*',
      component: NotFoundComponent
    }
  ]
})
