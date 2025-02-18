<template>
  <building-block :title="name" :block="block">
    <div class="columns">
      <div class="column is-6">
        <div class="content is-clearfix">
          <rich-text v-model="block.HTML"></rich-text>
        </div>
        <animation-selector v-model="block.Additional"></animation-selector> 
      </div>
      <div class="column is-6">
        <image-upload :width="940" :height="540" output="jpg" force-jpeg v-model="block.Picture"></image-upload>
        <control label="Embed (Vimeo or Youtube)" name="Embed" type="textarea"  v-model="block.PictureTwo" />
        <animation-selector v-model="block.AdditionalTwo"></animation-selector> 
      </div>
    </div>
  </building-block>
</template>

<script>
import BuildingBlock from './BuildingBlock.vue'
//import AnimationSelector from './AnimationSelector.vue'

let meta = {
  name: 'Video Right',
  svg: `
<svg viewBox="0 0 960 200" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <g id="background">
            <use fill="#F7F7F7" fill-rule="evenodd" xlink:href="#path-1"></use>
            <rect stroke="#979797" stroke-width="2" x="1" y="1" width="958" height="198" rx="4"></rect>
        </g>
        <g id="Group" transform="translate(16.000000, 21.000000)" fill="#D8D8D8">
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
        <path d="M680.803749,181 L937.003692,181 L937.003692,21 L680.803749,21 L680.803749,181 Z M795.550371,75.0690551 L795.550371,105.933663 L820.413438,90.5029821 L795.550371,75.0690551 Z M767.409015,153.881081 L912.919625,153.881081 L767.409015,153.881081 Z M736.57362,153.881081 L751.731651,153.881081 L736.57362,153.881081 Z M755.529273,164.754615 L755.529273,143.010792 L755.529273,164.754615 Z M719.60766,164.754615 L719.60766,143.010792 L719.60766,164.754615 Z M703.920558,164.754615 L703.920558,143.010792 L703.920558,164.754615 Z M834.59772,70.265225 C838.362884,76.1012294 840.537591,83.0440622 840.537591,90.5029821 C840.537591,111.172435 823.789102,127.927415 803.113158,127.927415 C782.437213,127.927415 765.688725,111.172435 765.688725,90.5029821 C765.688725,69.8302836 782.437213,53.0753033 803.113158,53.0753033 C811.974277,53.0753033 820.088855,56.1458595 826.5156,61.2840102 L834.59772,70.265225 Z" id="Stroke-354" stroke="#484751" stroke-width="5.63199997"></path>
    </g>
</svg>
  `,
  category: 'Video' // basic, interactive, form
}

export default {
  // COMPONENT
  // ______________________________________
  name: 'TwoColumn',
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
