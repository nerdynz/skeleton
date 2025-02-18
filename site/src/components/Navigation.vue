<template>
  <nav class="nav is-fixed">
    <div class="nav-left"></div>
    <div class="nav-center">
      <RouterLink to="/" class="nav-item logo-item">
        <img :src="`/assets/images/Rewind.png`" alt="Relapse Logo" />
        &nbsp;&nbsp;Relapse
      </RouterLink>
      
      <RouterLink v-for="ni in meta?.mainNav" :to="ni.slug" class="nav-item logo-item">{{ ni.title }}</RouterLink>
    </div>
    <div class="nav-right">
      <div class="nav-item"></div>
      
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useSSRMetadata } from "~/composables/ssr";

let meta = useSSRMetadata()

</script>
