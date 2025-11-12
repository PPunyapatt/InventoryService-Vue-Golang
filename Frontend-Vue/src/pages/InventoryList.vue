<template>
  <div class="container">

    <!-- รหัสสินค้า -->
    <v-row align="center">
      <v-col cols="2" style="margin-bottom: 20px;">
        <h3>รหัสสินค้า</h3>
      </v-col>
      <v-col cols="9">
        <v-text-field  
            label=""
          :rules="rules"
          v-model="state.text" 
          variant="outlined" 
          density="compact" 
          placeholder="XXXX-XXXX-XXXX-XXXX"
          @input="formatCode"
        />
      </v-col>
      <v-col cols="1" style="margin-bottom: 20px;">
        <v-btn 
            color="primary" 
            :disabled="state.disable_add"
            block
        >ADD</v-btn>
      </v-col>
    </v-row>
    <v-table 
        class="full-height-table"
        ref="tableWrapper" 
        fixed-header 
        style="margin-top: 10px;">
      <thead>
        <tr>
          <th class="text-center">id</th>
          <th class="text-center" style="width: 35%">รหัสสินค้า (16 หลัก)</th>
          <th class="text-center">บาร์โค้ดสินค้า</th>
          <th class="text-center" style="width: 20%">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, idx) in desserts" :key="item.name">
          <td class="text-center">{{ idx + 1 }}</td>
          <td class="text-center">{{ formatName(item.name) }}</td>
          <td class="text-center">
            <Barcode :value="item.name" />
          </td>
          <td class="text-center"><v-btn color="red" size="small">ลบ</v-btn></td>
        </tr>

        <tr v-if="state.loading">
          <td colspan="4" class="text-center">
            <v-progress-circular color="primary" indeterminate></v-progress-circular>
          </td>
        </tr>

      </tbody>
    </v-table>

  </div>
</template>

<script>
import { reactive, onMounted, ref } from 'vue'
import Barcode from '../components/Barcode.vue'
import InventoryService from "../services/api-services"

export default {
  components: { Barcode },
  setup() {
    const desserts = [
      { name: 'ASDV123BXXXXN899', calories: 159 },
      { name: 'ASDA123BAXXXN899', calories: 237 },
      { name: 'ASDC123BBXXXN899', calories: 262 },
      { name: 'ASDX123BCXXXN899', calories: 305 },
      { name: 'ASDB123BDXXXN899', calories: 356 },
      { name: 'Jelly bean', calories: 375 },
      { name: 'Lollipop', calories: 392 },
      { name: 'Honeycomb', calories: 408 },
      { name: 'Donut', calories: 452 },
    ]

    const state = reactive({ 
      text: "",
      loading: false,
      disable_add: true,
      inventories: []
    })

    const tableWrapper = ref(null)

    const formatCode = () => {
      state.text = state.text.toUpperCase()
      let cleaned = state.text.replace(/\W/g, '')
      if (cleaned.length > 16) cleaned = cleaned.slice(0, 16)
      const parts = cleaned.match(/.{1,4}/g)
      state.text = parts ? parts.join('-') : ''
    }

    const rules = [
      value => {
        // console.log("val: ", value);
      },
      value => {
        if (!value) return true // เว้นกรณีว่างไว้ให้ rule แรกจัดการ
        const cleaned = value.replace(/-/g, '')
        if (cleaned.length === 16) {
            state.disable_add = false
        }
        return cleaned.length === 16 || 'Must be 16 characters'
      }
    ]

    const onScroll = () => {
      const el = tableWrapper.value?.$el.querySelector('.v-table__wrapper')
      if (!el) return
      if (el.scrollTop + el.clientHeight >= el.scrollHeight - 1) {
        console.log("load more");
        state.loading = true
      }
    }

    const formatName = (name) => {
        const parts = name.match(/.{1,4}/g)
        return parts.join('-')
    }


    onMounted(() => {
        InventoryService.getInventories()
        .then(response => {
          state.inventories = response.data;
        //   console.log(response.data);
        })
        .catch(e => {
          console.log(e);
        });

        const el = tableWrapper.value?.$el.querySelector('.v-table__wrapper')
        if (el) el.addEventListener('scroll', onScroll)
    })

    return { 
      state,
      desserts,
      formatCode, 
      rules, 
      tableWrapper,
      formatName 
    }
  }
}
</script>

<style>
.container {
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 20px;
  max-width: 900px;
  width: 100%;
  box-sizing: border-box;
  margin: 30px auto;
}

th:nth-child(1) { width: 10%; }
th:nth-child(2) { width: 40%; }
th:nth-child(3) { width: 40%; }
th:nth-child(4) { width: 10%; }

.v-table.v-table--fixed-header > .v-table__wrapper > table > thead > tr > th {
  background-color: rgb(81, 199, 236) !important;
  font-size: 16px;
  font-weight: bold;
}

.full-height-table {
  height: calc(100vh - 150px); /* ลบ header, input row หรือ margin-top อื่น ๆ */
}
.v-table.v-table--fixed-header > .v-table__wrapper {
  height: 100%;
}

</style>
