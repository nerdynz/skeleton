import { createSSRApp } from "vue";
import {
  createMemoryHistory,
  createRouter,
  createWebHistory,
} from "vue-router";
import { routes } from "vue-router/auto-routes";
import App from "./App.vue";
import { DataLoaderPlugin } from 'unplugin-vue-router/data-loaders'
import { createHead } from '@unhead/vue'

import "./styles/main.css";

const isServer = typeof window === "undefined";
export function makeApp(context: any) {
  // (globalThis as any).payload = context
  const app = isServer ? createSSRApp(App) : createApp(App)
  const router = createRouter({
    history: isServer
      ? createMemoryHistory()
      : createWebHistory(import.meta.env.BASE_URL),
    routes,
  });

  app.use(DataLoaderPlugin, { router }) 
  const head = createHead()
  app.use(head)
  app.use(router);
  return { app, router, head };
}
