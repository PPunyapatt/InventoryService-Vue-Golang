<template>
  <v-dialog v-model="internalDialog" max-width="400">
    <v-card>
      <!-- <v-card-title class="text-h6">
        {{ title }}
      </v-card-title> -->
      <v-toolbar :title="title" color="red" height="50"></v-toolbar>
      <v-card-text>
        {{ message }}
      </v-card-text>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn class="btn" @click="close">Ok</v-btn>
        
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  show: { type: Boolean, default: false },
  title: { type: String, default: 'Error' },
  message: { type: String, default: 'Something went wrong.' }
})

const emit = defineEmits(['update:show'])

const internalDialog = ref(props.show)

watch(() => props.show, (val) => {
  internalDialog.value = val
})

function close() {
  internalDialog.value = false
  emit('close')
}
</script>

<style scoped>
.btn {
    color: white;
    background-color: #3a82d3
}
</style>
