import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import 'mdb-vue-ui-kit/css/mdb.min.css'
import { store } from './store/store'

createApp(App).use(store).use(router).mount('#app')
