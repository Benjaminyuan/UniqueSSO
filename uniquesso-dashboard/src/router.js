import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './views/Home.vue'
import Login from "./views/Login.vue"
import Cookies from 'js-cookies'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/cas/login',
    name: 'Login',
    component: Login,
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((from, to, next) => {
  const tgc = Cookies.getItem("CASTGC")
  if (to.path !== "/cas/login" && from.path !== "/cas/login" && !tgc) {
    next("/cas/login")
  } else next()
})

export default router
