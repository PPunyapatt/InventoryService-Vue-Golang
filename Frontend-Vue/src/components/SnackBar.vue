<template>
    <v-snackbar
      v-model="internalShow"
      :timeout="2000"
      location="bottom  end"
      color="success"
      elevation="2"
      :multi-line="false"
    >
      {{ message }}
    </v-snackbar>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  show: { type: Boolean, default: false },
  message: {type: String, required: true}
})
const emit = defineEmits(['update:show'])
const internalShow = ref(props.show)
const message = ref(props.message)

watch(() => props.show, (val) => {
  internalShow.value = val
})

watch(() => props.message, (val) => {
  message.value = val
})

watch(internalShow, (val) => {
  if (!val) {
    emit('update:show', false)
  }
})

</script>