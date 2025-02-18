<template>
  <building-block :title="name" :block="block">
    <div class="columns">
      <div class="column is-4">
        <div class="content is-clearfix">
          <rich-text v-model="block.contentOneHtml"></rich-text>
        </div>
         <!-- <animation-selector v-model="block.Additional"></animation-selector>  -->
      </div>
      <div class="column is-4">
        <div class="content is-clearfix">
          <rich-text v-model="block.contentTwoHtml"></rich-text>
        </div>
        <!-- <animation-selector v-model="block.AdditionalTwo"></animation-selector>  -->
      </div>
      <div class="column is-4">
        <div class="content is-clearfix">
          <rich-text v-model="block.contentThreeHtml"></rich-text>
        </div>
        <!-- <animation-selector v-model="block.AdditionalThree"></animation-selector>  -->
      </div>
    </div>
  </building-block>
</template>

<script>
import BuildingBlock from './BuildingBlock.vue'
//import AnimationSelector from './AnimationSelector.vue'

let meta = {
  name: 'Three Column',
  kind: 'three-column',
  svg: `
<svg viewBox="0 0 960 200" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <g id="background">
            <use fill="#F7F7F7" fill-rule="evenodd" xlink:href="#path-1"></use>
            <rect stroke="#979797" stroke-width="2" x="1" y="1" width="958" height="198" rx="4"></rect>
        </g>
        <g id="Group" transform="translate(20.000000, 22.000000)" fill="#D8D8D8">
            <g id="content" transform="translate(3.061674, 0.000000)">
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="0" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="34" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="68" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="102" width="289.938326" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="136" width="289.938326" height="15"></rect>
            </g>
            <text id="Content" fill="currentColor" font-family="Helvetica-Bold, Helvetica" font-size="72" font-weight="bold">
                <tspan x="3.53112954" y="69">Aa</tspan>
            </text>
        </g>
        <g id="Group" transform="translate(332.000000, 22.000000)" fill="#D8D8D8">
            <g id="content" transform="translate(3.061674, 0.000000)">
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="0" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="34" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="68" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="102" width="289.938326" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="136" width="289.938326" height="15"></rect>
            </g>
            <text id="Content" fill="currentColor" font-family="Helvetica-Bold, Helvetica" font-size="72" font-weight="bold">
                <tspan x="3.53112954" y="69">Aa</tspan>
            </text>
        </g>
        <g id="Group" transform="translate(647.000000, 22.000000)" fill="#D8D8D8">
            <g id="content" transform="translate(3.061674, 0.000000)">
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="0" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="34" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="116.278996" y="68" width="173.65933" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="102" width="289.938326" height="15"></rect>
                <rect id="Rectangle-2" fill="currentColor" x="0" y="136" width="289.938326" height="15"></rect>
            </g>
            <text id="Content" fill="currentColor" font-family="Helvetica-Bold, Helvetica" font-size="72" font-weight="bold">
                <tspan x="3.53112954" y="69">Aa</tspan>
            </text>
        </g>
    </g>
</svg>
  `,
  category: 'Basic' // basic, interactive, form
}

export default {
  // COMPONENT
  // ______________________________________
  name: 'ThreeColumn',
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
