import Vue from 'vue'
import VueRouter from 'vue-router'
import AreaList from '../components/AreaList.vue'
import About from '../components/About.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: AreaList
  },
  {
    path: '/about',
    name: 'About',
    component: About
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
