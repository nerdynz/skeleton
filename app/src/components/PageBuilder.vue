<template>
  <div
    ref="pageViewer"
    v-if="!disabled"
    class="master-page-template"
    :class="{ disabled: disabled }"
  >
    <div class="blocks-wrapper">
      <div class="blocks">
        <div
          class="block-outer"
          v-for="block in filteredBlocks"
          :key="block.UUID"
        >
          <component
            v-bind:is="fmtKebab(block.Type)"
            :block="block"
            @edit="editBlock"
            @move="moveBlock"
            @removed="removeBlock"
          >
          </component>
        </div>
      </div>
      <div class="blocks-overview">
        <button
          class="add-block button is-success u-fw"
          @click="openBlockSelector"
        >
          Add Block
        </button>
        <div class="block-scroller">
          <template v-for="block in filteredBlocks">
            <drop
              :key="block.UUID"
              @drop="
                (xferData) => {
                  dropBlock(block, xferData)
                }
              "
            >
              <drag :transfer-data="block">
                <building-block
                  :title="block.Type"
                  class="block-outer"
                  :key="block.UUID"
                  :block="block"
                  @edit="editBlock"
                  @move="moveBlock"
                  @removed="removeBlock"
                  is-hidden-sort
                  is-top-level
                >
                  <div
                    class="inner-block"
                    v-html="getBlockMetaSvg(block)"
                  ></div>
                </building-block>
              </drag>
            </drop>
          </template>
        </div>
      </div>
    </div>
    <b-modal
      animation=""
      class="block-selector-overlay"
      :active.sync="isBlockSelectorVisible"
      scroll="keep"
    >
      <div class="block-selector">
        <div class="modal-card" style="width: auto; height: 80vh">
          <header class="modal-card-head">
            <p class="modal-card-title">Select a layout</p>
          </header>
          <section class="modal-card-body">
            <div class="columns is-multiline u-p">
              <template
                v-show="
                  typeof isDevOnly === 'undefined' || (isDevOnly && userIsDev)
                "
                v-for="(
                  { name, svg, category, isDevOnly }, index
                ) in blocksMeta"
              >
                <div
                  class="column is-12 -u-p"
                  v-if="
                    index === 0 ||
                    (index > 0 && blocksMeta[index - 1].category !== category)
                  "
                  :key="category"
                >
                  <h5 class="title is-5">{{ category }}</h5>
                </div>
                <div
                  class="inner-block column is-4 u-rel"
                  :key="index"
                  @click="addBlock(name)"
                >
                  <div class="block-label">{{ name }}</div>
                  <div v-html="typeof svg === 'function' ? svg() : svg"></div>
                </div>
              </template>
            </div>
          </section>
        </div>
      </div>
    </b-modal>
  </div>
</template>

