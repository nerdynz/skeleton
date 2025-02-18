<template>
  <building-block :title="name" :block="block">
    <div class="columns">
      <div class="column is-6">
        <image-upload :width="940" :height="540" output="jpg" force-jpeg v-model="block.Picture"></image-upload>
        <control label="Embed (Vimeo or Youtube)" name="Embed" type="textarea"  v-model="block.PictureTwo" />
        <animation-selector v-model="block.Additional"></animation-selector> 
      </div>
      <div class="column is-6">
        <div class="content is-clearfix">
          <rich-text v-model="block.HTML"></rich-text>
        </div>
        <animation-selector v-model="block.AdditionalTwo"></animation-selector> 
      </div>
    </div>
  </building-block>
</template>

<script>
import BuildingBlock from './BuildingBlock.vue'
//import AnimationSelector from './AnimationSelector.vue'

let meta = {
  name: 'Video Left',
  svg: `
    <svg viewBox="0 0 960 200" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
        <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
            <g id="background">
                <use fill="#F7F7F7" fill-rule="evenodd" xlink:href="#path-1"></use>
                <rect stroke="#979797" stroke-width="2" x="1" y="1" width="958" height="198" rx="4"></rect>
            </g>
            <g id="Group" transform="translate(316.000000, 21.000000)" fill="#D8D8D8">
                <g id="content" transform="translate(10.000000, 0.000000)">
                    <rect id="Rectangle-2" fill="currentColor" x="103.008889" y="0" width="498.991111" height="15.8940397"></rect>
                    <rect id="Rectangle-2" fill="currentColor" x="103.008889" y="36.0264901" width="498.991111" height="15.8940397"></rect>
                    <rect id="Rectangle-2" fill="currentColor" x="103.008889" y="72.0529801" width="498.991111" height="15.8940397"></rect>
                    <rect id="Rectangle-2" fill="currentColor" x="0" y="108.07947" width="602" height="15.8940397"></rect>
                    <rect id="Rectangle-2" fill="currentColor" x="0" y="144.10596" width="602" height="15.8940397"></rect>
                </g>
                <text id="Content" fill="currentColor" font-family="Helvetica-Bold, Helvetica" font-size="72" font-weight="bold">
                    <tspan x="4.98046875" y="75">Aa</tspan>
                </text>
            </g>
            <path d="M22.8037489,181 L278.003692,181 L278.003692,21 L22.8037489,21 L22.8037489,181 Z M137.102492,75.0690551 L137.102492,105.933663 L161.868513,90.5029821 L137.102492,75.0690551 Z M109.070977,153.881081 L254.01363,153.881081 L109.070977,153.881081 Z M78.3559385,153.881081 L93.454805,153.881081 L78.3559385,153.881081 Z M97.2376046,164.754615 L97.2376046,143.010792 L97.2376046,164.754615 Z M61.4562007,164.754615 L61.4562007,143.010792 L61.4562007,164.754615 Z M45.8303287,164.754615 L45.8303287,143.010792 L45.8303287,164.754615 Z M175.997431,70.265225 C179.747899,76.1012294 181.914117,83.0440622 181.914117,90.5029821 C181.914117,111.172435 165.231001,127.927415 144.63576,127.927415 C124.040518,127.927415 107.357402,111.172435 107.357402,90.5029821 C107.357402,69.8302836 124.040518,53.0753033 144.63576,53.0753033 C153.462292,53.0753033 161.545197,56.1458595 167.946858,61.2840102 L175.997431,70.265225 Z" id="Stroke-354" stroke="#484751" stroke-width="5.63199997"></path>
        </g>
    </svg>
  `,
  category: 'Video' // basic, interactive, form
}

export default {
  // COMPONENT
  // ______________________________________
  name: 'ImageOffsetLeft',
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
