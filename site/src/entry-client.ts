import { makeApp } from '~/main'

const { app, router, head } = makeApp({})
router.push(window.location.pathname)

router.isReady().then(() => {
  app.mount('#app', true)
})
