<script lang="ts">
import { defineBasicLoader } from "unplugin-vue-router/data-loaders/basic";
export const usePageData = defineBasicLoader("/[slug]", async (route) => {
  return ssrFetch(`/${route.params.slug}`)
});
</script>

<script setup lang="ts">
import { useHead } from "@unhead/vue";
useHead({
  title: "My awesome site",
});

const route = useRoute();

const {
  data, // the data returned by the loader
  isLoading, // a boolean indicating if the loader is fetching data
  error, // an error object if the loader failed
  reload, // a function to refetch the data without navigating
} = usePageData()

</script>

<template>
  <div id="home-page" class="">
    {{ data }}
  </div>
</template>
