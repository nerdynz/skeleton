import 'virtual:svg-icons-register'

import Satchel from '@nerdynz/satchel'
import Oruga from '@oruga-ui/oruga-next'
import { createApp } from 'vue'

import Toast from 'vue-toastification'
// Import the CSS or use your own!

import { bulmaConfig } from '@oruga-ui/theme-bulma'

import '@oruga-ui/theme-bulma/dist/bulma.css'
import 'vue-toastification/dist/index.css'

import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import '@nerdynz/satchel/dist/style.css'
import './assets/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Satchel)
app.use(Oruga, bulmaConfig)

app.use(Toast)

app.mount('#app')
