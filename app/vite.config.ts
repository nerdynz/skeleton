import { fileURLToPath, URL } from 'node:url'
import path from 'path'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import ReactivityTransform from '@vue-macros/reactivity-transform/vite'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

const iconPath = path.resolve(process.cwd(), 'src/assets/icons')
console.log('iconPath', iconPath)

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    createSvgIconsPlugin({
      // Specify the icon folder to be cached
      iconDirs: [iconPath],
      // Specify symbolId format
      symbolId: 'icon-[dir]-[name]',
    }),
    vueJsx(),
    vueDevTools(),
    ReactivityTransform(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },

  server: {
    open: true,
    proxy: {
      // string shorthand
      '/upload': 'http://localhost:8080/',
      '/api': 'http://localhost:8080/',
      '/twirp': 'http://localhost:8080/',
      '/static': 'http://localhost:8080/',
      // // with options
      // '/api': {
      //   target: 'http://jsonplaceholder.typicode.com',
      //   changeOrigin: true,
      //   rewrite: (path) => path.replace(/^\/api/, '')
      // },
    },
  },
  // optimizeDeps: {
  //   include: [
  //     'google-protobuf'
  //   ]
  // }
  css: {
    preprocessorOptions: {},
  },
})
