<template>
  <svg ref="svg"></svg>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import JsBarcode from 'jsbarcode'

// รับ props
const props = defineProps({
  value: {
    type: String,
    required: true
  }
})

const svg = ref(null)

const renderBarcode = () => {
  if (svg.value) {
    JsBarcode(svg.value, props.value, {
      format: 'CODE39',
      lineColor: '#000',
      width: 1,
      height: 30,
      displayValue: false
    })
  }
}

// render ครั้งแรก
onMounted(renderBarcode)

// re-render เมื่อ props.value เปลี่ยน
watch(() => props.value, renderBarcode)
</script>

