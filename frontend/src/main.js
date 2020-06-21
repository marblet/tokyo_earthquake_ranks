import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueHead from 'vue-head'
import vuetify from './plugins/vuetify';

Vue.use(VueHead, {
  separator: '-'
})
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App),
  head: {
    meta: [
      {name: 'description', content: 'デフォルトのdescriptionです。'},
    ]
  }
}).$mount('#app')
