<template>
  <with-actions
    id="page-edit"
    :title="pageTitle"
    :actions="actions"
    class="has-scroll-body"
    header-class="p-4"
    footer-class="px-4 pb-2"
  >
    <form v-if="page" class="px-4">
      <section class="container is-relative">
        <o-tabs type="toggle" position="right">
          <o-tab-item label="Settings" icon="fad-cog">
            <div class="columns is-multiline mt-2">
              <div class="column is-8">
                <FormField
                  class=""
                  :validation="validation"
                  @changed="validate"
                  for="title"
                  label="Title"
                >
                  <o-input label="Title" name="title" type="text" v-model="page.title" />
                </FormField>
              </div>
              <div class="column is-4">
                <FormField
                  class=""
                  :validation="validation"
                  @changed="validate"
                  for="slug"
                  label="Slug"
                >
                  <o-input label="Slug" name="slug" type="text" v-model="page.slug" />
                </FormField>
              </div>
              <div class="column is-12">
                <FormField
                  class=""
                  :validation="validation"
                  @changed="validate"
                  for="summary"
                  label="Summary"
                >
                  <o-input label="Summary" name="summary" type="textarea" v-model="page.summary" />
                </FormField>
              </div>
            </div>
          </o-tab-item>
          <o-tab-item label="Blocks" icon="fad-cubes">
            <div class="master-page-template">
              <draggable
                v-model="page.blocks"
                item-key="blockUlid"
                @start="dragStart"
                @end="dragEnd"
              >
                <template #item="{ element }">
                  <component :is="element.kind" :block="element" @remove="removeBlock(element)" />
                </template>
              </draggable>
              <o-button class="is-primary add-block-btn is-fullwidth my-2" @click="modalOpen"
                >Add Block</o-button
              >
            </div>
          </o-tab-item>
        </o-tabs>
      </section>
    </form>

    <o-modal
      title="Select a block"
      min-width="60vw"
      animation=""
      class="block-selector-overlay"
      :is-visible="isBlockSelectorVisible"
      @closed="modalClosed"
      scroll="keep"
    >
      <div class="block-selector">
        <div class="columns is-multiline u-p">
          <template v-for="({ name, kind, svg, category }, index) in blocksMeta" :key="index">
            <div
              class="column is-12 -u-p"
              v-if="index === 0 || (index > 0 && blocksMeta[index - 1].category !== category)"
              :key="category"
            >
              <h5 class="title is-5">{{ category }}</h5>
            </div>
            <div class="inner-block column is-6 u-rel" @click="addBlock(kind)">
              <div class="block-label">{{ name }}</div>
              <div v-html="typeof svg === 'function' ? svg() : svg"></div>
            </div>
          </template>
        </div>
      </div>
    </o-modal>

    <loading :on="['Load']" />
  </with-actions>
</template>

<script lang="ts">
import {
  OneColumn,
  TwoColumn,
  ThreeColumn,
  OneColumnHero,
  HeroImage,
  ImageOffsetLeft,
  ImageOffsetRight,
  HeroVideo,
  ImageLinks,
  ImageBlocks,
} from '@/components/blocks'
import { BlockWithImage } from '@/api/pb/block.pb'
import Draggable from 'vuedraggable'
import { DeletePage, LoadFullPage, type PageWithBlockWithImage } from '@/api/pb/page.pb'

export default {
  components: {
    OneColumn,
    TwoColumn,
    ThreeColumn,
    OneColumnHero,
    HeroImage,
    ImageOffsetLeft,
    ImageOffsetRight,
    HeroVideo,
    ImageLinks,
    ImageBlocks,
    Draggable,
  },
}
</script>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { TYPE, useToast } from 'vue-toastification'
const toast = useToast()

const isBlockSelectorVisible = ref(false)

const route = useRoute()
const router = useRouter()

const isNew = computed(() => {
  return !route.params?.ulid
})

const ulid = $computed(() => {
  return route.params?.ulid as string
})

