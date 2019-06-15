import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/namespace/:namespace/deployments',
      name: 'deployments',
      component: () => import(/* webpackChunkName: "about" */ './views/Deployments.vue'),
    },
    {
      path: '/nodes',
      name: 'nodes',
      component: () => import(/* webpackChunkName: "about" */ './views/Nodes.vue'),
    },
    {
      path: '/namespaces/:namespace/deployment/:deployment/pods',
      name: 'pods',
      component: () => import(/* webpackChunkName: "about" */ './views/Pods.vue'),
    },
  ],
});
