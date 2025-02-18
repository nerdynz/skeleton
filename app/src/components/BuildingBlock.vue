<template>
  <div ref="blockOuter" class="building-block block-sort-pos">
    <description-bar class="has-close" :removeable="removeable" :value="title" @edit="edit" @removed="remove"
      :block="block" :is-hidden-sort="isHiddenSort" />
    <div ref="blockInner" :class="{ 'block-placeholder': custom }">
      <slot />
    </div>
  </div>
</template>
<script lang="ts" setup>
import DescriptionBar from './DescriptionBar.vue'

const props = defineProps({
  title: String,
  block: {
    type: Object,
    required: true,
  },
  removeable: {
    type: Boolean,
    default: true,
  },
  custom: Boolean,
  isTopLevel: Boolean,
  isHiddenSort: Boolean,
})

const emit = defineEmits<{
  (e: 'remove', ulid: string): void
  (e: 'edit', ulid: string): void
}>()

function remove() {
  emit('remove', props.block.blockUlid)
}
function edit() {
  emit('edit', props.block.blockUlid)
}
</script>
<style lang="scss">
.building-block {
  position: relative;
  border: 1px solid var(--bulma-border-color);
  margin-top: 1rem;
  border-radius: var(--bulma-radius);
  padding: 0.5rem;
  padding-top: 1.5rem;
  background-color: var(--white);
  // box-shadow: inset 0 1px 2px rgb(10 10 10 / 10%);

  .cropper {
    border-radius: var(--bulma-radius);
    overflow: hidden;
  }

  .ql-container.ql-snow,
  .ql-toolbar.ql-snow {
    border: none;
    padding: 0;
    // background-color: var(--white-ter);
    border-top-left-radius: var(--bulma-radius);
    border-top-right-radius: var(--bulma-radius);
    // padding-top: 0.5rem;
    // padding-left: 0.5rem;
    // padding-right: 0.5rem;
  }

  .ql-editor {
    padding: 0.25rem;
    border-bottom-left-radius: var(--bulma-radius);
    border-bottom-right-radius: var(--bulma-radius);
  }

  .ql-toolbar.ql-snow {
    margin-top: -5px;
    margin-left: -5px;
  }

  .content {
    margin: 0;
  }

  .animation-button {
    position: absolute;
    bottom: 0.5rem;
  }

  .upload {
    display: block;
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;

    .upload-draggable {
      &:hover {
        background-color: var(--white-bis);
      }

      border: none;
      text-align: center;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
