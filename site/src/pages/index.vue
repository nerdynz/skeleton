<script lang="ts">
import { defineBasicLoader } from "unplugin-vue-router/data-loaders/basic";
export const usePageData = defineBasicLoader("/", async (route) => {
  return ssrFetch('/home');
});
</script>

<script setup lang="ts">
import { useHead } from "@unhead/vue";
useHead({
  title: "My awesome site",
});

const route = useRoute();
const router = useRouter();

const {
  data, // the data returned by the loader
  isLoading, // a boolean indicating if the loader is fetching data
  error, // an error object if the loader failed
  reload, // a function to refetch the data without navigating
} = usePageData();

</script>

<template>
  <section id="home-page">
    <section class="hero is-link is-fullheight-with-navbar is-transparent">
      <div class="hero-body">
        <figure
          class="logo image is-5by3 has-text-centered zoomInDown animated"
        >
          <img
            :src="`/assets/images/Rewind.png`"
            alt="Relapse Logo"
            class="logo animate"
          />
        </figure>
        <div class="logo-text content fadeIn animated animation-delay-500">
          <h1>Relapse</h1>
          <br />
          <h2>{{ data.summary }}</h2>
        </div>
      </div>
    </section>
    <div class="more">
      <div class="learn-more">
        <a href="#more">
          Learn More<br />
          <i class="fa fa-angle-down"></i>
        </a>
      </div>
    </div>
  </section>
</template>
