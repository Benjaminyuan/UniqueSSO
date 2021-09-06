import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import Toasted from 'vue-toasted'
import Particles from 'particles.vue'

Vue.config.productionTip = false

Vue.use(Toasted, {
  theme: "bubble",
  position: "bottom-left",
  duration: 3000,
  keepOnHover: true,
})

Vue.use(Particles)

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')