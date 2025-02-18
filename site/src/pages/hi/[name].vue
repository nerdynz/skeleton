<script lang="ts">
import { defineBasicLoader } from "unplugin-vue-router/data-loaders/basic";

export const usePageData = defineBasicLoader("/hi/[name]", async (route) => {
  return new Promise((resolve, reject) => {
    if (typeof window === "undefined") {
      // resolve({ hello: "ssr" });
      // const context = useSSRContext();
      if ((globalThis as any).payload) {
        resolve((globalThis as any).payload);
      }
    } else {
      resolve({ hello: "ssr" });
    }
  });
});
</script>
<script setup lang="ts">
const route = useRoute();
const router = useRouter();

const name = computed(() => {
  if (route.name === "/hi/[name]") {
    return route.params.name;
  }
  return "this didnt work";
});

const {
  data, // the data returned by the loader
  isLoading, // a boolean indicating if the loader is fetching data
  error, // an error object if the loader failed
  reload, // a function to refetch the data without navigating
} = usePageData();
</script>

<template>
  <div>
    <div i-carbon-pedestrian text-4xl inline-block />
    <p>Hi, {{ name }}</p>
    <p text-sm op50>
      <em>Dynamic route!</em>
    </p>
    <div v-if="isLoading">loading</div>
    <div v-else>
      {{ data }}
    </div>

    <div>
      <button class="btn m-3 text-sm mt-8" @click="router.back()">Back</button>
    </div>
  </div>
</template>