const pageTitle = computed(() => {
  return `${isNew.value ? 'Create' : 'Edit'} Page`
})

// const isValid = computed(() => {
//   return Object.keys(validation).length === 0
// })

async function load() {
  // page.value = null
  // if (route.params?.ulid) {
  //   // EDIT
  //   page.value = await api.page.loadFullPage({
  //     ulid: route.params?.ulid as string,
  //   })
  // } else {
  //   page.value = api.page.createPage() as PageWithBlockWithImage
  // }
}
let page: PageWithBlockWithImage = await LoadFullPage({ ulid: ulid })

// let validation: Ref<Validation> = ref({
//   isValid: false,
// })

const actions = computed(() => {
  return [
    {
      name: isNew.value ? 'Create' : 'Save',
      action: async () => {
        let saved = await save()
      },
    },
    {
      name: 'Cancel',
      class: 'is-light',
      action: () => {
        router.replace({ name: 'PageList' })
      },
    },
    {
      name: 'Delete',
      class: 'is-danger is-right',
      action: () => {
          DeletePage({
            ulid: route.params?.ulid as string,
          })
          .then(() => {
            toast('Page Deleted', {
              type: TYPE.ERROR,
            })
            router.replace({ name: 'PageList' })
          })
      },
    },
  ]
})

// async function save(): Promise<boolean> {
//   let newPage = await api.page.saveFullPage(page.value!)
//   page.value = newPage

//   router.replace({ name: 'PageEdit', params: { ulid: newPage.pageUlid } })
//   return true
// }

// onMounted(() => {
//   load()
// })

// // blocks
// function addBlock(kind: string) {
//   console.log('kind', kind)
//   let block = api.block.createBlock() as unknown as BlockWithImage
//   block.kind = kind
//   block.contentOneHtml = '<p>asdf</p>'
//   if (page && page.value) {
//     block.pageUlid = page.value.pageUlid
//     console.log('block', block)
//     page.value.blocks.push(block)
//   }
//   isBlockSelectorVisible.value = false
//   resetBlockSortPosition()
// }

// function modalOpen() {
//   isBlockSelectorVisible.value = true
// }

// function modalClosed() {
//   isBlockSelectorVisible.value = false
// }

// const blocksMeta = computed(() => {
//   let metas: any = []
//   const blocks: any = {
//     OneColumn,
//     TwoColumn,
//     ThreeColumn,
//     OneColumnHero,
//     HeroImage,
//     ImageOffsetLeft,
//     ImageOffsetRight,
//     // HeroVideo,
//     // ImageLinks,
//     // ImageBlocks,
//   }
//   Object.keys(blocks).forEach((key, index) => {
//     let block: any = blocks[key]
//     metas.push(block.meta)
//   })
//   return metas
// })

// function removeBlock(deletingBlock: BlockWithImage) {
//   let blocks = page.value?.blocks

//   if (page && blocks && blocks.length) {
//     page.value!.blocks = blocks.filter((b) => {
//       return b.blockUlid !== deletingBlock.blockUlid
//     })
//   }
// }

// function dragStart(e: any) {}

// function dragEnd(e: any) {
//   console.log('this.blocks', page.value?.blocks)
//   resetBlockSortPosition()
// }

// function resetBlockSortPosition() {
//   let blocks = page.value?.blocks
//   if (blocks && blocks.length > 0) {
//     blocks.forEach((block, i) => {
//       block.sortPosition = i * 100
//     })
//   }
// }
</script>
<!-- <style lang="scss">
#page-edit {
  .tab-content {
    padding: 0;
  }

  .add-block-btn {
    position: sticky;
    top: 1px;
    right: 0;
    z-index: 21;
  }

  .b-tabs {
    .tabs {
      position: sticky;
      top: 0px;
      border-bottom: none;
      z-index: 20;
      background: $body-background-color;

      &.is-centered {
        justify-content: center;
      }
    }
  }
}
</style> -->
