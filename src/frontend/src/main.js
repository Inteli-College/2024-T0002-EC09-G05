import './index.css'

import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import Toaster from '@meforma/vue-toaster'
import { createSSRApp } from "vue";


import App from './App.vue'
import router from './router'

const app = createSSRApp(App)
app.use(Toaster)
//app.use(VueCookies)
app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
  unstyled: true,
  pt: "Tailwind"    
})

app.mount('#app')
