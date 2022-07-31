import vue from '@vitejs/plugin-vue'
import path from 'path'
import { fileURLToPath, URL } from 'url'
import { defineConfig } from 'vite'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

const iconPath = path.resolve(process.cwd(), 'src/assets/icons')
console.log('iconPath', iconPath)
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      script: {
        refSugar: true,
      },
    }),
    createSvgIconsPlugin({
      // Specify the icon folder to be cached
      iconDirs: [iconPath],
      // Specify symbolId format
      symbolId: 'icon-[dir]-[name]',
    }),
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
      '/api': 'http://localhost:4500/',
      '/twirp': 'http://localhost:4500/',
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
    preprocessorOptions: {
      scss: {
        additionalData: `@import "@/assets/variables.scss";`,
      },
    },
  },
})
