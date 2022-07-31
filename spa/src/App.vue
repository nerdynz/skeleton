<template>
  <tip ref="tooltip" />
  <aside v-if="!isLoginRoute" class="main-navigation">
    <nav class="navbar">
      <nav-link v-for="route in routes" :route="route" :key="route.path" />
    </nav>
  </aside>
  <router-view />
</template>

<script lang="ts">
// import Tooltip from '@/components/Tooltip.vue'
import { Options, Vue } from 'vue-class-component'
import { RouteRecord } from 'vue-router'
// import { RouteRecord } from 'vue-router'

@Options({
  components: {},
})
export default class App extends Vue {
  get isLoginRoute() {
    return this.$route.name === 'login' || this.$route.name === 'logout'
  }
  get routes() {
    return this.$router.getRoutes().filter((route: RouteRecord) => {
      return route.meta.inSidebar
    })
  }
}
</script>
<style lang="scss">
.main-navigation {
  .navbar {
    width: $sidebar-width;
    position: fixed;
    height: 100vh;
    // box-shadow: 1px 0px 5px 0 rgba(0, 0, 0, 0.1);
    border-right: 1px solid $grey-lighter;
    display: flex;
    flex-direction: column;
  }
}
.main-navigation + section {
  margin-left: $sidebar-width;
}
</style>