<script>
// import ContactFormRight from '~/components/blocks/ContactFormRight'
import * as Blocks from '~/components/blocks'
import BuildingBlock from '~/components/blocks/BuildingBlock'
import {
  byField,
  changeSort,
  changeSortByUUID,
  indexByField,
} from '~/helpers/filters.js'
import { fmtKebab } from '~/helpers/format.js'
export default {
  // COMPONENT
  // ______________________________________
  name: 'MasterPageTemplate',
  components: {
    ...Blocks,
    BuildingBlock,
    // ContactFormRight
  },
  props: {
    disabled: Boolean,
    page: Object,
  },
  computed: {
    userIsDev() {
      return false
    },
    ...mapGetters({
      offsetWindowHeight: 'src/offsetWindowHeight',
    }),
    filteredBlocks() {
      let blocks = this.page.Blocks || []
      blocks = blocks
        .filter((block) => {
          return !block.IsDeleted
        })
        .sort((a, b) => {
          return a.SortPosition - b.SortPosition
        })
      return blocks
    },
    basicBlocks() {
      return this.blocksMeta.filter((block) => {
        return block.type === 'basic'
      })
    },
    blocksMeta() {
      var metas = []
      Object.keys(Blocks).forEach((key, index) => {
        var block = Blocks[key]
        metas.push(block.meta)
      })
      return metas
    },
  },
  methods: {
    fmtKebab,
    getBlockMeta(block) {
      let meta = byField(this.blocksMeta, 'name', block.Type)
      return meta
    },
    getBlockMetaSvg(block) {
      var meta = this.getBlockMeta(block)
      if (meta) {
        return typeof meta.svg === 'function' ? meta.svg() : meta.svg
      }
      return ''
    },
    addBlock(type) {
      if (this.currentEditingBlock) {
        let index = indexByField(
          this.page.Blocks,
          'UUID',
          this.currentEditingBlock
        )
        this.page.Blocks[index].Type = type
      } else {
        var current = this.filteredBlocks
        let newBlock = this.$service.new('Block')
        newBlock.Type = type
        if (current.length > 0) {
          newBlock.SortPosition = current[current.length - 1].SortPosition + 10
        } else {
          newBlock.SortPosition = 10
        }
        if (!this.page.Blocks) this.page.Blocks = []
        this.page.Blocks.push(newBlock)
      }

      this.isBlockSelectorVisible = false
      this.$nextTick(() => {
        this.$refs.pageViewer.scrollTop = this.$refs.pageViewer.scrollHeight
      })
    },
    removeBlock(uuid) {
      let i = this.page.Blocks.findIndex((block) => {
        return block.UUID === uuid
      })
      this.page.Blocks[i].IsDeleted = true
    },
    openBlockSelector() {
      this.$refs.pageViewer.scrollTop = 0
      this.isBlockSelectorVisible = !this.isBlockSelectorVisible
    },
    editBlock(uuid) {
      this.currentEditingBlock = uuid
      this.openBlockSelector()
    },
    moveBlock({ uuid, direction }) {
      console.log(uuid, direction)
      let sortedBlocks = this.filteredBlocks
      let index = indexByField(sortedBlocks, 'UUID', uuid)
      if (index === 0 && direction === 'up') {
        return
      }
      if (index === sortedBlocks.length - 1 && direction === 'down') {
        return
      }
      let newIndex = direction === 'up' ? index - 1 : index + 1
      this.page.Blocks = changeSort(sortedBlocks, index, newIndex)
    },
    dropBlock(a, b) {
      this.changeBlockSort(a, b)
    },
    changeBlockSort(a, b) {
      // this.$store.commit(BLOCK_HOVERED, false)
      // let blocks = []
      // let sortEls = document.getElementsByClassName('block-sort-pos')
      // if (sortEls && sortEls.length > 0) {
      //   for (let i = 0; i < sortEls.length; i++) {
      //     let uuid = sortEls[i].getAttribute('uuid')
      //     let block = this.page.Blocks[indexByField(this.page.Blocks, 'UUID', uuid)]
      //     block.SortPosition = (i * 10)
      //     blocks.push(block)
      //   }
      // }
      this.page.Blocks = changeSortByUUID(this.page.Blocks, a.UUID, b.UUID)
    },
    isKind(kind) {
      kind = kind.toLowerCase()
      let specialFor = '' + this.page.SpecialPageFor
      let pageKind = '' + this.page.Kind

      if (specialFor.indexOf(':') > 0) {
        let s = specialFor.split(':')
        pageKind = s[0]
        specialFor = s[1]
      }
      return (
        pageKind.toLowerCase() === kind || specialFor.toLowerCase() === kind
      )
    },
  },
  watch: {
    isBlockSelectorVisible(isVis) {
      if (!isVis) {
        this.currentEditingBlock = null
      }
    },
  },
  data() {
    return {
      isBlockSelectorVisible: false,
      currentEditingBlock: null,
    }
  },

  // LIFECYCLE METHODS
  // ______________________________________
  beforeCreate() {},
  created() {},
  beforeMount() {},
  mounted() {},
  beforeUpdate() {},
  updated() {},
  beforeDestroy() {},
  destroyed() {},
}
</script>

<style lang="scss">
// @import '@bulma/utilities/_all';
// @import '@bulma/elements/button';

