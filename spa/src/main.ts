import Componenty from '@nerdynz/componenty/src/index';
import { titleCase } from '@nerdynz/componenty/src/utils/formatters';

import Oruga from '@oruga-ui/oruga-next';

import { bulmaConfig } from '@oruga-ui/theme-bulma';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import 'virtual:svg-icons-register';
import { createApp, nextTick } from 'vue';
import App from './App.vue';
import './assets/base.scss';
import router from './router';

import $api from '@/api';
import { createPinia } from 'pinia';
import { client, TwirpError } from 'twirpscript';
import { useUserAccessStore } from './store/userAccess';
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)
app.use(pinia) // pinia must come first


// ROUTER
app.use(router)
router.afterEach((to, from) => {
  nextTick( () => {
    document.title =`${(to.meta.title ? to.meta.title : '***')} - ${titleCase('skeleton')}` ;
  });
})

// app.use(store)
app.use(Oruga, {
  ...bulmaConfig, 
  ...{
    iconComponent: 'icon',
    iconPack: 'far'
  }
})

client.on("error", (context: any, err: TwirpError) => {
  // log or report
  console.log('context', context)
  if (err.code === 'unauthenticated') {
    router.push({ name: 'login' })
  }
});
client.on("error", (context: any, err: TwirpError) => {
  // log or report
  console.log('context', context)
  if (err.code === 'unauthenticated') {
    router.push({ name: 'login' })
  }
});

app.use(Componenty, {
  twirpClient: client
})

app.use({
  install: (app, store) => {
    app.config.globalProperties.$api = $api
  },
})


const userAccess =  useUserAccessStore()
client.use((context, next) => {
  const auth = userAccess.details.token
  if (auth) {
    console.log('hello', auth)
    context.headers['Authorization'] = `Basic ${auth}`
  }
  return next(context)
})

app.mount('#app')

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    // $store: typeof store
    $api: typeof $api
  }
}
