import './index.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import Toaster from '@meforma/vue-toaster'

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(Toaster)
app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
  unstyled: true
})

app.mount('#app')