.master-page-template {
  padding: 0rem;
  &.disabled {
    background: $grey-lighter;
    &::before {
      content: 'CUSTOM LAYOUT';
      color: $grey-light;
    }
  }

  .column {
    position: relative;
  }

  .animation-button {
    position: absolute;
    bottom: -1rem;
    right: 0.75rem;
    z-index: 50;

    .is-static {
      color: $black;
      background: $grey-light;
      border-color: $grey-light;
    }
  }

  .blocks-wrapper {
    height: calc(100vh - 7rem);
    display: flex;
    padding: 0 1rem 1rem;

    .blocks {
      width: calc(100% - 270px);
      height: 100%;
      overflow-y: scroll;
      overflow-x: hidden;
      .block-outer {
        margin: 0;
        &:nth-child(odd) {
          background-color: $grey-lightest;
        }
        &:nth-child(even) {
          background-color: $white;
        }
      }
    }

    .blocks-overview {
      position: relative;
      background: $white;
      width: 270px;
      padding-bottom: 2.5rem;
      .block-scroller {
        height: 100%;
        overflow-x: none;
        overflow-y: scroll;
      }

      .add-block {
        position: absolute;
        bottom: 0;
        border-radius: 0;
      }

      .building-block {
        padding: 1.5rem 0.25rem 0.15rem;
        min-height: 40px;
        cursor: move; /* fallback if grab cursor is unsupported */
        cursor: grab;
        cursor: -moz-grab;
        cursor: -webkit-grab;

        &:active {
          cursor: grabbing;
          cursor: -moz-grabbing;
          cursor: -webkit-grabbing;
        }
      }

      .inner-block {
        padding: 0;
        cursor: inherit;
      }
    }
  }

  .page-title {
    font-weight: bold;
    background: $grey-lighter;
    color: $grey-darker;
    position: absolute;
    height: 2rem;
    z-index: 10;
    top: 0;
    left: 0;
    right: 0;
    padding: 0.25rem 0.5rem;

    .preview-button {
      top: 0;
      bottom: 0;
      position: absolute;
      right: 0;
      color: $white;
      font-size: $size-small;
      border-radius: 0;
    }
  }

  .block-selector-overlay {
    background: $grey-lighter;
    border-top: 1px solid $white;
    // margin-top: 2rem;
    z-index: 1100;
    // overflow: hidden;
    // overflow-y: scroll;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;

    .block-category {
      font-size: $size-5;
      color: $grey-dark;
      border-bottom: 1px solid transparentize($grey, 0.5);
      margin-bottom: 0.5rem;
    }

    .modal-content {
      width: 100%;
    }
  }

  .building-block {
    // border: 1px dashed $grey;
    margin: 0;
    padding: 0;
    position: relative;
    min-height: 100px;
    padding: 2rem 0.5rem 2.5rem;

    // margin-bottom: 1rem;
    border-radius: $radius;

    // .content {
    //   height: 100%;
    // }

    .is-custom {
      display: flex;
      justify-content: center;
      flex-direction: column;
      text-align: center;
      min-height: 300px;
      margin: -0.5rem;
      font-size: $size-1;
    }

    .fr-element {
      // border: 1px dashed $grey;
      &.fr-view {
        min-height: 200px;
      }
    }

    .block-placeholder {
      min-height: 140px;
      text-transform: uppercase;
      color: $grey-light;
      font-size: $size-1;
      padding: 2rem;
      display: flex;
      justify-content: center;
      flex-direction: column;

      &.is-full-page {
        padding-top: 16rem;
        padding-bottom: 16rem;
      }
    }

    &.hero-image-slider {
      background-color: $grey-lighter;
      .block-image {
        .slim {
          border-radius: $radius;
          overflow: hidden;
          background-color: $white;
        }
      }
    }
  }

  .description {
    overflow: hidden;
    position: absolute;
    z-index: 20;
    right: 0;
    top: 0;
    left: 0;
    font-size: $size-small;
    background-color: $grey;
    color: $white;
    font-weight: $weight-semibold;
    padding: 0;
    border-top-left-radius: 0;
    border-bottom-right-radius: 0;
    user-select: none;
    // cursor: pointer;

    .fa {
      margin-top: 4px;
      margin-left: 2px;
      margin-right: 2px;
    }

    a {
      color: $grey-dark;
      font-weight: $weight-semibold;
      display: inline-block;
    }

    .desc-text {
      display: inline-block;
      padding-top: 2px;
      padding-right: 8px;
    }

    .desc-button {
      border-left: 1px solid $grey-dark;
      padding: 0;
      display: block;
      width: 30px;
      float: right;
      text-align: center;
      padding: 2px;

      &.is-danger {
        border-left: 1px solid $red;
        background-color: $red;
        color: $white;
      }

      &:hover {
        background-color: lighten($grey-dark, 3%);
        color: $white;
        border-left: 1px solid lighten($grey-dark, 3%);
      }
    }
  }

  .inner-block {
    cursor: pointer;
    padding: 0.5rem;
    position: relative;

    .block-label {
      color: $grey-dark;
      text-transform: uppercase;
      font-size: 0.75rem;
      padding-left: 0.25rem;
      font-weight: $weight-semibold;
    }

    .label {
      margin-bottom: 0;
    }
    svg {
      transform-origin: 0 0 0;
      width: 100%;
      max-width: 800px;
      #star {
        fill: $grey-dark;
      }
      #content {
        fill: $grey-dark;
      }
      #background {
        rect {
          stroke: $grey-dark;
          // fill: $button-background-color;
        }
      }
    }
    &:active {
      .block-label {
        color: darken($grey-darker, 2.5);
      }
      svg {
        #content {
          // fill: $button-active-color;
        }
        #background {
          rect {
            // stroke: $button-active-border-color;
          }
        }
      }
    }
    &:hover {
      .block-label {
        color: darken($grey-darker, 2.5);
      }
      svg {
        #content {
          fill: darken($grey-darker, 2.5);
        }
        #background {
          rect {
            stroke: darken($grey-darker, 2.5);
          }
        }
      }
    }
  }
}
</style>
