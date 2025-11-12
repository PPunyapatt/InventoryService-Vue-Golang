<!-- Barcode.vue -->
<template>
  <svg ref="svg"></svg>
</template>

<script>
import { onMounted, ref, watch } from 'vue'
import JsBarcode from 'jsbarcode'

export default {
  props: {
    value: { type: String, required: true }
  },
  setup(props) {
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

    onMounted(renderBarcode)

    // ถ้า props.value เปลี่ยน ให้ re-generate
    watch(() => props.value, renderBarcode)

    return { svg }
  }
}
</script>
