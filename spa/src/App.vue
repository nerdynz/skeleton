<template>
  <aside v-if="!isLoginRoute" class="main-navigation">
    <nav class="navbar">
      <nav-link v-for="route in routes" :route="route" :key="route.path" />
    </nav>
  </aside>
  <main class="main-body">
    <section class="main-content">
      <RouterView v-slot="{ Component }">
        {{ cacheComponent(Component) }}
        <template v-if="Component || cachedComponent">
          <Suspense>
            <component :is="Component || cachedComponent" />
            <template #fallback>
              <div style="font-size: 1.5rem">Loading Demo Nested...</div>
            </template>
          </Suspense>
        </template>
      </RouterView>
    </section>
  </main>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRoute, useRouter, type RouteRecord } from 'vue-router'

const cachedComponent = ref(null)
function cacheComponent(c: any) {
  cachedComponent.value = c || cachedComponent.value
}

const route = useRoute()
const router = useRouter()

const isLoginRoute = $computed(() => {
  return route.name === 'login' || route.name === 'logout'
})

const routes = $computed(() => {
  return router.getRoutes().filter((route: RouteRecord) => {
    return route.meta.inSidebar
  })
})

</script>
<style lang="scss">
.main-navigation {
  .navbar {
    width: var(--satchel-sidebar-width, 5rem);
    position: fixed;
    height: 100vh;
    // box-shadow: 1px 0px 5px 0 rgba(0, 0, 0, 0.1);
    border-right: 1px solid var(--bulma-grey-lighter);
    display: flex;
    flex-direction: column;
  }
}
.main-navigation + section {
  margin-left: var(--satchel-sidebar-width, 5rem);
}
</style>
