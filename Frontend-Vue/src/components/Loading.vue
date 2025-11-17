<template>
  <v-overlay
    v-model="internalShow"
    class="align-center justify-center"
  >
    <v-progress-circular
      color="primary"
      size="64"
      indeterminate
    />
  </v-overlay>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  show: { type: Boolean, default: false },
})

const emit = defineEmits(['update:show'])

const internalShow = ref(props.show)

watch(() => props.show, (val) => {
  internalShow.value = val
})

watch(internalShow, (val) => {
  emit('update:show', val)
})
</script>
