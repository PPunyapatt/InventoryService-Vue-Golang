<!-- <template>
    <div class="text-center pa-4">
    <v-dialog
      v-model="dialog"
      max-width="400"
      persistent
    >
      <v-card
        prepend-icon="mdi-map-marker"
        :text="text"
      >
        <template v-slot:actions>
          <v-spacer></v-spacer>
          
          <v-btn 
            @click="handleConfirm('confirm')" 
            class=" ms-4 text-white"
            color="blue-darken-4"
            variant="flat"
          >
            ตกลง
          </v-btn>

          <v-btn 
            @click="handleConfirm('cancel')"
            color="red" 
            class=" text-white"
            variant="flat"
          >
            ยกเลิก
          </v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { ref, watch } from 'vue'
export default {
  emits: ['confirm', 'cancel'],
  props: {
    open: { type: Boolean, required: true },
    text: { type: String, required: true }
  },
  setup(props, { emit }) {
    const dialog = ref(props.open)
    const text = ref(props.text)
    // dialog.value = props.open
    watch(() => props.open, (val) => {
        dialog.value = val
    })

    watch(() => props.text, (val) => {
        text.value = val
    })
    const handleConfirm = (val) => {
        dialog.value = false
        emit('confirm', val)
    }

    
    return { dialog, handleConfirm }
  }
}
</script> -->

<template>
  <div class="text-center pa-4">
    <v-dialog
      v-model="dialog"
      max-width="400"
      persistent
    >
      <v-card prepend-icon="mdi-map-marker" :text="props.text">
        <template #actions>
          <v-spacer></v-spacer>

          <v-btn
            @click="handleConfirm('confirm')"
            class="ms-4 text-white"
            color="blue-darken-4"
            variant="flat"
          >
            ตกลง
          </v-btn>

          <v-btn
            @click="handleConfirm('cancel')"
            color="red"
            class="text-white"
            variant="flat"
          >
            ยกเลิก
          </v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// props + emit
const props = defineProps({
  open: { type: Boolean, required: true },
  text: { type: String, required: true }
})

const emit = defineEmits(['confirm', 'update:open'])

// dialog เป็น local state ที่ sync กับ props.open
const dialog = ref(props.open)

watch(
  () => props.open,
  (val) => {
    dialog.value = val
  }
)

watch(dialog, (val) => {
  if (!val) {
    emit('update:open', val)
  }
})

const handleConfirm = (action) => {
  dialog.value = false
  emit('confirm', action)
}
</script>

