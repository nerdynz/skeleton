<template>
  <span class="building-block-description-bar">
    <span class="desc-text">{{ value }}</span>
    <button v-if="removeable" @click.stop="removed" class="delete" aria-label="close"></button>
  </span>
</template>

<script lang="ts" setup>
const props = defineProps({
  block: Object,
  value: String,
  removeable: Boolean,
  isHiddenSort: Boolean,
})
const emit = defineEmits<{
  (e: 'edit'): void
  (e: 'removed'): void
}>()

function edit() {
  emit('edit')
}
function removed() {
  emit('removed')
}
</script>

<style lang="scss">
.building-block-description-bar {
  position: absolute;
  top: 0.15rem;
  left: 0;
  right: 0.65rem;
  font-size: 0.6rem;
  color: var(--grey);
  transform: translate3d(0.35rem, 0.2rem, 0);

  .delete {
    position: absolute;
    right: 0;
    top: 0;
  }
}
</style>
