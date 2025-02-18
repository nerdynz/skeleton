<template>
  <building-block title="Hero Video" :block="block">
    <div class="columns is-multiline">
      <div class="column is-full">
        <div class="content is-clearfix">
          <image-crop :width="1920" :height="1080" output="jpg" force-jpeg v-model="block.imageOne" />
          <b-input label="Embed (Vimeo or Youtube)" name="Embed" type="textarea"  v-model="block.contentOneHtml" />
        </div>
        <!-- <animation-selector v-model="block.Additional"></animation-selector>  -->
      </div>
    </div>
  </building-block>
</template>

<script>
import BuildingBlock from './BuildingBlock.vue'
//import AnimationSelector from './AnimationSelector.vue'
let meta = {
  name: 'Hero Video',
  name: 'hero-video',
  svg: `
<svg viewBox="0 0 960 400" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <g id="background">
            <use fill="#F7F7F7" fill-rule="evenodd" xlink:href="#path-1"></use>
            <rect stroke="currentColor" stroke-width="2" x="1" y="1" width="958" height="398" rx="4"></rect>
        </g>
        <path d="M251,343 L708.957398,343 L708.957398,57 L251,57 L251,343 Z M456.109587,153.648436 L456.109587,208.818923 L500.552319,181.236581 L456.109587,153.648436 Z M405.806914,294.524932 L665.907129,294.524932 L405.806914,294.524932 Z M350.688644,294.524932 L377.783625,294.524932 L350.688644,294.524932 Z M384.571875,313.961375 L384.571875,275.094291 L384.571875,313.961375 Z M320.361991,313.961375 L320.361991,275.094291 L320.361991,313.961375 Z M292.321297,313.961375 L292.321297,275.094291 L292.321297,313.961375 Z M525.906723,145.06159 C532.636954,155.493447 536.524242,167.903761 536.524242,181.236581 C536.524242,218.183227 506.586319,248.132754 469.628068,248.132754 C432.669818,248.132754 402.731894,218.183227 402.731894,181.236581 C402.731894,144.284132 432.669818,114.334605 469.628068,114.334605 C485.467319,114.334605 499.972126,119.823224 511.459934,129.007668 L525.906723,145.06159 Z" id="Stroke-354" stroke="#484751" stroke-width="10.632"></path>
    </g>
</svg>
  `,
  category: 'Video' // basic, interactive, form
}
export default {
  // COMPONENT
  // ______________________________________
  meta: meta,
  components: {
    BuildingBlock,
//    AnimationSelector
  },
  props: {
    block: Object
  },
  computed: {
    name () {
      return meta.name
    }
  },
  methods: {
  },
  // watch: {},
  data () {
    return {
      isInit: true
    }
  },

  // LIFECYCLE METHODS
  // ______________________________________
  beforeCreate () {
  },
  created () {
  },
  beforeMount () {
  },
  mounted () {
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
  },
  destroyed () {
  }
}
</script>
